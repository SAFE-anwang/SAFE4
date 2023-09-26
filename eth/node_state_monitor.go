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
			if monitor.isSuperNode() {
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
			if monitor.isSuperNode() {
				monitor.lock.Lock()
				mnIDs, mnStates := monitor.collectMasterNodes()
				snIDs, snStates := monitor.collectSuperNodes()
				monitor.lock.Unlock()

				if len(mnIDs) != 0 && len(mnIDs) == len(mnStates) {
					hash, err := contract_api.UploadMasterNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, monitor.e.etherbase, mnIDs, mnStates)
					log.Info("upload-masternode-state", "ids", mnIDs, "states", mnStates, "hash", hash.Hex(), "error", err)
				}
				if len(snIDs) != 0 && len(snIDs) == len(snStates) {
					hash, err := contract_api.UploadSuperNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, monitor.e.etherbase, snIDs, snStates)
					log.Info("upload-supernode-state", "ids", snIDs, "states", snStates, "hash", hash.Hex(), "error", err)
				}
			}
		}
	}
}

func (monitor *NodeStateMonitor) broadcastLoop() {
	defer monitor.wg.Done()
	for {
		select {
		case <- monitor.broadcastTickerStopCh:
			return
		case <- monitor.broadcastTicker.C:
			addr := monitor.e.etherbase
			curBlock := monitor.e.blockchain.CurrentBlock()
			info1, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, addr)
			if err == nil {
				log.Info("broadcast-supernode-state", "id", info1.Id, "block hash", curBlock.Hash(), "block number", curBlock.Number())
				ping, _ := types.NewNodePing(info1.Id, types.SuperNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
				monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})
			} else {
				info2, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr)
				if err == nil {
					log.Info("broadcast-masternode-state", "id", info2.Id, "block hash", curBlock.Hash(), "block number", curBlock.Number())
					ping, _ := types.NewNodePing(info2.Id, types.MasterNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
					monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})
				}
			}
		}
	}
}

func (monitor *NodeStateMonitor) isSuperNode() bool {
	infos, err := contract_api.GetTopSuperNode(monitor.ctx, monitor.blockChainAPI)
	if err != nil {
		return false
	}
	if monitor.e.etherbase == (common.Address{}) {
		if wallets := monitor.e.AccountManager().Wallets(); len(wallets) > 0 {
			if accounts := wallets[0].Accounts(); len(accounts) > 0 {
				etherbase := accounts[0].Address
				monitor.e.lock.Lock()
				monitor.e.etherbase = etherbase
				monitor.e.lock.Unlock()
				log.Info("Etherbase automatically configured", "address", etherbase)
			}
		}
	}
	flag := false
	//log.Info("node-state-monitor", "coinbase", monitor.e.etherbase, "enode", monitor.e.p2pServer.NodeInfo().Enode)
	for _, info := range infos {
		//log.Info("node-state-monitor", "addr", info.Addr, "enode", info.Enode)
		if info.Addr == monitor.e.etherbase /*&& info.Enode == monitor.e.p2pServer.NodeInfo().Enode*/ {
			flag = true
			break
		}
	}
	return flag
}

func (monitor *NodeStateMonitor) collectMasterNodes() ([]int64, []uint8) {
	var ids []int64
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
					ids = append(ids, id)
					states = append(states, v.curState)
					delete(monitor.snMonitorInfos, id)
				}
			}
		}
	}
	return ids, states
}

func (monitor *NodeStateMonitor) collectSuperNodes() ([]int64, []uint8) {
	var ids []int64
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
					ids = append(ids, id)
					states = append(states, v.curState)
					delete(monitor.snMonitorInfos, id)
				}
			}
		}
	}
	return ids, states
}