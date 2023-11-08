// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package console

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"syscall"

	"github.com/dop251/goja"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/ethereum/go-ethereum/internal/jsre"
	"github.com/ethereum/go-ethereum/internal/jsre/deps"
	"github.com/ethereum/go-ethereum/internal/web3ext"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mattn/go-colorable"
	"github.com/peterh/liner"
)

var (
	// u: unlock, s: signXX, sendXX, n: newAccount, i: importXX
	passwordRegexp = regexp.MustCompile(`personal.[nusi]`)
	onlyWhitespace = regexp.MustCompile(`^\s*$`)
	exit           = regexp.MustCompile(`^\s*exit\s*;*\s*$`)
)

// HistoryFile is the file within the data directory to store input scrollback.
const HistoryFile = "history"

// DefaultPrompt is the default prompt line prefix to use for user input querying.
const DefaultPrompt = "> "

// Config is the collection of configurations to fine tune the behavior of the
// JavaScript console.
type Config struct {
	DataDir  string              // Data directory to store the console history at
	DocRoot  string              // Filesystem path from where to load JavaScript files from
	Client   *rpc.Client         // RPC client to execute Ethereum requests through
	Prompt   string              // Input prompt prefix string (defaults to DefaultPrompt)
	Prompter prompt.UserPrompter // Input prompter to allow interactive user feedback (defaults to TerminalPrompter)
	Printer  io.Writer           // Output writer to serialize any display strings to (defaults to os.Stdout)
	Preload  []string            // Absolute paths to JavaScript files to preload
}

// Console is a JavaScript interpreted runtime environment. It is a fully fledged
// JavaScript console attached to a running node via an external or in-process RPC
// client.
type Console struct {
	client   *rpc.Client         // RPC client to execute Ethereum requests through
	jsre     *jsre.JSRE          // JavaScript runtime environment running the interpreter
	prompt   string              // Input prompt prefix string
	prompter prompt.UserPrompter // Input prompter to allow interactive user feedback
	histPath string              // Absolute path to the console scrollback history
	history  []string            // Scroll history maintained by the console
	printer  io.Writer           // Output writer to serialize any display strings to

	interactiveStopped chan struct{}
	stopInteractiveCh  chan struct{}
	signalReceived     chan struct{}
	stopped            chan struct{}
	wg                 sync.WaitGroup
	stopOnce           sync.Once
}

// New initializes a JavaScript interpreted runtime environment and sets defaults
// with the config struct.
func New(config Config) (*Console, error) {
	// Handle unset config values gracefully
	if config.Prompter == nil {
		config.Prompter = prompt.Stdin
	}
	if config.Prompt == "" {
		config.Prompt = DefaultPrompt
	}
	if config.Printer == nil {
		config.Printer = colorable.NewColorableStdout()
	}

	// Initialize the console and return
	console := &Console{
		client:             config.Client,
		jsre:               jsre.New(config.DocRoot, config.Printer),
		prompt:             config.Prompt,
		prompter:           config.Prompter,
		printer:            config.Printer,
		histPath:           filepath.Join(config.DataDir, HistoryFile),
		interactiveStopped: make(chan struct{}),
		stopInteractiveCh:  make(chan struct{}),
		signalReceived:     make(chan struct{}, 1),
		stopped:            make(chan struct{}),
	}
	if err := os.MkdirAll(config.DataDir, 0700); err != nil {
		return nil, err
	}
	if err := console.init(config.Preload); err != nil {
		return nil, err
	}

	console.wg.Add(1)
	go console.interruptHandler()

	return console, nil
}

// init retrieves the available APIs from the remote RPC provider and initializes
// the console's JavaScript namespaces based on the exposed modules.
func (c *Console) init(preload []string) error {
	c.initConsoleObject()

	// Initialize the JavaScript <-> Go RPC bridge.
	bridge := newBridge(c.client, c.prompter, c.printer)
	if err := c.initWeb3(bridge); err != nil {
		return err
	}
	if err := c.initExtensions(); err != nil {
		return err
	}

	// Add bridge overrides for web3.js functionality.
	c.jsre.Do(func(vm *goja.Runtime) {
		c.initAdmin(vm, bridge)
		c.initPersonal(vm, bridge)
		c.initSysProperty(vm, bridge)
		c.initAccount(vm, bridge)
		c.initMasterNode(vm, bridge)
		c.initSuperNode(vm, bridge)
		c.initSNVote(vm, bridge)
		c.initProposal(vm, bridge)
	})

	// Preload JavaScript files.
	for _, path := range preload {
		if err := c.jsre.Exec(path); err != nil {
			failure := err.Error()
			if gojaErr, ok := err.(*goja.Exception); ok {
				failure = gojaErr.String()
			}
			return fmt.Errorf("%s: %v", path, failure)
		}
	}

	// Configure the input prompter for history and tab completion.
	if c.prompter != nil {
		if content, err := os.ReadFile(c.histPath); err != nil {
			c.prompter.SetHistory(nil)
		} else {
			c.history = strings.Split(string(content), "\n")
			c.prompter.SetHistory(c.history)
		}
		c.prompter.SetWordCompleter(c.AutoCompleteInput)
	}
	return nil
}

func (c *Console) initConsoleObject() {
	c.jsre.Do(func(vm *goja.Runtime) {
		console := vm.NewObject()
		console.Set("log", c.consoleOutput)
		console.Set("error", c.consoleOutput)
		vm.Set("console", console)
	})
}

func (c *Console) initWeb3(bridge *bridge) error {
	if err := c.jsre.Compile("bignumber.js", deps.BigNumberJS); err != nil {
		return fmt.Errorf("bignumber.js: %v", err)
	}
	if err := c.jsre.Compile("web3.js", deps.Web3JS); err != nil {
		return fmt.Errorf("web3.js: %v", err)
	}
	if _, err := c.jsre.Run("var Web3 = require('web3');"); err != nil {
		return fmt.Errorf("web3 require: %v", err)
	}
	var err error
	c.jsre.Do(func(vm *goja.Runtime) {
		transport := vm.NewObject()
		transport.Set("send", jsre.MakeCallback(vm, bridge.Send))
		transport.Set("sendAsync", jsre.MakeCallback(vm, bridge.Send))
		vm.Set("_consoleWeb3Transport", transport)
		_, err = vm.RunString("var web3 = new Web3(_consoleWeb3Transport)")
	})
	return err
}

// initExtensions loads and registers web3.js extensions.
func (c *Console) initExtensions() error {
	// Compute aliases from server-provided modules.
	apis, err := c.client.SupportedModules()
	if err != nil {
		return fmt.Errorf("api modules: %v", err)
	}
	aliases := map[string]struct{}{"eth": {}, "personal": {}}
	for api := range apis {
		if api == "web3" {
			continue
		}
		aliases[api] = struct{}{}
		if file, ok := web3ext.Modules[api]; ok {
			if err = c.jsre.Compile(api+".js", file); err != nil {
				return fmt.Errorf("%s.js: %v", api, err)
			}
		}
	}

	// Apply aliases.
	c.jsre.Do(func(vm *goja.Runtime) {
		web3 := getObject(vm, "web3")
		for name := range aliases {
			if v := web3.Get(name); v != nil {
				vm.Set(name, v)
			}
		}
	})
	return nil
}

// initAdmin creates additional admin APIs implemented by the bridge.
func (c *Console) initAdmin(vm *goja.Runtime, bridge *bridge) {
	if admin := getObject(vm, "admin"); admin != nil {
		admin.Set("sleepBlocks", jsre.MakeCallback(vm, bridge.SleepBlocks))
		admin.Set("sleep", jsre.MakeCallback(vm, bridge.Sleep))
		admin.Set("clearHistory", c.clearHistory)
	}
}

// initPersonal redirects account-related API methods through the bridge.
//
// If the console is in interactive mode and the 'personal' API is available, override
// the openWallet, unlockAccount, newAccount and sign methods since these require user
// interaction. The original web3 callbacks are stored in 'jeth'. These will be called
// by the bridge after the prompt and send the original web3 request to the backend.
func (c *Console) initPersonal(vm *goja.Runtime, bridge *bridge) {
	personal := getObject(vm, "personal")
	if personal == nil || c.prompter == nil {
		return
	}
	jeth := vm.NewObject()
	vm.Set("jeth", jeth)
	jeth.Set("openWallet", personal.Get("openWallet"))
	jeth.Set("unlockAccount", personal.Get("unlockAccount"))
	jeth.Set("newAccount", personal.Get("newAccount"))
	jeth.Set("sign", personal.Get("sign"))

	//Add get the public and private key of the specified account
	jeth.Set("getPublicAndPrivateKey", personal.Get("getPublicAndPrivateKey"))

	personal.Set("openWallet", jsre.MakeCallback(vm, bridge.OpenWallet))
	personal.Set("unlockAccount", jsre.MakeCallback(vm, bridge.UnlockAccount))
	personal.Set("newAccount", jsre.MakeCallback(vm, bridge.NewAccount))
	personal.Set("sign", jsre.MakeCallback(vm, bridge.Sign))

	//Add get the public and private key of the specified account
	personal.Set("getPublicAndPrivateKey", jsre.MakeCallback(vm, bridge.GetPublicAndPrivateKey))
}

func (c *Console) initSysProperty(vm *goja.Runtime, bridge *bridge) {
	sysproperty := getObject(vm, "sysproperty")
	if sysproperty == nil || c.prompter == nil {
		return
	}

	getJeth(vm).Set("addProperty", sysproperty.Get("add"))
	getJeth(vm).Set("applyUpdateProperty", sysproperty.Get("applyUpdate"))
	getJeth(vm).Set("vote4UpdateProperty", sysproperty.Get("vote4Update"))
	getJeth(vm).Set("getPropertyInfo", sysproperty.Get("getInfo"))
	getJeth(vm).Set("getUnconfirmedPropertyInfo", sysproperty.Get("getUnconfirmedInfo"))
	getJeth(vm).Set("getPropertyValue", sysproperty.Get("getValue"))
	getJeth(vm).Set("getAllProperties", sysproperty.Get("getAll"))
	getJeth(vm).Set("getAllUnconfirmedProperties", sysproperty.Get("getAllUnconfirmed"))
	getJeth(vm).Set("existProperty", sysproperty.Get("exist"))
	getJeth(vm).Set("existUnconfirmedProperty", sysproperty.Get("existUnconfirmed"))

	sysproperty.Set("add", jsre.MakeCallback(vm, bridge.AddProperty))
	sysproperty.Set("applyUpdate", jsre.MakeCallback(vm, bridge.ApplyUpdateProperty))
	sysproperty.Set("vote4Update", jsre.MakeCallback(vm, bridge.Vote4UpdateProperty))
	sysproperty.Set("getInfo", jsre.MakeCallback(vm, bridge.GetPropertyInfo))
	sysproperty.Set("getUnconfirmedInfo", jsre.MakeCallback(vm, bridge.GetUnconfirmedPropertyInfo))
	sysproperty.Set("getValue", jsre.MakeCallback(vm, bridge.GetPropertyValue))
	sysproperty.Set("getAll", jsre.MakeCallback(vm, bridge.GetAllProperties))
	sysproperty.Set("getAllUnconfirmed", jsre.MakeCallback(vm, bridge.GetAllUnconfirmedProperties))
	sysproperty.Set("exist", jsre.MakeCallback(vm, bridge.ExistProperty))
	sysproperty.Set("existUnconfirmed", jsre.MakeCallback(vm, bridge.ExistUnconfirmedProperty))
}

func (c *Console) initAccount(vm *goja.Runtime, bridge *bridge) {
	account := getObject(vm, "account")
	if account == nil || c.prompter == nil {
		return
	}

	getJeth(vm).Set("deposit", account.Get("deposit"))
	getJeth(vm).Set("withdraw", account.Get("withdraw"))
	getJeth(vm).Set("withdrawByID", account.Get("withdrawByID"))
	getJeth(vm).Set("transfer", account.Get("transfer"))
	getJeth(vm).Set("addLockDay", account.Get("addLockDay"))
	getJeth(vm).Set("getTotalAmount", account.Get("getTotalAmount"))
	getJeth(vm).Set("getAvailableAmount", account.Get("getAvailableAmount"))
	getJeth(vm).Set("getLockedAmount", account.Get("getLockedAmount"))
	getJeth(vm).Set("getUsedAmount", account.Get("getUsedAmount"))
	getJeth(vm).Set("getRecords", account.Get("getRecords"))
	getJeth(vm).Set("getRecord0", account.Get("getRecord0"))
	getJeth(vm).Set("getRecordByID", account.Get("getRecordByID"))
	getJeth(vm).Set("getRecordUseInfo", account.Get("getRecordUseInfo"))

	account.Set("deposit", jsre.MakeCallback(vm, bridge.Deposit))
	account.Set("withdraw", jsre.MakeCallback(vm, bridge.Withdraw))
	account.Set("withdrawByID", jsre.MakeCallback(vm, bridge.WithdrawByID))
	account.Set("transfer", jsre.MakeCallback(vm, bridge.Transfer))
	account.Set("addLockDay", jsre.MakeCallback(vm, bridge.AddLockDay))
	account.Set("getTotalAmount", jsre.MakeCallback(vm, bridge.GetTotalAmount))
	account.Set("getAvailableAmount", jsre.MakeCallback(vm, bridge.GetAvailableAmount))
	account.Set("getLockedAmount", jsre.MakeCallback(vm, bridge.GetLockedAmount))
	account.Set("getUsedAmount", jsre.MakeCallback(vm, bridge.GetUsedAmount))
	account.Set("getRecords", jsre.MakeCallback(vm, bridge.GetRecords))
	account.Set("getRecord0", jsre.MakeCallback(vm, bridge.GetRecord0))
	account.Set("getRecordByID", jsre.MakeCallback(vm, bridge.GetRecordByID))
	account.Set("getRecordUseInfo", jsre.MakeCallback(vm, bridge.GetRecordUseInfo))
}

func (c *Console) initMasterNode(vm *goja.Runtime, bridge *bridge) {
	masternode := getObject(vm, "masternode")
	if masternode == nil || c.prompter == nil {
		return
	}

	getJeth(vm).Set("startMasterNode", masternode.Get("start"))
	getJeth(vm).Set("stopMasterNode", masternode.Get("stop"))
	getJeth(vm).Set("restartMasterNode", masternode.Get("restart"))
	getJeth(vm).Set("registerMasterNode", masternode.Get("register"))
	getJeth(vm).Set("appendRegisterMasterNode", masternode.Get("appendRegister"))
	getJeth(vm).Set("turnRegisterMasterNode", masternode.Get("turnRegister"))
	getJeth(vm).Set("changeMasterNodeAddress", masternode.Get("changeAddress"))
	getJeth(vm).Set("changeMasterNodeEnode", masternode.Get("changeEnode"))
	getJeth(vm).Set("changeMasterNodeDescription", masternode.Get("changeDescription"))
	getJeth(vm).Set("changeMasterNodeIsOfficial", masternode.Get("changeIsOfficial"))
	getJeth(vm).Set("getMasterNodeInfo", masternode.Get("getInfo"))
	getJeth(vm).Set("getMasterNodeInfoByID", masternode.Get("getInfoByID"))
	getJeth(vm).Set("getNextMasterNode", masternode.Get("getNext"))
	getJeth(vm).Set("getAllMasterNodes", masternode.Get("getAll"))
	getJeth(vm).Set("getOfficialMasterNodes", masternode.Get("getOfficials"))
	getJeth(vm).Set("getMasterNodeNum", masternode.Get("getNum"))
	getJeth(vm).Set("existMasterNode", masternode.Get("exist"))
	getJeth(vm).Set("existMasterNodeID", masternode.Get("existID"))
	getJeth(vm).Set("existMasterNodeEnode", masternode.Get("existEnode"))
	getJeth(vm).Set("existMasterNodeLockID", masternode.Get("existLockID"))

	masternode.Set("start", jsre.MakeCallback(vm, bridge.StartMasterNode))
	masternode.Set("stop", jsre.MakeCallback(vm, bridge.StopMasterNode))
	masternode.Set("restart", jsre.MakeCallback(vm, bridge.RestartMasterNode))
	masternode.Set("register", jsre.MakeCallback(vm, bridge.RegisterMasterNode))
	masternode.Set("appendRegister", jsre.MakeCallback(vm, bridge.AppendRegisterMasterNode))
	masternode.Set("turnRegister", jsre.MakeCallback(vm, bridge.TurnRegisterMasterNode))
	masternode.Set("changeAddress", jsre.MakeCallback(vm, bridge.ChangeMasterNodeAddress))
	masternode.Set("changeEnode", jsre.MakeCallback(vm, bridge.ChangeMasterNodeEnode))
	masternode.Set("changeDescription", jsre.MakeCallback(vm, bridge.ChangeMasterNodeDescription))
	masternode.Set("changeIsOfficial", jsre.MakeCallback(vm, bridge.ChangeMasterNodeIsOfficial))
	masternode.Set("getInfo", jsre.MakeCallback(vm, bridge.GetMasterNodeInfo))
	masternode.Set("getInfoByID", jsre.MakeCallback(vm, bridge.GetMasterNodeInfoByID))
	masternode.Set("getNext", jsre.MakeCallback(vm, bridge.GetNextMasterNode))
	masternode.Set("getAll", jsre.MakeCallback(vm, bridge.GetAllMasterNodes))
	masternode.Set("getOfficials", jsre.MakeCallback(vm, bridge.GetOfficialMasterNodes))
	masternode.Set("getNum", jsre.MakeCallback(vm, bridge.GetMasterNodeNum))
	masternode.Set("exist", jsre.MakeCallback(vm, bridge.ExistMasterNode))
	masternode.Set("existID", jsre.MakeCallback(vm, bridge.ExistMasterNodeID))
	masternode.Set("existEnode", jsre.MakeCallback(vm, bridge.ExistMasterNodeEnode))
	masternode.Set("existLockID", jsre.MakeCallback(vm, bridge.ExistMasterNodeLockID))
}

func (c *Console) initSuperNode(vm *goja.Runtime, bridge *bridge) {
	supernode := getObject(vm, "supernode")
	if supernode == nil || c.prompter == nil {
		return
	}

	getJeth(vm).Set("startSuperNode", supernode.Get("start"))
	getJeth(vm).Set("stopSuperNode", supernode.Get("stop"))
	getJeth(vm).Set("restartSuperNode", supernode.Get("restart"))
	getJeth(vm).Set("registerSuperNode", supernode.Get("register"))
	getJeth(vm).Set("appendRegisterSuperNode", supernode.Get("appendRegister"))
	getJeth(vm).Set("turnRegisterSuperNode", supernode.Get("turnRegister"))
	getJeth(vm).Set("changeSuperNodeAddress", supernode.Get("changeAddress"))
	getJeth(vm).Set("changeSuperNodeName", supernode.Get("changeName"))
	getJeth(vm).Set("changeSuperNodeEnode", supernode.Get("changeEnode"))
	getJeth(vm).Set("changeSuperNodeDescription", supernode.Get("changeDescription"))
	getJeth(vm).Set("changeSuperNodeIsOfficial", supernode.Get("changeIsOfficial"))
	getJeth(vm).Set("getSuperNodeInfo", supernode.Get("getInfo"))
	getJeth(vm).Set("getSuperNodeInfoByID", supernode.Get("getInfoByID"))
	getJeth(vm).Set("getAllSuperNodes", supernode.Get("getAll"))
	getJeth(vm).Set("getTopSuperNodes", supernode.Get("getTops"))
	getJeth(vm).Set("getOfficialSuperNodes", supernode.Get("getOfficials"))
	getJeth(vm).Set("getSuperNodeNum", supernode.Get("getNum"))
	getJeth(vm).Set("existSuperNode", supernode.Get("exist"))
	getJeth(vm).Set("existSuperNodeID", supernode.Get("existID"))
	getJeth(vm).Set("existSuperNodeName", supernode.Get("existName"))
	getJeth(vm).Set("existSuperNodeEnode", supernode.Get("existEnode"))
	getJeth(vm).Set("existSuperNodeLockID", supernode.Get("existLockID"))

	supernode.Set("start", jsre.MakeCallback(vm, bridge.StartSuperNode))
	supernode.Set("stop", jsre.MakeCallback(vm, bridge.StopSuperNode))
	supernode.Set("restart", jsre.MakeCallback(vm, bridge.RestartSuperNode))
	supernode.Set("register", jsre.MakeCallback(vm, bridge.RegisterSuperNode))
	supernode.Set("appendRegister", jsre.MakeCallback(vm, bridge.AppendRegisterSuperNode))
	supernode.Set("turnRegister", jsre.MakeCallback(vm, bridge.TurnRegisterSuperNode))
	supernode.Set("changeAddress", jsre.MakeCallback(vm, bridge.ChangeSuperNodeAddress))
	supernode.Set("changeName", jsre.MakeCallback(vm, bridge.ChangeSuperNodeName))
	supernode.Set("changeEnode", jsre.MakeCallback(vm, bridge.ChangeSuperNodeEnode))
	supernode.Set("changeDescription", jsre.MakeCallback(vm, bridge.ChangeSuperNodeDescription))
	supernode.Set("changeIsOfficial", jsre.MakeCallback(vm, bridge.ChangeSuperNodeIsOfficial))
	supernode.Set("getInfo", jsre.MakeCallback(vm, bridge.GetSuperNodeInfo))
	supernode.Set("getInfoByID", jsre.MakeCallback(vm, bridge.GetSuperNodeInfoByID))
	supernode.Set("getAll", jsre.MakeCallback(vm, bridge.GetAllSuperNodes))
	supernode.Set("getTops", jsre.MakeCallback(vm, bridge.GetTopSuperNodes))
	supernode.Set("getOfficials", jsre.MakeCallback(vm, bridge.GetOfficialSuperNodes))
	supernode.Set("getNum", jsre.MakeCallback(vm, bridge.GetSuperNodeNum))
	supernode.Set("exist", jsre.MakeCallback(vm, bridge.ExistSuperNode))
	supernode.Set("existID", jsre.MakeCallback(vm, bridge.ExistSuperNodeID))
	supernode.Set("existName", jsre.MakeCallback(vm, bridge.ExistSuperNodeName))
	supernode.Set("existEnode", jsre.MakeCallback(vm, bridge.ExistSuperNodeEnode))
	supernode.Set("existLockID", jsre.MakeCallback(vm, bridge.ExistSuperNodeLockID))
}

func (c *Console) initSNVote(vm *goja.Runtime, bridge *bridge) {
	snvote := getObject(vm, "snvote")
	if snvote == nil || c.prompter == nil {
		return
	}

	getJeth(vm).Set("voteOrApproval", snvote.Get("voteOrApproval"))
	getJeth(vm).Set("removeVoteOrApproval", snvote.Get("removeVoteOrApproval"))
	getJeth(vm).Set("proxyVote", snvote.Get("proxyVote"))
	getJeth(vm).Set("getSuperNodes4Voter", snvote.Get("getSuperNodes4Voter"))
	getJeth(vm).Set("getRecordIDs4Voter", snvote.Get("getRecordIDs4Voter"))
	getJeth(vm).Set("getVoters4SN", snvote.Get("getVoters4SN"))
	getJeth(vm).Set("getVoteNum4SN", snvote.Get("getVoteNum4SN"))
	getJeth(vm).Set("getProxies4Voter", snvote.Get("getProxies4Voter"))
	getJeth(vm).Set("getProxiedRecordIDs4Voter", snvote.Get("getProxiedRecordIDs4Voter"))
	getJeth(vm).Set("getVoters4Proxy", snvote.Get("getVoters4Proxy"))
	getJeth(vm).Set("getVoteNum4Proxy", snvote.Get("getVoteNum4Proxy"))

	snvote.Set("voteOrApproval", jsre.MakeCallback(vm, bridge.VoteOrApproval))
	snvote.Set("removeVoteOrApproval", jsre.MakeCallback(vm, bridge.RemoveVoteOrApproval))
	snvote.Set("proxyVote", jsre.MakeCallback(vm, bridge.ProxyVote))
	snvote.Set("getSuperNodes4Voter", jsre.MakeCallback(vm, bridge.GetSuperNodes4Voter))
	snvote.Set("getRecordIDs4Voter", jsre.MakeCallback(vm, bridge.GetRecordIDs4Voter))
	snvote.Set("getVoters4SN", jsre.MakeCallback(vm, bridge.GetVoters4SN))
	snvote.Set("getVoteNum4SN", jsre.MakeCallback(vm, bridge.GetVoteNum4SN))
	snvote.Set("getProxies4Voter", jsre.MakeCallback(vm, bridge.GetProxies4Voter))
	snvote.Set("getProxiedRecordIDs4Voter", jsre.MakeCallback(vm, bridge.GetProxiedRecordIDs4Voter))
	snvote.Set("getVoters4Proxy", jsre.MakeCallback(vm, bridge.GetVoters4Proxy))
	snvote.Set("getVoteNum4Proxy", jsre.MakeCallback(vm, bridge.GetVoteNum4Proxy))
}

func (c *Console) initProposal(vm *goja.Runtime, bridge *bridge) {
	proposal := getObject(vm, "proposal")
	if proposal == nil || c.prompter == nil {
		return
	}

	getJeth(vm).Set("createProposal", proposal.Get("create"))
	getJeth(vm).Set("vote4Proposal", proposal.Get("vote"))
	getJeth(vm).Set("changeProposalTitle", proposal.Get("changeTitle"))
	getJeth(vm).Set("changeProposalPayAmount", proposal.Get("changePayAmount"))
	getJeth(vm).Set("changeProposalPayTimes", proposal.Get("changePayTimes"))
	getJeth(vm).Set("changeProposalStartPayTime", proposal.Get("changeStartPayTime"))
	getJeth(vm).Set("changeProposalEndPayTime", proposal.Get("changeEndPayTime"))
	getJeth(vm).Set("changeProposalDescription", proposal.Get("changeDescription"))
	getJeth(vm).Set("getProposalInfo", proposal.Get("getInfo"))
	getJeth(vm).Set("getAllProposals", proposal.Get("getAll"))
	getJeth(vm).Set("getMineProposals", proposal.Get("getMines"))
	getJeth(vm).Set("existProposal", proposal.Get("exist"))

	proposal.Set("create", jsre.MakeCallback(vm, bridge.CreateProposal))
	proposal.Set("vote", jsre.MakeCallback(vm, bridge.Vote4Proposal))
	proposal.Set("changeTitle", jsre.MakeCallback(vm, bridge.ChangeProposalTitle))
	proposal.Set("changePayAmount", jsre.MakeCallback(vm, bridge.ChangeProposalPayAmount))
	proposal.Set("changePayTimes", jsre.MakeCallback(vm, bridge.ChangeProposalPayTimes))
	proposal.Set("changeStartPayTime", jsre.MakeCallback(vm, bridge.ChangeProposalStartPayTime))
	proposal.Set("changeEndPayTime", jsre.MakeCallback(vm, bridge.ChangeProposalEndPayTime))
	proposal.Set("changeDescription", jsre.MakeCallback(vm, bridge.ChangeProposalDescription))
	proposal.Set("getInfo", jsre.MakeCallback(vm, bridge.GetProposalInfo))
	proposal.Set("getAll", jsre.MakeCallback(vm, bridge.GetAllProposals))
	proposal.Set("getMines", jsre.MakeCallback(vm, bridge.GetMineProposals))
	proposal.Set("exist", jsre.MakeCallback(vm, bridge.ExistProposal))
}

func (c *Console) clearHistory() {
	c.history = nil
	c.prompter.ClearHistory()
	if err := os.Remove(c.histPath); err != nil {
		fmt.Fprintln(c.printer, "can't delete history file:", err)
	} else {
		fmt.Fprintln(c.printer, "history file deleted")
	}
}

// consoleOutput is an override for the console.log and console.error methods to
// stream the output into the configured output stream instead of stdout.
func (c *Console) consoleOutput(call goja.FunctionCall) goja.Value {
	var output []string
	for _, argument := range call.Arguments {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	fmt.Fprintln(c.printer, strings.Join(output, " "))
	return goja.Null()
}

// AutoCompleteInput is a pre-assembled word completer to be used by the user
// input prompter to provide hints to the user about the methods available.
func (c *Console) AutoCompleteInput(line string, pos int) (string, []string, string) {
	// No completions can be provided for empty inputs
	if len(line) == 0 || pos == 0 {
		return "", nil, ""
	}
	// Chunck data to relevant part for autocompletion
	// E.g. in case of nested lines eth.getBalance(eth.coinb<tab><tab>
	start := pos - 1
	for ; start > 0; start-- {
		// Skip all methods and namespaces (i.e. including the dot)
		if line[start] == '.' || (line[start] >= 'a' && line[start] <= 'z') || (line[start] >= 'A' && line[start] <= 'Z') {
			continue
		}
		// Handle web3 in a special way (i.e. other numbers aren't auto completed)
		if start >= 3 && line[start-3:start] == "web3" {
			start -= 3
			continue
		}
		// We've hit an unexpected character, autocomplete form here
		start++
		break
	}
	return line[:start], c.jsre.CompleteKeywords(line[start:pos]), line[pos:]
}

// Welcome show summary of current Geth instance and some metadata about the
// console's available modules.
func (c *Console) Welcome() {
	message := "Welcome to the Geth JavaScript console!\n\n"

	// Print some generic Geth metadata
	if res, err := c.jsre.Run(`
		var message = "instance: " + web3.version.node + "\n";
		try {
			message += "coinbase: " + eth.coinbase + "\n";
		} catch (err) {}
		message += "at block: " + eth.blockNumber + " (" + new Date(1000 * eth.getBlock(eth.blockNumber).timestamp) + ")\n";
		try {
			message += " datadir: " + admin.datadir + "\n";
		} catch (err) {}
		message
	`); err == nil {
		message += res.String()
	}
	// List all the supported modules for the user to call
	if apis, err := c.client.SupportedModules(); err == nil {
		modules := make([]string, 0, len(apis))
		for api, version := range apis {
			modules = append(modules, fmt.Sprintf("%s:%s", api, version))
		}
		sort.Strings(modules)
		message += " modules: " + strings.Join(modules, " ") + "\n"
	}
	message += "\nTo exit, press ctrl-d or type exit"
	fmt.Fprintln(c.printer, message)
}

// Evaluate executes code and pretty prints the result to the specified output
// stream.
func (c *Console) Evaluate(statement string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(c.printer, "[native] error: %v\n", r)
		}
	}()
	c.jsre.Evaluate(statement, c.printer)

	// Avoid exiting Interactive when jsre was interrupted by SIGINT.
	c.clearSignalReceived()
}

// interruptHandler runs in its own goroutine and waits for signals.
// When a signal is received, it interrupts the JS interpreter.
func (c *Console) interruptHandler() {
	defer c.wg.Done()

	// During Interactive, liner inhibits the signal while it is prompting for
	// input. However, the signal will be received while evaluating JS.
	//
	// On unsupported terminals, SIGINT can also happen while prompting.
	// Unfortunately, it is not possible to abort the prompt in this case and
	// the c.readLines goroutine leaks.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	defer signal.Stop(sig)

	for {
		select {
		case <-sig:
			c.setSignalReceived()
			c.jsre.Interrupt(errors.New("interrupted"))
		case <-c.stopInteractiveCh:
			close(c.interactiveStopped)
			c.jsre.Interrupt(errors.New("interrupted"))
		case <-c.stopped:
			return
		}
	}
}

func (c *Console) setSignalReceived() {
	select {
	case c.signalReceived <- struct{}{}:
	default:
	}
}

func (c *Console) clearSignalReceived() {
	select {
	case <-c.signalReceived:
	default:
	}
}

// StopInteractive causes Interactive to return as soon as possible.
func (c *Console) StopInteractive() {
	select {
	case c.stopInteractiveCh <- struct{}{}:
	case <-c.stopped:
	}
}

// Interactive starts an interactive user session, where in.put is propted from
// the configured user prompter.
func (c *Console) Interactive() {
	var (
		prompt      = c.prompt             // the current prompt line (used for multi-line inputs)
		indents     = 0                    // the current number of input indents (used for multi-line inputs)
		input       = ""                   // the current user input
		inputLine   = make(chan string, 1) // receives user input
		inputErr    = make(chan error, 1)  // receives liner errors
		requestLine = make(chan string)    // requests a line of input
	)

	defer func() {
		c.writeHistory()
	}()

	// The line reader runs in a separate goroutine.
	go c.readLines(inputLine, inputErr, requestLine)
	defer close(requestLine)

	for {
		// Send the next prompt, triggering an input read.
		requestLine <- prompt

		select {
		case <-c.interactiveStopped:
			fmt.Fprintln(c.printer, "node is down, exiting console")
			return

		case <-c.signalReceived:
			// SIGINT received while prompting for input -> unsupported terminal.
			// I'm not sure if the best choice would be to leave the console running here.
			// Bash keeps running in this case. node.js does not.
			fmt.Fprintln(c.printer, "caught interrupt, exiting")
			return

		case err := <-inputErr:
			if err == liner.ErrPromptAborted {
				// When prompting for multi-line input, the first Ctrl-C resets
				// the multi-line state.
				prompt, indents, input = c.prompt, 0, ""
				continue
			}
			return

		case line := <-inputLine:
			// User input was returned by the prompter, handle special cases.
			if indents <= 0 && exit.MatchString(line) {
				return
			}
			if onlyWhitespace.MatchString(line) {
				continue
			}
			// Append the line to the input and check for multi-line interpretation.
			input += line + "\n"
			indents = countIndents(input)
			if indents <= 0 {
				prompt = c.prompt
			} else {
				prompt = strings.Repeat(".", indents*3) + " "
			}
			// If all the needed lines are present, save the command and run it.
			if indents <= 0 {
				if len(input) > 0 && input[0] != ' ' && !passwordRegexp.MatchString(input) {
					if command := strings.TrimSpace(input); len(c.history) == 0 || command != c.history[len(c.history)-1] {
						c.history = append(c.history, command)
						if c.prompter != nil {
							c.prompter.AppendHistory(command)
						}
					}
				}
				c.Evaluate(input)
				input = ""
			}
		}
	}
}

// readLines runs in its own goroutine, prompting for input.
func (c *Console) readLines(input chan<- string, errc chan<- error, prompt <-chan string) {
	for p := range prompt {
		line, err := c.prompter.PromptInput(p)
		if err != nil {
			errc <- err
		} else {
			input <- line
		}
	}
}

// countIndents returns the number of identations for the given input.
// In case of invalid input such as var a = } the result can be negative.
func countIndents(input string) int {
	var (
		indents     = 0
		inString    = false
		strOpenChar = ' '   // keep track of the string open char to allow var str = "I'm ....";
		charEscaped = false // keep track if the previous char was the '\' char, allow var str = "abc\"def";
	)

	for _, c := range input {
		switch c {
		case '\\':
			// indicate next char as escaped when in string and previous char isn't escaping this backslash
			if !charEscaped && inString {
				charEscaped = true
			}
		case '\'', '"':
			if inString && !charEscaped && strOpenChar == c { // end string
				inString = false
			} else if !inString && !charEscaped { // begin string
				inString = true
				strOpenChar = c
			}
			charEscaped = false
		case '{', '(':
			if !inString { // ignore brackets when in string, allow var str = "a{"; without indenting
				indents++
			}
			charEscaped = false
		case '}', ')':
			if !inString {
				indents--
			}
			charEscaped = false
		default:
			charEscaped = false
		}
	}

	return indents
}

// Execute runs the JavaScript file specified as the argument.
func (c *Console) Execute(path string) error {
	return c.jsre.Exec(path)
}

// Stop cleans up the console and terminates the runtime environment.
func (c *Console) Stop(graceful bool) error {
	c.stopOnce.Do(func() {
		// Stop the interrupt handler.
		close(c.stopped)
		c.wg.Wait()
	})

	c.jsre.Stop(graceful)
	return nil
}

func (c *Console) writeHistory() error {
	if err := os.WriteFile(c.histPath, []byte(strings.Join(c.history, "\n")), 0600); err != nil {
		return err
	}
	return os.Chmod(c.histPath, 0600) // Force 0600, even if it was different previously
}
