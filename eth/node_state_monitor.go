package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	StateRunning = iota + 1
	StateStop
)

const StateBroadcastDuration = 300
const StateUploadDuration = 37
const MaxMissNum = 10

const CoinbaseDuration = 30
const batchSize = 15

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

	broadcastStopCh chan struct{}
	uploadStopCh    chan struct{}
	coinbaseStopCh  chan struct{}

	wg   sync.WaitGroup
	lock sync.RWMutex

	mnMonitorInfos map[int64]MonitorInfo
	snMonitorInfos map[int64]MonitorInfo

	enode string
}

func newNodeStateMonitor() *NodeStateMonitor {
	monitor := &NodeStateMonitor{}
	monitor.ctx, monitor.cancelCtx = context.WithCancel(context.Background())
	monitor.mnMonitorInfos = make(map[int64]MonitorInfo)
	monitor.snMonitorInfos = make(map[int64]MonitorInfo)
	return monitor
}

func (monitor *NodeStateMonitor) Start(e *Ethereum) {
	monitor.e = e
	monitor.blockChainAPI = monitor.e.GetPublicBlockChainAPI()
	monitor.transactionPoolAPI = monitor.e.GetPublicTransactionPoolAPI()
	monitor.enode = contract_api.CompressEnode(monitor.e.p2pServer.NodeInfo().Enode)

	monitor.wg.Add(1)
	monitor.uploadStopCh = make(chan struct{})
	go monitor.uploadLoop()

	monitor.wg.Add(1)
	monitor.broadcastStopCh = make(chan struct{})
	go monitor.broadcastLoop()

	monitor.wg.Add(1)
	monitor.coinbaseStopCh = make(chan struct{})
	go monitor.coinbaseLoop()
}

func (monitor *NodeStateMonitor) Stop() {
	close(monitor.uploadStopCh)
	close(monitor.broadcastStopCh)
	close(monitor.coinbaseStopCh)
	monitor.wg.Wait()
	monitor.cancelCtx()
	log.Info("Node monitor stopped")
}

func (monitor *NodeStateMonitor) HandlePing(ping *types.NodePing) error {
	monitor.e.eventMux.Post(core.NodePingEvent{Ping: ping})

	if monitor.isSyncing() {
		return nil
	}

	localHeight := monitor.e.blockchain.CurrentBlock().NumberU64()
	remoteHeight := ping.CurHeight.Uint64()

	addr, err := monitor.e.Etherbase()
	if err == nil && monitor.isTopSuperNode(addr) {
		log.Debug("Handle received nodePing", "node-type", ping.NodeType, "node-id", ping.Id, "node-height", ping.CurHeight)

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

		curTime := time.Now().Unix()
		nodeType := ping.NodeType.Int64()
		if nodeType == int64(types.MasterNodeType) {
			info, err := contract_api.GetMasterNodeInfoByID(monitor.ctx, monitor.blockChainAPI, ping.Id, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
			if err != nil || hexutils.BytesToHex(pub)[1:] == GetPubKeyFromEnode(info.Enode) {
				return fmt.Errorf("invalid masternode ping, verify signature failed")
			}

			if _, err := CheckPublicIP(info.Enode); err != nil {
				return fmt.Errorf("invalid masternode ping, %s", err.Error())
			}

			monitor.lock.Lock()
			monitor.mnMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
			if ping.Version.Int64() < int64(types.NodePingVersion) || localHeight > remoteHeight + 20 {
				monitor.mnMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateStop, MaxMissNum, curTime}
			}
			monitor.lock.Unlock()
		} else if nodeType == int64(types.SuperNodeType) {
			info, err := contract_api.GetSuperNodeInfoByID(monitor.ctx, monitor.blockChainAPI, ping.Id, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
			if err != nil || hexutils.BytesToHex(pub)[1:] == GetPubKeyFromEnode(info.Enode) {
				return fmt.Errorf("invalid supernode ping, verify signature failed")
			}

			if _, err := CheckPublicIP(info.Enode); err != nil {
				return fmt.Errorf("invalid supernode ping, %s", err.Error())
			}

			monitor.lock.Lock()
			monitor.snMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
			if ping.Version.Int64() < int64(types.NodePingVersion) || localHeight > remoteHeight + 20 {
				monitor.snMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateStop, MaxMissNum, curTime}
			}
			monitor.lock.Unlock()
		}
	}
	return nil
}

func (monitor *NodeStateMonitor) uploadMNState(addr common.Address, ids []*big.Int, states []*big.Int) {
	hash, err := contract_api.UploadMasterNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, addr, ids, states)
	log.Info("Upload masternode state", "caller", addr, "ids", ids, "states", states, "hash", hash.Hex(), "error", err)
	//if err != nil {
	//	for i, id := range ids {
	//		if states[i].Int64() == StateRunning {
	//			monitor.mnMonitorInfos[id.Int64()] = MonitorInfo{StateRunning, 0, time.Now().Unix()}
	//		} else {
	//			monitor.mnMonitorInfos[id.Int64()] = MonitorInfo{StateStop, 5, time.Now().Unix()}
	//		}
	//	}
	//}
}

func (monitor *NodeStateMonitor) uploadSNState(addr common.Address, ids []*big.Int, states []*big.Int) {
	hash, err := contract_api.UploadSuperNodeStates(monitor.ctx, monitor.blockChainAPI, monitor.transactionPoolAPI, addr, ids, states)
	log.Info("Upload supernode state", "caller", addr, "ids", ids, "states", states, "hash", hash.Hex(), "error", err)
	//if err != nil {
	//	for i, id := range ids {
	//		if states[i].Int64() == StateRunning {
	//			monitor.snMonitorInfos[id.Int64()] = MonitorInfo{StateRunning, 0, time.Now().Unix()}
	//		} else {
	//			monitor.snMonitorInfos[id.Int64()] = MonitorInfo{StateStop, 5, time.Now().Unix()}
	//		}
	//	}
	//}
}

func (monitor *NodeStateMonitor) uploadLoop() {
	ticker := time.NewTicker(StateUploadDuration * time.Second)
	defer ticker.Stop()
	defer monitor.wg.Done()
	once := 1
	for {
		select {
		case <-monitor.uploadStopCh:
			log.Info("Exit node-state-monitor uploadLoop")
			return
		case <-ticker.C:
			if monitor.isSyncing() {
				if once == 1 {
					log.Info("Syncing now, wait a moment...")
					once = 0
				}
				continue
			}
			addr, err := monitor.e.Etherbase()
			if err != nil {
				continue
			}
			if monitor.isTopSuperNode(addr) {
				monitor.lock.Lock()
				mnIDs, mnStates := monitor.collectMasterNodes(addr)
				snIDs, snStates := monitor.collectSuperNodes(addr)
				monitor.lock.Unlock()

				if len(mnIDs) != 0 && len(mnIDs) == len(mnStates) {
					monitor.uploadMNState(addr, mnIDs, mnStates)
				}
				if len(snIDs) != 0 && len(snIDs) == len(snStates) {
					monitor.uploadSNState(addr, snIDs, snStates)
				}
			}
		}
	}
}

func (monitor *NodeStateMonitor) broadcastLoop() {
	ticker := time.NewTicker(StateBroadcastDuration * time.Second)
	defer ticker.Stop()
	defer monitor.wg.Done()
	lastAddr := common.Address{}
	once := 1
	for {
		select {
		case <-monitor.broadcastStopCh:
			log.Info("Exit node-state-monitor broadcastLoop")
			return
		case <-ticker.C:
			if monitor.isSyncing() {
				if once == 1 {
					log.Info("Syncing now, wait a moment...")
					once = 0
				}
				continue
			}

			addr, err := monitor.e.Etherbase()
			if err != nil {
				continue
			}
			curTime := time.Now().Unix()
			curBlock := monitor.e.blockchain.CurrentBlock()
			info1, err := contract_api.GetSuperNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
			if err == nil && info1.Id.Int64() != 0 { // supernode
				if !contract_api.CompareEnode(monitor.enode, info1.Enode) {
					if lastAddr != addr {
						log.Error("Broadcast supernode ping failed, incompatible enode", "local-enode", monitor.enode, "node-enode", info1.Enode)
						lastAddr = addr
					}
					continue
				}

				ping, err := types.NewNodePing(info1.Id, types.SuperNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
				if err != nil {
					if lastAddr != addr {
						log.Error("Broadcast supernode ping failed", "error", err)
						lastAddr = addr
					}
					continue
				}

				monitor.e.handler.BroadcastNodePing(ping)
				monitor.lock.Lock()
				monitor.snMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
				monitor.lock.Unlock()
				log.Debug("Broadcast supernode ping", "id", ping.Id, "height", ping.CurHeight)
				continue
			}
			info2, err := contract_api.GetMasterNodeInfo(monitor.ctx, monitor.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
			if err == nil && info2.Id.Int64() != 0 { // masternode
				if !contract_api.CompareEnode(monitor.enode, info2.Enode) {
					if lastAddr != addr {
						log.Error("Broadcast masternode ping failed, incompatible enode", "local-enode", monitor.enode, "node-enode", info1.Enode)
						lastAddr = addr
					}
					continue
				}

				ping, err := types.NewNodePing(info2.Id, types.MasterNodeType, curBlock.Hash(), curBlock.Number(), monitor.e.p2pServer.Config.PrivateKey)
				if err != nil {
					if lastAddr != addr {
						log.Error("Broadcast masternode ping failed", "error", err)
						lastAddr = addr
					}
					continue
				}

				monitor.e.handler.BroadcastNodePing(ping)
				monitor.lock.Lock()
				monitor.mnMonitorInfos[ping.Id.Int64()] = MonitorInfo{StateRunning, 0, curTime}
				monitor.lock.Unlock()
				log.Debug("Broadcast masternode ping", "id", ping.Id, "height", ping.CurHeight)
			}
		}
	}
}

func (monitor *NodeStateMonitor) coinbaseLoop() {
	ticker := time.NewTicker(CoinbaseDuration * time.Second)
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

	curTime := time.Now().Unix()
	var info types.MasterNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.mnMonitorInfos[id]; ok {
			if v.curState != StateRunning || curTime > v.lastTime+StateUploadDuration {
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
		if v, ok := monitor.mnMonitorInfos[id]; ok {
			log.Debug("collect-masternode-state", "id", id, "global-state", info.State, "local-state", v.curState, "missNum", v.missNum)
			if v.curState != info.State.Int64() {
				if v.curState == StateRunning || (v.curState == StateStop && v.missNum >= MaxMissNum) {
					flag := false
					entries, err := contract_api.GetMasterNodeUploadEntries(monitor.ctx, monitor.blockChainAPI, info.Id, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
					if err == nil {
						for _, entry := range entries {
							if entry.State.Int64() == v.curState && entry.Caller == from {
								flag = true
								break
							}
						}
					}
					if !flag {
						ids = append(ids, info.Id)
						states = append(states, big.NewInt(v.curState))
					}
					delete(monitor.mnMonitorInfos, id)
					if len(ids) == batchSize {
						break
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

	curTime := time.Now().Unix()
	var info types.SuperNodeInfo
	for _, info = range infos {
		id := info.Id.Int64()
		if v, ok := monitor.snMonitorInfos[id]; ok {
			if v.curState != StateRunning || curTime > v.lastTime+StateUploadDuration {
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
		if v, ok := monitor.snMonitorInfos[id]; ok {
			log.Debug("collect-supernode-state", "id", id, "global-state", info.State, "local-state", v.curState, "missNum", v.missNum)
			if v.curState != info.State.Int64() {
				if v.curState == StateRunning || (v.curState == StateStop && v.missNum >= MaxMissNum) {
					flag := false
					entries, err := contract_api.GetSuperNodeUploadEntries(monitor.ctx, monitor.blockChainAPI, info.Id, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
					if err == nil {
						for _, entry := range entries {
							if entry.State.Int64() == v.curState && entry.Caller == from {
								flag = true
								break
							}
						}
					}
					if !flag {
						ids = append(ids, info.Id)
						states = append(states, big.NewInt(v.curState))
					}
					delete(monitor.snMonitorInfos, id)
					if len(ids) == batchSize {
						break
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