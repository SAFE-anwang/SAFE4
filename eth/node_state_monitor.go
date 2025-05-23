package eth

import (
	"context"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"math/rand"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	StateRunning = iota + 1
	StateStop
)

const mnStateBroadcastDuration = 3600 // 60 minute
const snStateBroadcastDuration = 180  // 3 minute

const mnStateUploadDuration = 311
const snStateUploadDuration = 67

const mnMaxMissNum = 12 // 0.5 day
const snMaxMissNum = 4  // 12 minute = 24 block

const coinbaseDuration = 60
const addSnDuration = 30

const batchSize = 20

type MonitorInfo struct {
	curState int64
	missNum  int64 // no node-ping msg
	lastTime int64
}

type NodeStateMonitor struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI

	broadcastStopCh     chan struct{}
	uploadSnStateStopCh chan struct{}
	uploadMnStateStopCh chan struct{}
	coinbaseStopCh      chan struct{}
	addSnStopCh         chan struct{}

	wg     sync.WaitGroup
	mnLock sync.RWMutex
	snLock sync.RWMutex

	mnMonitorInfos map[int64]MonitorInfo
	snMonitorInfos map[int64]MonitorInfo

	knownPings        *knownCache
	lastMnPingHeights map[int64]int64
	lastSnPingHeights map[int64]int64

	trustedPeers map[enode.ID]struct{}
	trustedLock  sync.RWMutex

	enode string
	exit  int32
}

func newNodeStateMonitor() *NodeStateMonitor {
	monitor := &NodeStateMonitor{}
	monitor.ctx, monitor.cancelCtx = context.WithCancel(context.Background())
	monitor.mnMonitorInfos = make(map[int64]MonitorInfo)
	monitor.snMonitorInfos = make(map[int64]MonitorInfo)
	monitor.knownPings = newKnownCache(20480)
	monitor.lastMnPingHeights = make(map[int64]int64)
	monitor.lastSnPingHeights = make(map[int64]int64)
	atomic.StoreInt32(&monitor.exit, 0)
	return monitor
}

func (monitor *NodeStateMonitor) Start(e *Ethereum) {
	monitor.e = e
	monitor.blockChainAPI = monitor.e.GetPublicBlockChainAPI()
	monitor.transactionPoolAPI = monitor.e.GetPublicTransactionPoolAPI()
	monitor.enode = contract_api.CompressEnode(monitor.e.p2pServer.NodeInfo().Enode)

	monitor.wg.Add(1)
	monitor.uploadSnStateStopCh = make(chan struct{})
	go monitor.uploadSnStateLoop()

	monitor.wg.Add(1)
	monitor.uploadMnStateStopCh = make(chan struct{})
	go monitor.uploadMnStateLoop()

	monitor.wg.Add(1)
	monitor.broadcastStopCh = make(chan struct{})
	go monitor.broadcastLoop()

	monitor.wg.Add(1)
	monitor.coinbaseStopCh = make(chan struct{})
	go monitor.coinbaseLoop()

	monitor.wg.Add(1)
	monitor.addSnStopCh = make(chan struct{})
	go monitor.addSuperNodePeerLoop()
}

func (monitor *NodeStateMonitor) Stop() {
	atomic.StoreInt32(&monitor.exit, 1)
	close(monitor.addSnStopCh)
	close(monitor.uploadSnStateStopCh)
	close(monitor.uploadMnStateStopCh)
	close(monitor.broadcastStopCh)
	close(monitor.coinbaseStopCh)
	monitor.wg.Wait()
	monitor.cancelCtx()
	log.Info("Node monitor stopped")
}

func (monitor *NodeStateMonitor) HandlePing(ping *types.NodePing) error {
	if atomic.LoadInt32(&monitor.exit) == 1 {
		return fmt.Errorf("node state monitor is exiting")
	}

	h := ping.Hash()
	if monitor.knownPings.Contains(h) {
		return nil
	}
	monitor.knownPings.Add(h)

	nodeType := ping.NodeType.Int64()
	id := ping.Id.Int64()
	pingHeight := ping.CurHeight.Int64()
	log.Trace("handleNodePing", "node-type", nodeType, "node-id", id, "node-height", pingHeight, "hash", h)

	if nodeType == int64(types.MasterNodeType) {
		monitor.mnLock.Lock()
		if monitor.lastMnPingHeights[id] >= pingHeight { // discard expired ping
			monitor.mnLock.Unlock()
			return nil
		}
		if monitor.lastMnPingHeights[id] != 0 && monitor.lastMnPingHeights[id] > pingHeight-110 { // decrease broadcast frequency
			monitor.mnLock.Unlock()
			return nil
		}
		monitor.lastMnPingHeights[id] = pingHeight
		monitor.mnLock.Unlock()
	} else {
		monitor.snLock.Lock()
		if monitor.lastSnPingHeights[id] >= pingHeight { // discard expired ping
			monitor.snLock.Unlock()
			return nil
		}
		monitor.lastSnPingHeights[id] = pingHeight
		monitor.snLock.Unlock()
	}

	monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})

	if monitor.isSyncing() {
		return nil
	}

	addr, err := monitor.e.Etherbase()
	if err != nil || !monitor.isTopSuperNode(addr) {
		return nil
	}

	log.Debug("Handle received nodePing", "node-type", nodeType, "node-id", id, "node-height", pingHeight)
	localHeight := monitor.e.blockchain.CurrentBlock().Number().Int64()

	// recover the public key from the signature
	r, s := ping.R.Bytes(), ping.S.Bytes()
	v := byte(ping.V.Uint64() - 27)
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = v
	pub, err := crypto.Ecrecover(ping.Hash().Bytes(), sig)
	if err != nil || len(pub) == 0 || pub[0] != 4 {
		return fmt.Errorf("invalid nodePing, recover public key failed")
	}

	if nodeType == int64(types.MasterNodeType) {
		info, err := contract_api.GetMasterNodeInfoByID(monitor.ctx, monitor.blockChainAPI, ping.Id, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
		if err != nil || hexutils.BytesToHex(pub)[1:] == GetPubKeyFromEnode(info.Enode) {
			return fmt.Errorf("invalid masternode ping, verify signature failed")
		}

		if _, err := CheckPublicIP(info.Enode); err != nil {
			return fmt.Errorf("invalid masternode ping, %s", err.Error())
		}

		monitor.mnLock.Lock()
		monitor.mnMonitorInfos[id] = MonitorInfo{StateRunning, 0, time.Now().Unix()}
		if ping.Version.Int64() < int64(types.NodePingVersion) || localHeight > pingHeight+360 {
			monitor.mnMonitorInfos[id] = MonitorInfo{StateStop, mnMaxMissNum, time.Now().Unix()}
		}
		monitor.mnLock.Unlock()
	} else {
		info, err := contract_api.GetSuperNodeInfoByID(monitor.ctx, monitor.blockChainAPI, ping.Id, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
		if err != nil || hexutils.BytesToHex(pub)[1:] == GetPubKeyFromEnode(info.Enode) {
			return fmt.Errorf("invalid supernode ping, verify signature failed")
		}

		if _, err := CheckPublicIP(info.Enode); err != nil {
			return fmt.Errorf("invalid supernode ping, %s", err.Error())
		}

		monitor.snLock.Lock()
		monitor.snMonitorInfos[id] = MonitorInfo{StateRunning, 0, time.Now().Unix()}
		if ping.Version.Int64() < int64(types.NodePingVersion) || localHeight > pingHeight+6 {
			monitor.snMonitorInfos[id] = MonitorInfo{StateStop, snMaxMissNum, time.Now().Unix()}
		}
		monitor.snLock.Unlock()
	}
	return nil
}

func (monitor *NodeStateMonitor) uploadMnState(addr common.Address, ids []*big.Int, states []*big.Int) {
	hash, err := contract_api.UploadMasterNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, addr, ids, states)
	log.Info("Upload masternode state", "caller", addr, "ids", ids, "states", states, "hash", hash.Hex(), "error", err)
}

func (monitor *NodeStateMonitor) uploadMnStateLoop() {
	ticker := time.NewTicker(mnStateUploadDuration * time.Second)
	defer ticker.Stop()
	defer monitor.wg.Done()
	once := 1
	for {
		select {
		case <-monitor.uploadMnStateStopCh:
			log.Info("Exit node-state-monitor uploadMnStateLoop")
			return
		case <-ticker.C:
			if monitor.isSyncing() {
				if once == 1 {
					log.Info("uploadMnStateLoop: syncing now, wait a moment...")
					once = 0
				}
				continue
			}
			addr, err := monitor.e.Etherbase()
			if err != nil {
				continue
			}
			if monitor.isTopSuperNode(addr) {
				mnIDs, mnStates := monitor.collectMasterNodes(addr)
				if len(mnIDs) != 0 && len(mnIDs) == len(mnStates) {
					monitor.uploadMnState(addr, mnIDs, mnStates)
				}
			}
		}
	}
}

func (monitor *NodeStateMonitor) uploadSnState(addr common.Address, ids []*big.Int, states []*big.Int) {
	hash, err := contract_api.UploadSuperNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, addr, ids, states)
	log.Info("Upload supernode state", "caller", addr, "ids", ids, "states", states, "hash", hash.Hex(), "error", err)
}

func (monitor *NodeStateMonitor) uploadSnStateLoop() {
	ticker := time.NewTicker(snStateUploadDuration * time.Second)
	defer ticker.Stop()
	defer monitor.wg.Done()
	once := 1
	for {
		select {
		case <-monitor.uploadSnStateStopCh:
			log.Info("Exit node-state-monitor uploadSnStateLoop")
			return
		case <-ticker.C:
			if monitor.isSyncing() {
				if once == 1 {
					log.Info("uploadSnStateLoop: syncing now, wait a moment...")
					once = 0
				}
				continue
			}
			addr, err := monitor.e.Etherbase()
			if err != nil {
				continue
			}
			if monitor.isTopSuperNode(addr) {
				snIDs, snStates := monitor.collectSuperNodes(addr)
				if len(snIDs) != 0 && len(snIDs) == len(snStates) {
					monitor.uploadSnState(addr, snIDs, snStates)
				}
			}
		}
	}
}

var lastAddr common.Address

func (monitor *NodeStateMonitor) broadcastLoop() {
	mnTicker := time.NewTicker(mnStateBroadcastDuration * time.Second)
	snTicker := time.NewTicker(snStateBroadcastDuration * time.Second)
	defer mnTicker.Stop()
	defer snTicker.Stop()
	defer monitor.wg.Done()

	for monitor.isSyncing() {
		time.Sleep(5 * time.Second)
		if atomic.LoadInt32(&monitor.exit) == 1 {
			return
		}
	}

	// wait 60s
	tempTime := time.Now().Unix()
	for time.Now().Unix() < tempTime+60 {
		time.Sleep(5 * time.Second)
		if atomic.LoadInt32(&monitor.exit) == 1 {
			return
		}
	}

	monitor.snBroadcastPing()
	monitor.mnBroadcastPing()

	for {
		select {
		case <-monitor.broadcastStopCh:
			log.Info("Exit node-state-monitor broadcastLoop")
			return
		case <-snTicker.C:
			monitor.snBroadcastPing()
		case <-mnTicker.C:
			monitor.mnBroadcastPing()
		}
	}
}

func (monitor *NodeStateMonitor) snBroadcastPing() {
	addr, err := monitor.e.Etherbase()
	if err != nil {
		return
	}

	info, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil || info.Id.Int64() == 0 {
		return
	}

	if !contract_api.CompareEnode(monitor.enode, info.Enode) {
		if lastAddr != addr {
			log.Error("Broadcast supernode ping failed, incompatible enode", "local-enode", monitor.enode, "node-enode", info.Enode)
			lastAddr = addr
		}
		return
	}

	curBlock := monitor.e.blockchain.CurrentBlock()
	ping, err := types.NewNodePing(info.Id, types.SuperNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
	if err != nil {
		if lastAddr != addr {
			log.Error("Broadcast supernode ping failed", "error", err)
			lastAddr = addr
		}
		return
	}

	monitor.e.handler.BroadcastNodePing(ping)
	monitor.snLock.Lock()
	monitor.snMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, time.Now().Unix()}
	monitor.lastSnPingHeights[ping.Id.Int64()] = curBlock.Number().Int64()
	monitor.snLock.Unlock()
	log.Info("Broadcast supernode ping", "id", ping.Id, "height", ping.CurHeight)
}

func (monitor *NodeStateMonitor) mnBroadcastPing() {
	addr, err := monitor.e.Etherbase()
	if err != nil {
		return
	}

	info, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil || info.Id.Int64() == 0 {
		return
	}

	if !contract_api.CompareEnode(monitor.enode, info.Enode) {
		if lastAddr != addr {
			log.Error("Broadcast masternode ping failed, incompatible enode", "local-enode", monitor.enode, "node-enode", info.Enode)
			lastAddr = addr
		}
		return
	}

	curBlock := monitor.e.blockchain.CurrentBlock()
	ping, err := types.NewNodePing(info.Id, types.MasterNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
	if err != nil {
		if lastAddr != addr {
			log.Error("Broadcast masternode ping failed", "error", err)
			lastAddr = addr
		}
		return
	}

	monitor.e.handler.BroadcastNodePing(ping)
	monitor.mnLock.Lock()
	monitor.mnMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, time.Now().Unix()}
	monitor.lastMnPingHeights[ping.Id.Int64()] = curBlock.Number().Int64()
	monitor.mnLock.Unlock()
	log.Info("Broadcast masternode ping", "id", ping.Id, "height", ping.CurHeight)
}

func (monitor *NodeStateMonitor) coinbaseLoop() {
	ticker := time.NewTicker(coinbaseDuration * time.Second)
	defer ticker.Stop()
	defer monitor.wg.Done()
	lastAddr := common.Address{}
	for {
		select {
		case <-monitor.coinbaseStopCh:
			log.Info("Exit node-state-monitor coinbaseLoop")
			return
		case <-ticker.C:
			if _, state := monitor.e.miner.Pending(); state == nil {
				continue
			}
			if monitor.e.accountManager == nil {
				continue
			}
			wallets := monitor.e.accountManager.Wallets()
			for _, wallet := range wallets {
				accounts := wallet.Accounts()
				for _, account := range accounts {
					if monitor.isSuperNode(account.Address) {
						if lastAddr != account.Address {
							monitor.e.SetEtherbase(account.Address)
							lastAddr = account.Address
						}
						goto END
					}
				}
			}
			for _, wallet := range wallets {
				accounts := wallet.Accounts()
				for _, account := range accounts {
					if monitor.isMasterNode(account.Address) {
						if lastAddr != account.Address {
							monitor.e.SetEtherbase(account.Address)
							lastAddr = account.Address
						}
						goto END
					}
				}
			}
		END:
			// do nothing
		}
	}
}

func (monitor *NodeStateMonitor) isTopSuperNode(addr common.Address) bool {
	topAddrs, err := contract_api.GetTopSuperNodes(monitor.ctx, monitor.blockChainAPI, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return false
	}
	for _, snAddr := range topAddrs {
		info, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, snAddr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
		if err != nil || info.Id.Int64() == 0 {
			continue
		}
		if info.Addr == addr && contract_api.CompareEnode(monitor.enode, info.Enode) {
			return true
		}
	}
	return false
}

func (monitor *NodeStateMonitor) isSuperNode(addr common.Address) bool {
	info, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil || info.Id.Int64() == 0 {
		return false
	}
	return contract_api.CompareEnode(monitor.enode, info.Enode)
}

func (monitor *NodeStateMonitor) isMasterNode(addr common.Address) bool {
	info, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil || info.Id.Int64() == 0 {
		return false
	}
	return contract_api.CompareEnode(monitor.enode, info.Enode)
}

func (monitor *NodeStateMonitor) collectMasterNodes(from common.Address) ([]*big.Int, []*big.Int) {
	var ids []*big.Int
	var states []*big.Int

	pending, queued := monitor.e.txPool.Stats()
	if pending+queued >= 800 {
		return ids, states
	}

	num, err := contract_api.GetMasterNodeNum(monitor.ctx, monitor.blockChainAPI, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return ids, states
	}

	batch := num.Int64() / 100
	if num.Int64()%100 != 0 {
		batch++
	}

	var infos []types.MasterNodeInfo
	for i := int64(0); i < batch; i++ {
		mnAddrs, err := contract_api.GetAllMasterNodes(monitor.ctx, monitor.blockChainAPI, big.NewInt(i*100), big.NewInt(100), rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
		if err != nil {
			return ids, states
		}
		for _, addr := range mnAddrs {
			info, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
			if err != nil || info.Id.Int64() == 0 {
				continue
			}
			infos = append(infos, *info)
		}
	}

	var info types.MasterNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		if len(info.Enode) == 0 {
			monitor.mnLock.Lock()
			monitor.mnMonitorInfos[id] = MonitorInfo{StateStop, mnMaxMissNum, time.Now().Unix()}
			monitor.mnLock.Unlock()
			continue
		}
		monitor.mnLock.Lock()
		if v, ok := monitor.mnMonitorInfos[id]; ok {
			if time.Now().Unix() > v.lastTime+mnStateBroadcastDuration {
				v.curState = StateStop
				v.missNum++
				v.lastTime = time.Now().Unix()
				monitor.mnMonitorInfos[id] = v
			}
		} else {
			monitor.mnMonitorInfos[id] = MonitorInfo{StateStop, 0, time.Now().Unix()}
		}
		monitor.mnLock.Unlock()
	}

	// upload running state first
	for _, info = range infos {
		id := info.Id.Int64()
		monitor.mnLock.Lock()
		v, ok := monitor.mnMonitorInfos[id]
		monitor.mnLock.Unlock()
		if ok {
			log.Debug("collect-masternode-state-running", "id", id, "global-state", info.State, "local-state", v.curState, "missNum", v.missNum)
			if v.curState != info.State.Int64() {
				if v.curState == StateRunning || len(info.Enode) == 0 {
					tempState, err := contract_api.GetMasterNodeUploadState(monitor.ctx, monitor.blockChainAPI, info.Id, from, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
					if err != nil {
						continue
					}
					if tempState.Int64() != v.curState {
						ids = append(ids, info.Id)
						states = append(states, big.NewInt(v.curState))
						monitor.mnLock.Lock()
						delete(monitor.mnMonitorInfos, id)
						monitor.mnLock.Unlock()
						if len(ids) >= batchSize {
							break
						}
					}
				}
			}
		}
	}

	if len(ids) >= batchSize {
		return ids, states
	}

	pending, queued = monitor.e.txPool.Stats()
	if pending+queued >= 100 {
		return ids, states
	}

	// upload stopped state later
	for _, info = range infos {
		id := info.Id.Int64()
		monitor.mnLock.Lock()
		v, ok := monitor.mnMonitorInfos[id]
		monitor.mnLock.Unlock()
		if ok {
			log.Debug("collect-masternode-state-stop", "id", id, "global-state", info.State, "local-state", v.curState, "missNum", v.missNum)
			if v.curState != info.State.Int64() {
				if v.curState == StateStop && v.missNum >= mnMaxMissNum {
					tempState, err := contract_api.GetMasterNodeUploadState(monitor.ctx, monitor.blockChainAPI, info.Id, from, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
					if err != nil {
						continue
					}
					if tempState.Int64() != v.curState {
						ids = append(ids, info.Id)
						states = append(states, big.NewInt(v.curState))
						monitor.mnLock.Lock()
						delete(monitor.mnMonitorInfos, id)
						monitor.mnLock.Unlock()
						if len(ids) == batchSize {
							break
						}
					}
				}
			}
		}
	}
	return ids, states
}

func (monitor *NodeStateMonitor) collectSuperNodes(from common.Address) ([]*big.Int, []*big.Int) {
	var ids []*big.Int
	var states []*big.Int

	num, err := contract_api.GetSuperNodeNum(monitor.ctx, monitor.blockChainAPI, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return ids, states
	}

	batch := num.Int64() / 100
	if num.Int64()%100 != 0 {
		batch++
	}

	var infos []types.SuperNodeInfo
	for i := int64(0); i < batch; i++ {
		snAddrs, err := contract_api.GetAllSuperNodes(monitor.ctx, monitor.blockChainAPI, big.NewInt(i*100), big.NewInt(100), rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
		if err != nil {
			return ids, states
		}
		for _, addr := range snAddrs {
			info, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
			if err != nil || info.Id.Int64() == 0 {
				continue
			}
			infos = append(infos, *info)
		}
	}

	var info types.SuperNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		monitor.snLock.Lock()
		if v, ok := monitor.snMonitorInfos[id]; ok {
			if time.Now().Unix() > v.lastTime+snStateBroadcastDuration {
				v.curState = StateStop
				v.missNum++
				v.lastTime = time.Now().Unix()
				monitor.snMonitorInfos[id] = v
			}
		} else {
			monitor.snMonitorInfos[id] = MonitorInfo{StateStop, 0, time.Now().Unix()}
		}
		monitor.snLock.Unlock()
	}

	for _, info = range infos {
		id := info.Id.Int64()
		monitor.snLock.Lock()
		v, ok := monitor.snMonitorInfos[id]
		monitor.snLock.Unlock()
		if ok {
			log.Debug("collect-supernode-state", "id", id, "global-state", info.State, "local-state", v.curState, "missNum", v.missNum)
			if v.curState != info.State.Int64() {
				if v.curState == StateRunning || (v.curState == StateStop && v.missNum >= snMaxMissNum) {
					tempState, err := contract_api.GetSuperNodeUploadState(monitor.ctx, monitor.blockChainAPI, info.Id, from, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
					if err != nil {
						continue
					}
					if tempState.Int64() != v.curState {
						ids = append(ids, info.Id)
						states = append(states, big.NewInt(v.curState))
						monitor.snLock.Lock()
						delete(monitor.snMonitorInfos, id)
						monitor.snLock.Unlock()
						if len(ids) >= batchSize {
							break
						}
					}
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
	ret = enode[pos1+len("enode://") : pos2]
	return ret
}

func isPrivateIP(ip net.IP) bool {
	privateIPRange := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	for _, cidr := range privateIPRange {
		_, ipRange, _ := net.ParseCIDR(cidr)
		if ipRange.Contains(ip) {
			return true
		}
	}

	return false
}

func CheckPublicIP(url string) (bool, error) {
	node, err := enode.Parse(enode.ValidSchemes, url)
	if err != nil {
		return false, fmt.Errorf("invalid enode: %v", err)
	}

	ip := node.IP()
	if ip == nil {
		return false, fmt.Errorf("invalid ip")
	}

	privateIPFlag := isPrivateIP(ip)
	if privateIPFlag {
		return false, fmt.Errorf("ip is private")
	}

	return true, nil
}

// knownCache is a cache for known hashes.
type knownCache struct {
	hashes mapset.Set
	max    int
}

// newKnownCache creates a new knownCache with a max capacity.
func newKnownCache(max int) *knownCache {
	return &knownCache{
		max:    max,
		hashes: mapset.NewSet(),
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Add adds a list of elements to the set.
func (k *knownCache) Add(hashes ...common.Hash) {
	for k.hashes.Cardinality() > max(0, k.max-len(hashes)) {
		k.hashes.Pop()
	}
	for _, hash := range hashes {
		k.hashes.Add(hash)
	}
}

// Contains returns whether the given item is in the set.
func (k *knownCache) Contains(hash common.Hash) bool {
	return k.hashes.Contains(hash)
}

// Cardinality returns the number of elements in the set.
func (k *knownCache) Cardinality() int {
	return k.hashes.Cardinality()
}

func (monitor *NodeStateMonitor) GetTops(hash common.Hash) ([]common.Address, error) {
	return contract_api.GetTopSuperNodes(monitor.ctx, monitor.blockChainAPI, rpc.BlockNumberOrHashWithHash(hash, false))
}

func (monitor *NodeStateMonitor) GetSuperNodeInfo(snAddr common.Address, hash common.Hash) (*types.SuperNodeInfo, error) {
	return contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, snAddr, rpc.BlockNumberOrHashWithHash(hash, false))
}

func (monitor *NodeStateMonitor) ExistSuperNodeEnode(enode string, hash common.Hash) (bool, error) {
	return contract_api.ExistSuperNodeEnode(monitor.ctx, monitor.blockChainAPI, enode, rpc.BlockNumberOrHashWithHash(hash, false))
}

func (monitor *NodeStateMonitor) addSuperNodePeer() {
	if monitor.e.p2pServer == nil || monitor.e.blockchain == nil || monitor.blockChainAPI == nil {
		log.Trace("AddSuperNodePeer wait for prepare")
		return
	}

	hash := monitor.e.blockchain.CurrentBlock().Hash()
	if exist, _ := monitor.ExistSuperNodeEnode(monitor.enode, hash); !exist {
		flagSeedNode := false
		for _, url := range params.MainnetBootnodes {
			if contract_api.CompareEnode(monitor.enode, url) {
				flagSeedNode = true
				break
			}
		}
		if !flagSeedNode {
			return
		}

		peerCount := monitor.e.p2pServer.PeerCount()
		if peerCount != 0 {
			return
		}

		topAddrs, err := monitor.GetTops(hash)
		if err != nil {
			log.Error("node-state-monitor: addSuperNodePeer-1 get top supernodes failed", "hash", hash, "error", err)
			return
		}

		if len(topAddrs) == 0 {
			return
		}

		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(topAddrs))
		randomSuperNodeAddr := topAddrs[randomIndex]
		info, err := monitor.GetSuperNodeInfo(randomSuperNodeAddr, hash)
		if err != nil {
			log.Error("node-state-monitor: addSuperNodePeer-1 get supernode failed", "hash", hash, "error", err)
			return
		}

		node, err := enode.Parse(enode.ValidSchemes, info.Enode)
		if err != nil {
			log.Trace("node-state-monitor: invalid enode", "snAddr", info.Addr, "enode", info.Enode)
			return
		}

		monitor.e.p2pServer.AddPeer(node)
		return
	}

	topAddrs, err := monitor.GetTops(hash)
	if err != nil {
		log.Error("node-state-monitor: addSuperNodePeer-2 get top supernodes failed", "hash", hash, "error", err)
		return
	}

	for _, snAddr := range topAddrs {
		info, err := monitor.GetSuperNodeInfo(snAddr, hash)
		if err != nil {
			log.Error("node-state-monitor: addSuperNodePeer-2 get supernode failed", "hash", hash, "error", err)
			continue
		}

		if contract_api.CompareEnode(monitor.enode, info.Enode) {
			continue
		}

		node, err := enode.Parse(enode.ValidSchemes, info.Enode)
		if err != nil {
			log.Trace("node-state-monitor: invalid enode", "snAddr", info.Addr, "enode", info.Enode)
			continue
		}

		flag := false
		peersInfo := monitor.e.p2pServer.PeersInfo()
		for _, peer := range peersInfo {
			if contract_api.CompareEnode(peer.Enode, info.Enode) {
				flag = true
				break
			}
		}
		if !flag {
			monitor.e.p2pServer.AddPeer(node)
		} else {
			monitor.trustedLock.RLock()
			_, alreadyTrusted := monitor.trustedPeers[node.ID()]
			monitor.trustedLock.RUnlock()

			if !alreadyTrusted {
				monitor.e.p2pServer.AddTrustedPeer(node)
				monitor.trustedLock.Lock()
				monitor.trustedPeers[node.ID()] = struct{}{}
				monitor.trustedLock.Unlock()
			}
		}
	}
}

func (monitor *NodeStateMonitor) addSuperNodePeerLoop() {
	ticker := time.NewTicker(addSnDuration * time.Second)
	defer ticker.Stop()
	defer monitor.wg.Done()
	for {
		select {
		case <-monitor.addSnStopCh:
			log.Info("Exit node-state-monitor addSuperNodeLoop")
			return
		case <-ticker.C:
			monitor.addSuperNodePeer()
		}
	}
}
