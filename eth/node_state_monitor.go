package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
	"sync"
	"time"
)

const (
	StateRunning = iota + 1
	StateStop
)

const StateBroadcastDuration = 60
const StateUploadDuration    = 120
const MaxMissNum = 5

type MonitorInfo struct {
	curState int64
	missNum  int64 // no node-ping msg
	lastTime int64
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
	monitor.uploadTicker = time.NewTicker(StateUploadDuration * time.Second)
	monitor.uploadTickerStopCh = make(chan struct{})
	go monitor.uploadLoop()

	monitor.wg.Add(1)
	monitor.broadcastTicker = time.NewTicker(StateBroadcastDuration * time.Second)
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
			if monitor.isSyncing() {
				break
			}

			addr, err := monitor.e.Etherbase()
			if err == nil && monitor.isSuperNode(addr) {
				ping := ev.Ping
				log.Trace("node-state-monitor", "ping", ping)

				// recover the public key from the signature
				r, s := ping.R.Bytes(), ping.S.Bytes()
				v := byte(ping.V.Uint64() - 27)
				sig := make([]byte, crypto.SignatureLength)
				copy(sig[32-len(r):32], r)
				copy(sig[64-len(s):64], s)
				sig[64] = v
				pub, err := crypto.Ecrecover(ping.Hash().Bytes(), sig)
				if err != nil || len(pub) == 0 || pub[0] != 4 {
					log.Warn("node-state-monitor", "ping", ping, "error", "recover public key failed")
					break
				}

				curTime := time.Now().Unix()
				nodeType := ping.NodeType.Int64()
				if nodeType == int64(types.MasterNodeType) {
					monitor.lock.Lock()
					monitor.mnMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
					monitor.lock.Unlock()
					info, err := contract_api.GetMasterNodeInfoByID(monitor.ctx, monitor.blockChainAPI, ping.Id, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(monitor.e.blockchain.CurrentBlock().Number().Int64())))
					if err != nil || hexutils.BytesToHex(pub)[1:] == GetPubKeyFromEnode(info.Enode) {
						log.Warn("node-state-monitor", "ping", ping, "error", "verify signature failed")
						break
					}
				} else if nodeType == int64(types.SuperNodeType) {
					monitor.lock.Lock()
					monitor.snMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
					monitor.lock.Unlock()
					info, err := contract_api.GetSuperNodeInfoByID(monitor.ctx, monitor.blockChainAPI, ping.Id, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(monitor.e.blockchain.CurrentBlock().Number().Int64())))
					if err != nil || hexutils.BytesToHex(pub)[1:] == GetPubKeyFromEnode(info.Enode) {
						log.Warn("node-state-monitor", "ping", ping, "error", "verify signature failed")
						break
					}
				}
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

			if monitor.isSyncing() {
				if lastAddr != addr {
					log.Warn("syncing now, wait...")
					lastAddr = addr
				}
				break
			}

			curTime := time.Now().Unix()
			curBlock := monitor.e.blockchain.CurrentBlock()
			info1, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(curBlock.Number().Int64())))
			if err == nil {
				if monitor.enode != info1.Enode {
					if lastAddr != addr {
						log.Error("broadcast-supernode-state", "local", monitor.enode, "state", info1.Enode, "error", "incompatible enode")
						lastAddr = addr
					}
					break
				}
				ping, _ := types.NewNodePing(info1.Id, types.SuperNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
				monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})
				monitor.lock.Lock()
				monitor.snMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
				monitor.lock.Unlock()
			} else {
				info2, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(curBlock.Number().Int64())))
				if err == nil {
					if monitor.enode != info2.Enode {
						if lastAddr != addr {
							log.Error("broadcast-masternode-state", "local", monitor.enode, "state", info2.Enode, "error", "incompatible enode")
							lastAddr = addr
						}
						break
					}
					ping, _ := types.NewNodePing(info2.Id, types.MasterNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
					monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})
					monitor.lock.Lock()
					monitor.mnMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
					monitor.lock.Unlock()
				}
			}
		}
	}
}

func (monitor *NodeStateMonitor) isSuperNode(addr common.Address) bool {
	topAddrs, err := contract_api.GetTopSuperNodes(monitor.ctx, monitor.blockChainAPI, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(monitor.e.blockchain.CurrentBlock().Number().Int64())))
	if err != nil {
		return false
	}
	flag := false
	for i := range topAddrs {
		if topAddrs[i] == addr { //info.Addr == addr && info.Enode == monitor.enode {
			flag = true
			break
		}
	}
	return flag
}

func (monitor *NodeStateMonitor) collectMasterNodes() ([]*big.Int, []*big.Int) {
	var ids []*big.Int
	var states []*big.Int

	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(monitor.e.blockchain.CurrentBlock().Number().Int64()))
	num, err := contract_api.GetMasterNodeNum(monitor.ctx, monitor.blockChainAPI, blockNrOrHash)
	if err != nil {
		return ids, states
	}

	batch := num.Int64() / 100
	if num.Int64() % 100 != 0 {
		batch++
	}

	var infos []types.MasterNodeInfo
	for i := int64(0); i < batch; i++ {
		mnAddrs, err := contract_api.GetAllMasterNodes(monitor.ctx, monitor.blockChainAPI, big.NewInt(i*100), big.NewInt(100), blockNrOrHash)
		if err != nil {
			return ids, states
		}
		for _, addr := range mnAddrs {
			info, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, blockNrOrHash)
			if err != nil {
				continue
			}
			infos = append(infos, *info)
		}
	}

	curTime := time.Now().Unix()
	var info types.MasterNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.mnMonitorInfos[id]; ok {
			if v.curState != StateRunning || curTime > v.lastTime + StateUploadDuration {
				v.curState = StateStop
				v.missNum++
				v.lastTime = curTime
				monitor.mnMonitorInfos[id] = v
			}
		} else {
			monitor.mnMonitorInfos[id] = MonitorInfo{StateStop, 1, curTime}
		}
	}

	for _, info = range infos {
		id := info.Id.Int64()
		log.Trace("collect-masternode-state", "id", id, "global-state", info.State, "local-state", monitor.mnMonitorInfos[id].curState, "missNum", monitor.mnMonitorInfos[id].missNum)
		if v, ok := monitor.mnMonitorInfos[id]; ok {
			if v.curState != info.State.Int64() {
				if v.curState == StateRunning || (v.curState == StateStop && v.missNum >= MaxMissNum) {
					ids = append(ids, info.Id)
					states = append(states, big.NewInt(v.curState))
					delete(monitor.mnMonitorInfos, id)
				}
			}
		}
	}
	return ids, states
}

func (monitor *NodeStateMonitor) collectSuperNodes() ([]*big.Int, []*big.Int) {
	var ids []*big.Int
	var states []*big.Int
	infos, err := contract_api.GetAllSuperNodes(monitor.ctx, monitor.blockChainAPI, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(monitor.e.blockchain.CurrentBlock().Number().Int64())))
	if err != nil {
		return ids, states
	}
	curTime := time.Now().Unix()
	var info types.SuperNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.snMonitorInfos[id]; ok {
			if v.curState != StateRunning || curTime > v.lastTime + StateUploadDuration {
				v.curState = StateStop
				v.missNum++
				v.lastTime = curTime
				monitor.snMonitorInfos[id] = v
			}
		} else {
			monitor.snMonitorInfos[id] = MonitorInfo{StateStop, 1, curTime}
		}
	}
	for _, info = range infos {
		id := info.Id.Int64()
		log.Trace("collect-supernode-state", "id", id, "global-state", info.State, "local-state", monitor.snMonitorInfos[id].curState, "missNum", monitor.snMonitorInfos[id].missNum)
		if v, ok := monitor.snMonitorInfos[id]; ok {
			if v.curState != info.State.Int64() {
				if v.curState == StateRunning || (v.curState == StateStop && v.missNum >= MaxMissNum) {
					ids = append(ids, info.Id)
					states = append(states, big.NewInt(v.curState))
					delete(monitor.snMonitorInfos, id)
				}
			}
		}
	}
	return ids, states
}

func (monitor *NodeStateMonitor) isSyncing() bool {
	progress := monitor.e.APIBackend.SyncProgress()
	return progress.CurrentBlock < progress.HighestBlock
}

func GetPubKeyFromEnode(enode string) string {
	ret := ""
	pos1 := strings.Index(enode, "enode://")
	pos2 := strings.Index(enode, "@")
	if pos1 == -1 || pos2 == -1 {
		return ret
	}
	ret = enode[pos1 + len("enode://") : pos2]
	return ret
}