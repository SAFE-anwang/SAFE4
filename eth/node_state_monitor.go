package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
	"sync"
	"time"
)

const (
	StateRunning uint8 = iota + 1
	StateStop
)

const StateBroadcastDuration = time.Second * 20
const StateUploadDuration    = time.Second * 60
const MaxMissNum = 3

type MonitorInfo struct {
	curState uint8
	missNum  int // no node-ping msg
}

type NodeStateMonitor struct {
	ctx         context.Context
	cancelCtx   context.CancelFunc

	e           *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI

	eventSub    *event.TypeMuxSubscription

	broadcastTicker        *time.Ticker
	broadcastTickerStopCh  chan struct{}

	uploadTicker      *time.Ticker
	uploadTickerStopCh chan struct{}

	wg          sync.WaitGroup
	lock        sync.RWMutex

	mnMonitorInfos    map[int64]MonitorInfo
	snMonitorInfos    map[int64]MonitorInfo

	enode       string
}

func newNodeStateMonitor(e *Ethereum) (*NodeStateMonitor, error) {
	monitor := &NodeStateMonitor{}
	monitor.ctx, monitor.cancelCtx = context.WithCancel(context.Background())
	monitor.e = e
	monitor.blockChainAPI = monitor.e.GetPublicBlockChainAPI()
	monitor.transactionPoolAPI = monitor.e.GetPublicTransactionPoolAPI()
	monitor.mnMonitorInfos = make(map[int64]MonitorInfo)
	monitor.snMonitorInfos = make(map[int64]MonitorInfo)
	return monitor, nil
}

func (monitor *NodeStateMonitor) Start() {
	temp := monitor.e.p2pServer.NodeInfo().Enode
	arr := strings.Split(temp, "?")
	if len(arr) == 0 {
		log.Error("invalid local enode")
		return
	}
	monitor.enode = arr[0]

	monitor.wg.Add(1)
	monitor.eventSub = monitor.e.eventMux.Subscribe(core.NodePingEvent{})
	go monitor.loop()

	monitor.wg.Add(1)
	monitor.uploadTicker = time.NewTicker(StateUploadDuration)
	monitor.uploadTickerStopCh = make(chan struct{})
	go monitor.uploadLoop()

	monitor.wg.Add(1)
	monitor.broadcastTicker = time.NewTicker(StateBroadcastDuration)
	monitor.broadcastTickerStopCh = make(chan struct{})
	go monitor.broadcastLoop()
}

func (monitor *NodeStateMonitor) Stop() {
	monitor.eventSub.Unsubscribe()
	monitor.uploadTickerStopCh <- struct{}{}
	monitor.uploadTicker.Stop()
	monitor.broadcastTickerStopCh <- struct{}{}
	monitor.broadcastTicker.Stop()
	monitor.wg.Wait()
	monitor.cancelCtx()
	log.Info("Node monitor stopped")
}

func (monitor *NodeStateMonitor) loop() {
	defer monitor.wg.Done()
	for obj := range monitor.eventSub.Chan() {
		switch ev := obj.Data.(type) {
		case core.NodePingEvent:
			addr, err := monitor.e.Etherbase()
			if err == nil && monitor.isSuperNode(addr) {
				ping := ev.Ping
				//log.Info("node-state-monitor", "ping", ping)
				nodeType := ping.NodeType.Int64()
				monitor.lock.Lock()
				if nodeType == int64(types.MasterNodeType) {
					monitor.mnMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0}
				} else if nodeType == int64(types.SuperNodeType) {
					monitor.snMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0}
				}
				monitor.lock.Unlock()
			}
		}
	}
}

func (monitor *NodeStateMonitor) uploadLoop() {
	defer monitor.wg.Done()
	for {
		select {
		case <- monitor.uploadTickerStopCh:
			return
		case <- monitor.uploadTicker.C:
			addr, err := monitor.e.Etherbase()
			if err == nil && monitor.isSuperNode(addr) {
				monitor.lock.Lock()
				mnIDs, mnStates := monitor.collectMasterNodes()
				snIDs, snStates := monitor.collectSuperNodes()
				monitor.lock.Unlock()

				if len(mnIDs) != 0 && len(mnIDs) == len(mnStates) {
					hash, err := contract_api.UploadMasterNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, addr, mnIDs, mnStates)
					log.Info("upload-masternode-state", "ids", mnIDs, "states", mnStates, "hash", hash.Hex(), "error", err)
				}
				if len(snIDs) != 0 && len(snIDs) == len(snStates) {
					hash, err := contract_api.UploadSuperNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, addr, snIDs, snStates)
					log.Info("upload-supernode-state", "ids", snIDs, "states", snStates, "hash", hash.Hex(), "error", err)
				}
			}
		}
	}
}

func (monitor *NodeStateMonitor) broadcastLoop() {
	defer monitor.wg.Done()
	lastAddr := common.Address{}
	for {
		select {
		case <- monitor.broadcastTickerStopCh:
			return
		case <- monitor.broadcastTicker.C:
			addr, err := monitor.e.Etherbase()
			if err != nil {
				break
			}
			curBlock := monitor.e.blockchain.CurrentBlock()
			info1, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, addr)
			if err == nil {
				if monitor.enode != info1.Enode {
					if lastAddr != addr {
						log.Error("incompatible supernode enode", "local", monitor.enode, "state", info1.Enode)
						lastAddr = addr
					}
					break
				}
				//log.Info("broadcast-supernode-state", "id", info1.Id, "block hash", curBlock.Hash(), "block number", curBlock.Number())
				ping, _ := types.NewNodePing(info1.Id, types.SuperNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
				monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})
			} else {
				info2, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr)
				if err == nil {
					if monitor.enode != info2.Enode {
						if lastAddr != addr {
							log.Error("incompatible masternode enode", "local", monitor.enode, "state", info2.Enode)
							lastAddr = addr
						}
						break
					}
					//log.Info("broadcast-masternode-state", "id", info2.Id, "block hash", curBlock.Hash(), "block number", curBlock.Number())
					ping, _ := types.NewNodePing(info2.Id, types.MasterNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
					monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})
				}
			}
		}
	}
}

func (monitor *NodeStateMonitor) isSuperNode(addr common.Address) bool {
	infos, err := contract_api.GetTopSuperNode(monitor.ctx, monitor.blockChainAPI)
	if err != nil {
		return false
	}
	flag := false
	//log.Info("node-state-monitor", "coinbase", addr, "enode", monitor.enode)
	for _, info := range infos {
		//log.Info("node-state-monitor", "addr", info.Addr, "enode", info.Enode)
		if info.Addr == addr && info.Enode == monitor.enode {
			flag = true
			break
		}
	}
	return flag
}

func (monitor *NodeStateMonitor) collectMasterNodes() ([]*big.Int, []uint8) {
	var ids []*big.Int
	var states []uint8
	infos, err := contract_api.GetAllMasterNode(monitor.ctx, monitor.blockChainAPI)
	if err != nil {
		return ids, states
	}
	var info types.MasterNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.mnMonitorInfos[id]; ok {
			if v.curState != StateRunning {
				v.curState = StateStop
				v.missNum++
				monitor.mnMonitorInfos[id] = v
			}
		} else {
			monitor.mnMonitorInfos[id] = MonitorInfo{StateStop, 1}
		}
	}
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.mnMonitorInfos[id]; ok {
			if v.curState != info.StateInfo.State {
				if v.curState == StateRunning || (v.curState == StateStop && v.missNum >= MaxMissNum) {
					ids = append(ids, info.Id)
					states = append(states, v.curState)
					delete(monitor.mnMonitorInfos, id)
				}
			}
		}
	}
	return ids, states
}

func (monitor *NodeStateMonitor) collectSuperNodes() ([]*big.Int, []uint8) {
	var ids []*big.Int
	var states []uint8
	infos, err := contract_api.GetAllSuperNode(monitor.ctx, monitor.blockChainAPI)
	if err != nil {
		return ids, states
	}
	var info types.SuperNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.snMonitorInfos[id]; ok {
			if v.curState != StateRunning {
				v.curState = StateStop
				v.missNum++
				monitor.snMonitorInfos[id] = v
			}
		} else {
			monitor.snMonitorInfos[id] = MonitorInfo{StateStop, 1}
		}
	}
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.snMonitorInfos[id]; ok {
			if v.curState != info.StateInfo.State {
				if v.curState == StateRunning || (v.curState == StateStop && v.missNum >= MaxMissNum) {
					ids = append(ids, info.Id)
					states = append(states, v.curState)
					delete(monitor.snMonitorInfos, id)
				}
			}
		}
	}
	return ids, states
}