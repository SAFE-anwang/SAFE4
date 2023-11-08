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
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/ethereum/go-ethereum/accounts/scwallet"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/ethereum/go-ethereum/internal/jsre"
	"github.com/ethereum/go-ethereum/rpc"
)

// bridge is a collection of JavaScript utility methods to bride the .js runtime
// environment and the Go RPC connection backing the remote method calls.
type bridge struct {
	client   *rpc.Client         // RPC client to execute Ethereum requests through
	prompter prompt.UserPrompter // Input prompter to allow interactive user feedback
	printer  io.Writer           // Output writer to serialize any display strings to
}

// newBridge creates a new JavaScript wrapper around an RPC client.
func newBridge(client *rpc.Client, prompter prompt.UserPrompter, printer io.Writer) *bridge {
	return &bridge{
		client:   client,
		prompter: prompter,
		printer:  printer,
	}
}

func getJeth(vm *goja.Runtime) *goja.Object {
	jeth := vm.Get("jeth")
	if jeth == nil {
		panic(vm.ToValue("jeth object does not exist"))
	}
	return jeth.ToObject(vm)
}

// NewAccount is a wrapper around the personal.newAccount RPC method that uses a
// non-echoing password prompt to acquire the passphrase and executes the original
// RPC method (saved in jeth.newAccount) with it to actually execute the RPC call.
func (b *bridge) NewAccount(call jsre.Call) (goja.Value, error) {
	var (
		password string
		confirm  string
		err      error
	)
	switch {
	// No password was specified, prompt the user for it
	case len(call.Arguments) == 0:
		if password, err = b.prompter.PromptPassword("Passphrase: "); err != nil {
			return nil, err
		}
		if confirm, err = b.prompter.PromptPassword("Repeat passphrase: "); err != nil {
			return nil, err
		}
		if password != confirm {
			return nil, fmt.Errorf("passwords don't match!")
		}
	// A single string password was specified, use that
	case len(call.Arguments) == 1 && call.Argument(0).ToString() != nil:
		password = call.Argument(0).ToString().String()
	default:
		return nil, fmt.Errorf("expected 0 or 1 string argument")
	}
	// Password acquired, execute the call and return
	newAccount, callable := goja.AssertFunction(getJeth(call.VM).Get("newAccount"))
	if !callable {
		return nil, fmt.Errorf("jeth.newAccount is not callable")
	}
	ret, err := newAccount(goja.Null(), call.VM.ToValue(password))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// OpenWallet is a wrapper around personal.openWallet which can interpret and
// react to certain error messages, such as the Trezor PIN matrix request.
func (b *bridge) OpenWallet(call jsre.Call) (goja.Value, error) {
	// Make sure we have a wallet specified to open
	if call.Argument(0).ToObject(call.VM).ClassName() != "String" {
		return nil, fmt.Errorf("first argument must be the wallet URL to open")
	}
	wallet := call.Argument(0)

	var passwd goja.Value
	if goja.IsUndefined(call.Argument(1)) || goja.IsNull(call.Argument(1)) {
		passwd = call.VM.ToValue("")
	} else {
		passwd = call.Argument(1)
	}
	// Open the wallet and return if successful in itself
	openWallet, callable := goja.AssertFunction(getJeth(call.VM).Get("openWallet"))
	if !callable {
		return nil, fmt.Errorf("jeth.openWallet is not callable")
	}
	val, err := openWallet(goja.Null(), wallet, passwd)
	if err == nil {
		return val, nil
	}

	// Wallet open failed, report error unless it's a PIN or PUK entry
	switch {
	case strings.HasSuffix(err.Error(), usbwallet.ErrTrezorPINNeeded.Error()):
		val, err = b.readPinAndReopenWallet(call)
		if err == nil {
			return val, nil
		}
		val, err = b.readPassphraseAndReopenWallet(call)
		if err != nil {
			return nil, err
		}

	case strings.HasSuffix(err.Error(), scwallet.ErrPairingPasswordNeeded.Error()):
		// PUK input requested, fetch from the user and call open again
		input, err := b.prompter.PromptPassword("Please enter the pairing password: ")
		if err != nil {
			return nil, err
		}
		passwd = call.VM.ToValue(input)
		if val, err = openWallet(goja.Null(), wallet, passwd); err != nil {
			if !strings.HasSuffix(err.Error(), scwallet.ErrPINNeeded.Error()) {
				return nil, err
			}
			// PIN input requested, fetch from the user and call open again
			input, err := b.prompter.PromptPassword("Please enter current PIN: ")
			if err != nil {
				return nil, err
			}
			if val, err = openWallet(goja.Null(), wallet, call.VM.ToValue(input)); err != nil {
				return nil, err
			}
		}

	case strings.HasSuffix(err.Error(), scwallet.ErrPINUnblockNeeded.Error()):
		// PIN unblock requested, fetch PUK and new PIN from the user
		var pukpin string
		input, err := b.prompter.PromptPassword("Please enter current PUK: ")
		if err != nil {
			return nil, err
		}
		pukpin = input
		input, err = b.prompter.PromptPassword("Please enter new PIN: ")
		if err != nil {
			return nil, err
		}
		pukpin += input

		if val, err = openWallet(goja.Null(), wallet, call.VM.ToValue(pukpin)); err != nil {
			return nil, err
		}

	case strings.HasSuffix(err.Error(), scwallet.ErrPINNeeded.Error()):
		// PIN input requested, fetch from the user and call open again
		input, err := b.prompter.PromptPassword("Please enter current PIN: ")
		if err != nil {
			return nil, err
		}
		if val, err = openWallet(goja.Null(), wallet, call.VM.ToValue(input)); err != nil {
			return nil, err
		}

	default:
		// Unknown error occurred, drop to the user
		return nil, err
	}
	return val, nil
}

func (b *bridge) readPassphraseAndReopenWallet(call jsre.Call) (goja.Value, error) {
	wallet := call.Argument(0)
	input, err := b.prompter.PromptPassword("Please enter your passphrase: ")
	if err != nil {
		return nil, err
	}
	openWallet, callable := goja.AssertFunction(getJeth(call.VM).Get("openWallet"))
	if !callable {
		return nil, fmt.Errorf("jeth.openWallet is not callable")
	}
	return openWallet(goja.Null(), wallet, call.VM.ToValue(input))
}

func (b *bridge) readPinAndReopenWallet(call jsre.Call) (goja.Value, error) {
	wallet := call.Argument(0)
	// Trezor PIN matrix input requested, display the matrix to the user and fetch the data
	fmt.Fprintf(b.printer, "Look at the device for number positions\n\n")
	fmt.Fprintf(b.printer, "7 | 8 | 9\n")
	fmt.Fprintf(b.printer, "--+---+--\n")
	fmt.Fprintf(b.printer, "4 | 5 | 6\n")
	fmt.Fprintf(b.printer, "--+---+--\n")
	fmt.Fprintf(b.printer, "1 | 2 | 3\n\n")

	input, err := b.prompter.PromptPassword("Please enter current PIN: ")
	if err != nil {
		return nil, err
	}
	openWallet, callable := goja.AssertFunction(getJeth(call.VM).Get("openWallet"))
	if !callable {
		return nil, fmt.Errorf("jeth.openWallet is not callable")
	}
	return openWallet(goja.Null(), wallet, call.VM.ToValue(input))
}

// UnlockAccount is a wrapper around the personal.unlockAccount RPC method that
// uses a non-echoing password prompt to acquire the passphrase and executes the
// original RPC method (saved in jeth.unlockAccount) with it to actually execute
// the RPC call.
func (b *bridge) UnlockAccount(call jsre.Call) (goja.Value, error) {
	if len(call.Arguments) < 1 {
		return nil, fmt.Errorf("usage: unlockAccount(account, [ password, duration ])")
	}

	account := call.Argument(0)
	// Make sure we have an account specified to unlock.
	if goja.IsUndefined(account) || goja.IsNull(account) || account.ExportType().Kind() != reflect.String {
		return nil, fmt.Errorf("first argument must be the account to unlock")
	}

	// If password is not given or is the null value, prompt the user for it.
	var passwd goja.Value
	if goja.IsUndefined(call.Argument(1)) || goja.IsNull(call.Argument(1)) {
		fmt.Fprintf(b.printer, "Unlock account %s\n", account)
		input, err := b.prompter.PromptPassword("Passphrase: ")
		if err != nil {
			return nil, err
		}
		passwd = call.VM.ToValue(input)
	} else {
		if call.Argument(1).ExportType().Kind() != reflect.String {
			return nil, fmt.Errorf("password must be a string")
		}
		passwd = call.Argument(1)
	}

	// Third argument is the duration how long the account should be unlocked.
	duration := goja.Null()
	if !goja.IsUndefined(call.Argument(2)) && !goja.IsNull(call.Argument(2)) {
		if !isNumber(call.Argument(2)) {
			return nil, fmt.Errorf("unlock duration must be a number")
		}
		duration = call.Argument(2)
	}

	// Send the request to the backend and return.
	unlockAccount, callable := goja.AssertFunction(getJeth(call.VM).Get("unlockAccount"))
	if !callable {
		return nil, fmt.Errorf("jeth.unlockAccount is not callable")
	}
	return unlockAccount(goja.Null(), account, passwd, duration)
}

// Sign is a wrapper around the personal.sign RPC method that uses a non-echoing password
// prompt to acquire the passphrase and executes the original RPC method (saved in
// jeth.sign) with it to actually execute the RPC call.
func (b *bridge) Sign(call jsre.Call) (goja.Value, error) {
	if nArgs := len(call.Arguments); nArgs < 2 {
		return nil, fmt.Errorf("usage: sign(message, account, [ password ])")
	}
	var (
		message = call.Argument(0)
		account = call.Argument(1)
		passwd  = call.Argument(2)
	)

	if goja.IsUndefined(message) || message.ExportType().Kind() != reflect.String {
		return nil, fmt.Errorf("first argument must be the message to sign")
	}
	if goja.IsUndefined(account) || account.ExportType().Kind() != reflect.String {
		return nil, fmt.Errorf("second argument must be the account to sign with")
	}

	// if the password is not given or null ask the user and ensure password is a string
	if goja.IsUndefined(passwd) || goja.IsNull(passwd) {
		fmt.Fprintf(b.printer, "Give password for account %s\n", account)
		input, err := b.prompter.PromptPassword("Password: ")
		if err != nil {
			return nil, err
		}
		passwd = call.VM.ToValue(input)
	} else if passwd.ExportType().Kind() != reflect.String {
		return nil, fmt.Errorf("third argument must be the password to unlock the account")
	}

	// Send the request to the backend and return
	sign, callable := goja.AssertFunction(getJeth(call.VM).Get("sign"))
	if !callable {
		return nil, fmt.Errorf("jeth.sign is not callable")
	}
	return sign(goja.Null(), message, account, passwd)
}

// Sleep will block the console for the specified number of seconds.
func (b *bridge) Sleep(call jsre.Call) (goja.Value, error) {
	if nArgs := len(call.Arguments); nArgs < 1 {
		return nil, fmt.Errorf("usage: sleep(<number of seconds>)")
	}
	sleepObj := call.Argument(0)
	if goja.IsUndefined(sleepObj) || goja.IsNull(sleepObj) || !isNumber(sleepObj) {
		return nil, fmt.Errorf("usage: sleep(<number of seconds>)")
	}
	sleep := sleepObj.ToFloat()
	time.Sleep(time.Duration(sleep * float64(time.Second)))
	return call.VM.ToValue(true), nil
}

// SleepBlocks will block the console for a specified number of new blocks optionally
// until the given timeout is reached.
func (b *bridge) SleepBlocks(call jsre.Call) (goja.Value, error) {
	// Parse the input parameters for the sleep.
	var (
		blocks = int64(0)
		sleep  = int64(9999999999999999) // indefinitely
	)
	nArgs := len(call.Arguments)
	if nArgs == 0 {
		return nil, fmt.Errorf("usage: sleepBlocks(<n blocks>[, max sleep in seconds])")
	}
	if nArgs >= 1 {
		if goja.IsNull(call.Argument(0)) || goja.IsUndefined(call.Argument(0)) || !isNumber(call.Argument(0)) {
			return nil, fmt.Errorf("expected number as first argument")
		}
		blocks = call.Argument(0).ToInteger()
	}
	if nArgs >= 2 {
		if goja.IsNull(call.Argument(1)) || goja.IsUndefined(call.Argument(1)) || !isNumber(call.Argument(1)) {
			return nil, fmt.Errorf("expected number as second argument")
		}
		sleep = call.Argument(1).ToInteger()
	}

	// Poll the current block number until either it or a timeout is reached.
	deadline := time.Now().Add(time.Duration(sleep) * time.Second)
	var lastNumber hexutil.Uint64
	if err := b.client.Call(&lastNumber, "eth_blockNumber"); err != nil {
		return nil, err
	}
	for time.Now().Before(deadline) {
		var number hexutil.Uint64
		if err := b.client.Call(&number, "eth_blockNumber"); err != nil {
			return nil, err
		}
		if number != lastNumber {
			lastNumber = number
			blocks--
		}
		if blocks <= 0 {
			break
		}
		time.Sleep(time.Second)
	}
	return call.VM.ToValue(true), nil
}

type jsonrpcCall struct {
	ID     int64
	Method string
	Params []interface{}
}

// Send implements the web3 provider "send" method.
func (b *bridge) Send(call jsre.Call) (goja.Value, error) {
	// Remarshal the request into a Go value.
	reqVal, err := call.Argument(0).ToObject(call.VM).MarshalJSON()
	if err != nil {
		return nil, err
	}

	var (
		rawReq = string(reqVal)
		dec    = json.NewDecoder(strings.NewReader(rawReq))
		reqs   []jsonrpcCall
		batch  bool
	)
	dec.UseNumber() // avoid float64s
	if rawReq[0] == '[' {
		batch = true
		dec.Decode(&reqs)
	} else {
		batch = false
		reqs = make([]jsonrpcCall, 1)
		dec.Decode(&reqs[0])
	}

	// Execute the requests.
	var resps []*goja.Object
	for _, req := range reqs {
		resp := call.VM.NewObject()
		resp.Set("jsonrpc", "2.0")
		resp.Set("id", req.ID)

		var result json.RawMessage
		if err = b.client.Call(&result, req.Method, req.Params...); err == nil {
			if result == nil {
				// Special case null because it is decoded as an empty
				// raw message for some reason.
				resp.Set("result", goja.Null())
			} else {
				JSON := call.VM.Get("JSON").ToObject(call.VM)
				parse, callable := goja.AssertFunction(JSON.Get("parse"))
				if !callable {
					return nil, fmt.Errorf("JSON.parse is not a function")
				}
				resultVal, err := parse(goja.Null(), call.VM.ToValue(string(result)))
				if err != nil {
					setError(resp, -32603, err.Error(), nil)
				} else {
					resp.Set("result", resultVal)
				}
			}
		} else {
			code := -32603
			var data interface{}
			if err, ok := err.(rpc.Error); ok {
				code = err.ErrorCode()
			}
			if err, ok := err.(rpc.DataError); ok {
				data = err.ErrorData()
			}
			setError(resp, code, err.Error(), data)
		}
		resps = append(resps, resp)
	}
	// Return the responses either to the callback (if supplied)
	// or directly as the return value.
	var result goja.Value
	if batch {
		result = call.VM.ToValue(resps)
	} else {
		result = resps[0]
	}
	if fn, isFunc := goja.AssertFunction(call.Argument(1)); isFunc {
		fn(goja.Null(), goja.Null(), result)
		return goja.Undefined(), nil
	}
	return result, nil
}

func setError(resp *goja.Object, code int, msg string, data interface{}) {
	err := make(map[string]interface{})
	err["code"] = code
	err["message"] = msg
	if data != nil {
		err["data"] = data
	}
	resp.Set("error", err)
}

// isNumber returns true if input value is a JS number.
func isNumber(v goja.Value) bool {
	k := v.ExportType().Kind()
	return k >= reflect.Int && k <= reflect.Float64
}

func getObject(vm *goja.Runtime, name string) *goja.Object {
	v := vm.Get(name)
	if v == nil {
		return nil
	}
	return v.ToObject(vm)
}

//Add get the public and private key of the specified account
func (b *bridge) GetPublicAndPrivateKey(call jsre.Call) (goja.Value, error) {
	if len(call.Arguments) < 1 {
		return nil, fmt.Errorf("usage: getpublicandprivatekey(account, [ password ])")
	}

	account := call.Argument(0)
	if goja.IsUndefined(account) || account.ExportType().Kind() != reflect.String {
		return nil, fmt.Errorf("first argument must be the account")
	}

	// If password is not given or is the null value, prompt the user for it.
	var passwd goja.Value
	if goja.IsUndefined(call.Argument(1)) || goja.IsNull(call.Argument(1)) {
		fmt.Fprintf(b.printer, "Unlock account %s\n", account)
		input, err := b.prompter.PromptPassword("Passphrase: ")
		if err != nil {
			return nil, err
		}
		passwd = call.VM.ToValue(input)
	} else {
		if call.Argument(1).ExportType().Kind() != reflect.String {
			return nil, fmt.Errorf("password must be a string")
		}
		passwd = call.Argument(1)
	}

	// Send the request to the backend and return.
	getPublicAndPrivateKey, callable := goja.AssertFunction(getJeth(call.VM).Get("getPublicAndPrivateKey"))
	if !callable {
		return nil, fmt.Errorf("jeth.getPublicAndPrivateKey is not callable")
	}
	return getPublicAndPrivateKey(goja.Null(), account, passwd)
}

/****************************** property ******************************/

func (b *bridge) AddProperty(call jsre.Call) (goja.Value, error) {
	addProperty, callable := goja.AssertFunction(getJeth(call.VM).Get("addProperty"))
	if !callable {
		return nil, fmt.Errorf("jeth.addProperty is not callable")
	}
	return addProperty(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) ApplyUpdateProperty(call jsre.Call) (goja.Value, error) {
	applyUpdateProperty, callable := goja.AssertFunction(getJeth(call.VM).Get("applyUpdateProperty"))
	if !callable {
		return nil, fmt.Errorf("jeth.applyUpdateProperty is not callable")
	}
	return applyUpdateProperty(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) Vote4UpdateProperty(call jsre.Call) (goja.Value, error) {
	vote4UpdateProperty, callable := goja.AssertFunction(getJeth(call.VM).Get("vote4UpdateProperty"))
	if !callable {
		return nil, fmt.Errorf("jeth.vote4UpdateProperty is not callable")
	}
	return vote4UpdateProperty(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) GetPropertyInfo(call jsre.Call) (goja.Value, error) {
	getPropertyInfo, callable := goja.AssertFunction(getJeth(call.VM).Get("getPropertyInfo"))
	if !callable {
		return nil, fmt.Errorf("jeth.getPropertyInfo is not callable")
	}
	return getPropertyInfo(goja.Null(), call.Argument(0))
}

func (b *bridge) GetUnconfirmedPropertyInfo(call jsre.Call) (goja.Value, error) {
	getUnconfirmedPropertyInfo, callable := goja.AssertFunction(getJeth(call.VM).Get("getUnconfirmedPropertyInfo"))
	if !callable {
		return nil, fmt.Errorf("jeth.getUnconfirmedPropertyInfo is not callable")
	}
	return getUnconfirmedPropertyInfo(goja.Null(), call.Argument(0))
}

func (b *bridge) GetPropertyValue(call jsre.Call) (goja.Value, error) {
	getPropertyValue, callable := goja.AssertFunction(getJeth(call.VM).Get("getPropertyValue"))
	if !callable {
		return nil, fmt.Errorf("jeth.getPropertyValue is not callable")
	}
	return getPropertyValue(goja.Null(), call.Argument(0))
}

func (b *bridge) GetAllProperties(call jsre.Call) (goja.Value, error) {
	getAllProperties, callable := goja.AssertFunction(getJeth(call.VM).Get("getAllProperties"))
	if !callable {
		return nil, fmt.Errorf("jeth.getAllProperties is not callable")
	}
	return getAllProperties(goja.Null())
}

func (b *bridge) GetAllUnconfirmedProperties(call jsre.Call) (goja.Value, error) {
	getAllUnconfirmedProperties, callable := goja.AssertFunction(getJeth(call.VM).Get("getAllUnconfirmedProperties"))
	if !callable {
		return nil, fmt.Errorf("jeth.getAllUnconfirmedProperties is not callable")
	}
	return getAllUnconfirmedProperties(goja.Null())
}

func (b *bridge) ExistProperty(call jsre.Call) (goja.Value, error) {
	existProperty, callable := goja.AssertFunction(getJeth(call.VM).Get("existProperty"))
	if !callable {
		return nil, fmt.Errorf("jeth.existProperty is not callable")
	}
	return existProperty(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistUnconfirmedProperty(call jsre.Call) (goja.Value, error) {
	existUnconfirmedProperty, callable := goja.AssertFunction(getJeth(call.VM).Get("existUnconfirmedProperty"))
	if !callable {
		return nil, fmt.Errorf("jeth.existUnconfirmedProperty is not callable")
	}
	return existUnconfirmedProperty(goja.Null(), call.Argument(0))
}

/****************************** account_manager ******************************/

func (b *bridge) Deposit(call jsre.Call) (goja.Value, error) {
	deposit, callable := goja.AssertFunction(getJeth(call.VM).Get("deposit"))
	if !callable {
		return nil, fmt.Errorf("jeth.deposit is not callable")
	}
	return deposit(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) Withdraw(call jsre.Call) (goja.Value, error) {
	withdraw, callable := goja.AssertFunction(getJeth(call.VM).Get("withdraw"))
	if !callable {
		return nil, fmt.Errorf("jeth.withdraw is not callable")
	}
	return withdraw(goja.Null(), call.Argument(0))
}

func (b *bridge) WithdrawByID(call jsre.Call) (goja.Value, error) {
	withdrawByID, callable := goja.AssertFunction(getJeth(call.VM).Get("withdrawByID"))
	if !callable {
		return nil, fmt.Errorf("jeth.withdrawByID is not callable")
	}
	return withdrawByID(goja.Null(), call.Argument(0), call.Argument(1))
}

func (b *bridge) Transfer(call jsre.Call) (goja.Value, error) {
	transfer, callable := goja.AssertFunction(getJeth(call.VM).Get("transfer"))
	if !callable {
		return nil, fmt.Errorf("jeth.transfer is not callable")
	}
	return transfer(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) AddLockDay(call jsre.Call) (goja.Value, error) {
	addLockDay, callable := goja.AssertFunction(getJeth(call.VM).Get("addLockDay"))
	if !callable {
		return nil, fmt.Errorf("jeth.addLockDay is not callable")
	}
	return addLockDay(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) GetTotalAmount(call jsre.Call) (goja.Value, error) {
	getTotalAmount, callable := goja.AssertFunction(getJeth(call.VM).Get("getTotalAmount"))
	if !callable {
		return nil, fmt.Errorf("jeth.getTotalAmount is not callable")
	}
	return getTotalAmount(goja.Null(), call.Argument(0))
}

func (b *bridge) GetAvailableAmount(call jsre.Call) (goja.Value, error) {
	getAvailableAmount, callable := goja.AssertFunction(getJeth(call.VM).Get("getAvailableAmount"))
	if !callable {
		return nil, fmt.Errorf("jeth.getAvailableAmount is not callable")
	}
	return getAvailableAmount(goja.Null(), call.Argument(0))
}

func (b *bridge) GetLockedAmount(call jsre.Call) (goja.Value, error) {
	getLockedAmount, callable := goja.AssertFunction(getJeth(call.VM).Get("getLockedAmount"))
	if !callable {
		return nil, fmt.Errorf("jeth.getLockedAmount is not callable")
	}
	return getLockedAmount(goja.Null(), call.Argument(0))
}

func (b *bridge) GetUsedAmount(call jsre.Call) (goja.Value, error) {
	getUsedAmount, callable := goja.AssertFunction(getJeth(call.VM).Get("getUsedAmount"))
	if !callable {
		return nil, fmt.Errorf("jeth.getUsedAmount is not callable")
	}
	return getUsedAmount(goja.Null(), call.Argument(0))
}

func (b *bridge) GetRecords(call jsre.Call) (goja.Value, error) {
	getRecords, callable := goja.AssertFunction(getJeth(call.VM).Get("getRecords"))
	if !callable {
		return nil, fmt.Errorf("jeth.getRecords is not callable")
	}
	return getRecords(goja.Null(), call.Argument(0))
}

func (b *bridge) GetRecord0(call jsre.Call) (goja.Value, error) {
	getRecord0, callable := goja.AssertFunction(getJeth(call.VM).Get("getRecord0"))
	if !callable {
		return nil, fmt.Errorf("jeth.getRecord0 is not callable")
	}
	return getRecord0(goja.Null(), call.Argument(0))
}

func (b *bridge) GetRecordByID(call jsre.Call) (goja.Value, error) {
	getRecordByID, callable := goja.AssertFunction(getJeth(call.VM).Get("getRecordByID"))
	if !callable {
		return nil, fmt.Errorf("jeth.getRecordByID is not callable")
	}
	return getRecordByID(goja.Null(), call.Argument(0))
}

func (b *bridge) GetRecordUseInfo(call jsre.Call) (goja.Value, error) {
	getRecordUseInfo, callable := goja.AssertFunction(getJeth(call.VM).Get("getRecordUseInfo"))
	if !callable {
		return nil, fmt.Errorf("jeth.getRecordUseInfo is not callable")
	}
	return getRecordUseInfo(goja.Null(), call.Argument(0))
}

/****************************** masternode ******************************/

func (b *bridge) StartMasterNode(call jsre.Call) (goja.Value, error) {
	startMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("startMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.startMasterNode is not callable")
	}
	return startMasterNode(goja.Null(), call.Argument(0))
}

func (b *bridge) StopMasterNode(call jsre.Call) (goja.Value, error) {
	stopMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("stopMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.stopMasterNode is not callable")
	}
	return stopMasterNode(goja.Null(), call.Argument(0))
}

func (b *bridge) RestartMasterNode(call jsre.Call) (goja.Value, error) {
	restartMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("restartMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.restartMasterNode is not callable")
	}
	return restartMasterNode(goja.Null(), call.Argument(0))
}

func (b *bridge) RegisterMasterNode(call jsre.Call) (goja.Value, error) {
	registerMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("registerMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.registerMasterNode is not callable")
	}
	return registerMasterNode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3), call.Argument(4), call.Argument(5), call.Argument(6), call.Argument(7), call.Argument(8))
}

func (b *bridge) AppendRegisterMasterNode(call jsre.Call) (goja.Value, error) {
	appendRegisterMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("appendRegisterMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.appendRegisterMasterNode is not callable")
	}
	return appendRegisterMasterNode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) TurnRegisterMasterNode(call jsre.Call) (goja.Value, error) {
	turnRegisterMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("turnRegisterMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.turnRegisterMasterNode is not callable")
	}
	return turnRegisterMasterNode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeMasterNodeAddress(call jsre.Call) (goja.Value, error) {
	changeMasterNodeAddress, callable := goja.AssertFunction(getJeth(call.VM).Get("changeMasterNodeAddress"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeMasterNodeAddress is not callable")
	}
	return changeMasterNodeAddress(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeMasterNodeEnode(call jsre.Call) (goja.Value, error) {
	changeMasterNodeEnode, callable := goja.AssertFunction(getJeth(call.VM).Get("changeMasterNodeEnode"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeMasterNodeEnode is not callable")
	}
	return changeMasterNodeEnode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeMasterNodeDescription(call jsre.Call) (goja.Value, error) {
	changeMasterNodeDescription, callable := goja.AssertFunction(getJeth(call.VM).Get("changeMasterNodeDescription"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeMasterNodeDescription is not callable")
	}
	return changeMasterNodeDescription(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeMasterNodeIsOfficial(call jsre.Call) (goja.Value, error) {
	changeMasterNodeIsOfficial, callable := goja.AssertFunction(getJeth(call.VM).Get("changeMasterNodeIsOfficial"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeMasterNodeIsOfficial is not callable")
	}
	return changeMasterNodeIsOfficial(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) GetMasterNodeInfo(call jsre.Call) (goja.Value, error) {
	getMasterNodeInfo, callable := goja.AssertFunction(getJeth(call.VM).Get("getMasterNodeInfo"))
	if !callable {
		return nil, fmt.Errorf("jeth.getMasterNodeInfo is not callable")
	}
	return getMasterNodeInfo(goja.Null(), call.Argument(0))
}

func (b *bridge) GetMasterNodeInfoByID(call jsre.Call) (goja.Value, error) {
	getMasterNodeInfoByID, callable := goja.AssertFunction(getJeth(call.VM).Get("getMasterNodeInfoByID"))
	if !callable {
		return nil, fmt.Errorf("jeth.getMasterNodeInfoByID is not callable")
	}
	return getMasterNodeInfoByID(goja.Null(), call.Argument(0))
}

func (b *bridge) GetNextMasterNode(call jsre.Call) (goja.Value, error) {
	getNextMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("getNextMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.getNextMasterNode is not callable")
	}
	return getNextMasterNode(goja.Null())
}

func (b *bridge) GetAllMasterNodes(call jsre.Call) (goja.Value, error) {
	getAllMasterNodes, callable := goja.AssertFunction(getJeth(call.VM).Get("getAllMasterNodes"))
	if !callable {
		return nil, fmt.Errorf("jeth.getAllMasterNodes is not callable")
	}
	return getAllMasterNodes(goja.Null())
}

func (b *bridge) GetOfficialMasterNodes(call jsre.Call) (goja.Value, error) {
	getOfficialMasterNodes, callable := goja.AssertFunction(getJeth(call.VM).Get("getOfficialMasterNodes"))
	if !callable {
		return nil, fmt.Errorf("jeth.getOfficialMasterNodes is not callable")
	}
	return getOfficialMasterNodes(goja.Null())
}

func (b *bridge) GetMasterNodeNum(call jsre.Call) (goja.Value, error) {
	getMasterNodeNum, callable := goja.AssertFunction(getJeth(call.VM).Get("getMasterNodeNum"))
	if !callable {
		return nil, fmt.Errorf("jeth.getMasterNodeNum is not callable")
	}
	return getMasterNodeNum(goja.Null())
}

func (b *bridge) ExistMasterNode(call jsre.Call) (goja.Value, error) {
	existMasterNode, callable := goja.AssertFunction(getJeth(call.VM).Get("existMasterNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.existMasterNode is not callable")
	}
	return existMasterNode(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistMasterNodeID(call jsre.Call) (goja.Value, error) {
	existMasterNodeID, callable := goja.AssertFunction(getJeth(call.VM).Get("existMasterNodeID"))
	if !callable {
		return nil, fmt.Errorf("jeth.existMasterNodeID is not callable")
	}
	return existMasterNodeID(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistMasterNodeEnode(call jsre.Call) (goja.Value, error) {
	existMasterNodeEnode, callable := goja.AssertFunction(getJeth(call.VM).Get("existMasterNodeEnode"))
	if !callable {
		return nil, fmt.Errorf("jeth.existMasterNodeEnode is not callable")
	}
	return existMasterNodeEnode(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistMasterNodeLockID(call jsre.Call) (goja.Value, error) {
	existMasterNodeLockID, callable := goja.AssertFunction(getJeth(call.VM).Get("existMasterNodeLockID"))
	if !callable {
		return nil, fmt.Errorf("jeth.existMasterNodeLockID is not callable")
	}
	return existMasterNodeLockID(goja.Null(), call.Argument(0))
}

/****************************** supernode ******************************/

func (b *bridge) StartSuperNode(call jsre.Call) (goja.Value, error) {
	startSuperNode, callable := goja.AssertFunction(getJeth(call.VM).Get("startSuperNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.startSuperNode is not callable")
	}
	return startSuperNode(goja.Null(), call.Argument(0))
}

func (b *bridge) StopSuperNode(call jsre.Call) (goja.Value, error) {
	stopSuperNode, callable := goja.AssertFunction(getJeth(call.VM).Get("stopSuperNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.stopSuperNode is not callable")
	}
	return stopSuperNode(goja.Null(), call.Argument(0))
}

func (b *bridge) RestartSuperNode(call jsre.Call) (goja.Value, error) {
	restartSuperNode, callable := goja.AssertFunction(getJeth(call.VM).Get("restartSuperNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.restartSuperNode is not callable")
	}
	return restartSuperNode(goja.Null(), call.Argument(0))
}

func (b *bridge) RegisterSuperNode(call jsre.Call) (goja.Value, error) {
	registerSuperNode, callable := goja.AssertFunction(getJeth(call.VM).Get("registerSuperNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.registerSuperNode is not callable")
	}
	return registerSuperNode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3), call.Argument(4), call.Argument(5), call.Argument(6), call.Argument(7), call.Argument(8), call.Argument(9), call.Argument(10))
}

func (b *bridge) AppendRegisterSuperNode(call jsre.Call) (goja.Value, error) {
	appendRegisterSuperNode, callable := goja.AssertFunction(getJeth(call.VM).Get("appendRegisterSuperNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.appendRegisterSuperNode is not callable")
	}
	return appendRegisterSuperNode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) TurnRegisterSuperNode(call jsre.Call) (goja.Value, error) {
	turnRegisterSuperNode, callable := goja.AssertFunction(getJeth(call.VM).Get("turnRegisterSuperNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.turnRegisterSuperNode is not callable")
	}
	return turnRegisterSuperNode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeSuperNodeAddress(call jsre.Call) (goja.Value, error) {
	changeSuperNodeAddress, callable := goja.AssertFunction(getJeth(call.VM).Get("changeSuperNodeAddress"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeSuperNodeAddress is not callable")
	}
	return changeSuperNodeAddress(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeSuperNodeName(call jsre.Call) (goja.Value, error) {
	changeSuperNodeName, callable := goja.AssertFunction(getJeth(call.VM).Get("changeSuperNodeName"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeSuperNodeName is not callable")
	}
	return changeSuperNodeName(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeSuperNodeEnode(call jsre.Call) (goja.Value, error) {
	changeSuperNodeEnode, callable := goja.AssertFunction(getJeth(call.VM).Get("changeSuperNodeEnode"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeSuperNodeEnode is not callable")
	}
	return changeSuperNodeEnode(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeSuperNodeDescription(call jsre.Call) (goja.Value, error) {
	changeSuperNodeDescription, callable := goja.AssertFunction(getJeth(call.VM).Get("changeSuperNodeDescription"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeSuperNodeDescription is not callable")
	}
	return changeSuperNodeDescription(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeSuperNodeIsOfficial(call jsre.Call) (goja.Value, error) {
	changeSuperNodeIsOfficial, callable := goja.AssertFunction(getJeth(call.VM).Get("changeSuperNodeIsOfficial"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeSuperNodeIsOfficial is not callable")
	}
	return changeSuperNodeIsOfficial(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) GetSuperNodeInfo(call jsre.Call) (goja.Value, error) {
	getSuperNodeInfo, callable := goja.AssertFunction(getJeth(call.VM).Get("getSuperNodeInfo"))
	if !callable {
		return nil, fmt.Errorf("jeth.getSuperNodeInfo is not callable")
	}
	return getSuperNodeInfo(goja.Null(), call.Argument(0))
}

func (b *bridge) GetSuperNodeInfoByID(call jsre.Call) (goja.Value, error) {
	getSuperNodeInfoByID, callable := goja.AssertFunction(getJeth(call.VM).Get("getSuperNodeInfoByID"))
	if !callable {
		return nil, fmt.Errorf("jeth.getSuperNodeInfoByID is not callable")
	}
	return getSuperNodeInfoByID(goja.Null(), call.Argument(0))
}

func (b *bridge) GetAllSuperNodes(call jsre.Call) (goja.Value, error) {
	getAllSuperNodes, callable := goja.AssertFunction(getJeth(call.VM).Get("getAllSuperNodes"))
	if !callable {
		return nil, fmt.Errorf("jeth.getAllSuperNodes is not callable")
	}
	return getAllSuperNodes(goja.Null())
}

func (b *bridge) GetTopSuperNodes(call jsre.Call) (goja.Value, error) {
	getTopSuperNodes, callable := goja.AssertFunction(getJeth(call.VM).Get("getTopSuperNodes"))
	if !callable {
		return nil, fmt.Errorf("jeth.getTopSuperNodes is not callable")
	}
	return getTopSuperNodes(goja.Null())
}

func (b *bridge) GetOfficialSuperNodes(call jsre.Call) (goja.Value, error) {
	getOfficialSuperNodes, callable := goja.AssertFunction(getJeth(call.VM).Get("getOfficialSuperNodes"))
	if !callable {
		return nil, fmt.Errorf("jeth.getOfficialSuperNodes is not callable")
	}
	return getOfficialSuperNodes(goja.Null())
}

func (b *bridge) GetSuperNodeNum(call jsre.Call) (goja.Value, error) {
	getSuperNodeNum, callable := goja.AssertFunction(getJeth(call.VM).Get("getSuperNodeNum"))
	if !callable {
		return nil, fmt.Errorf("jeth.getSuperNodeNum is not callable")
	}
	return getSuperNodeNum(goja.Null())
}

func (b *bridge) ExistSuperNode(call jsre.Call) (goja.Value, error) {
	existSuperNode, callable := goja.AssertFunction(getJeth(call.VM).Get("existSuperNode"))
	if !callable {
		return nil, fmt.Errorf("jeth.existSuperNode is not callable")
	}
	return existSuperNode(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistSuperNodeID(call jsre.Call) (goja.Value, error) {
	existSuperNodeID, callable := goja.AssertFunction(getJeth(call.VM).Get("existSuperNodeID"))
	if !callable {
		return nil, fmt.Errorf("jeth.existSuperNodeID is not callable")
	}
	return existSuperNodeID(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistSuperNodeName(call jsre.Call) (goja.Value, error) {
	existSuperNodeName, callable := goja.AssertFunction(getJeth(call.VM).Get("existSuperNodeName"))
	if !callable {
		return nil, fmt.Errorf("jeth.existSuperNodeName is not callable")
	}
	return existSuperNodeName(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistSuperNodeEnode(call jsre.Call) (goja.Value, error) {
	existSuperNodeEnode, callable := goja.AssertFunction(getJeth(call.VM).Get("existSuperNodeEnode"))
	if !callable {
		return nil, fmt.Errorf("jeth.existSuperNodeEnode is not callable")
	}
	return existSuperNodeEnode(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistSuperNodeLockID(call jsre.Call) (goja.Value, error) {
	existSuperNodeLockID, callable := goja.AssertFunction(getJeth(call.VM).Get("existSuperNodeLockID"))
	if !callable {
		return nil, fmt.Errorf("jeth.existSuperNodeLockID is not callable")
	}
	return existSuperNodeLockID(goja.Null(), call.Argument(0))
}

/****************************** snvote ******************************/

func (b *bridge) VoteOrApproval(call jsre.Call) (goja.Value, error) {
	voteOrApproval, callable := goja.AssertFunction(getJeth(call.VM).Get("voteOrApproval"))
	if !callable {
		return nil, fmt.Errorf("jeth.voteOrApproval is not callable")
	}
	return voteOrApproval(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3))
}

func (b *bridge) RemoveVoteOrApproval(call jsre.Call) (goja.Value, error) {
	removeVoteOrApproval, callable := goja.AssertFunction(getJeth(call.VM).Get("removeVoteOrApproval"))
	if !callable {
		return nil, fmt.Errorf("jeth.removeVoteOrApproval is not callable")
	}
	return removeVoteOrApproval(goja.Null(), call.Argument(0), call.Argument(1))
}

func (b *bridge) ProxyVote(call jsre.Call) (goja.Value, error) {
	proxyVote, callable := goja.AssertFunction(getJeth(call.VM).Get("proxyVote"))
	if !callable {
		return nil, fmt.Errorf("jeth.proxyVote is not callable")
	}
	return proxyVote(goja.Null(), call.Argument(0), call.Argument(1))
}

func (b *bridge) GetSuperNodes4Voter(call jsre.Call) (goja.Value, error) {
	getSuperNodes4Voter, callable := goja.AssertFunction(getJeth(call.VM).Get("getSuperNodes4Voter"))
	if !callable {
		return nil, fmt.Errorf("jeth.getSuperNodes4Voter is not callable")
	}
	return getSuperNodes4Voter(goja.Null(), call.Argument(0))
}

func (b *bridge) GetRecordIDs4Voter(call jsre.Call) (goja.Value, error) {
	getRecordIDs4Voter, callable := goja.AssertFunction(getJeth(call.VM).Get("getRecordIDs4Voter"))
	if !callable {
		return nil, fmt.Errorf("jeth.getRecordIDs4Voter is not callable")
	}
	return getRecordIDs4Voter(goja.Null(), call.Argument(0))
}

func (b *bridge) GetVoters4SN(call jsre.Call) (goja.Value, error) {
	getVoters4SN, callable := goja.AssertFunction(getJeth(call.VM).Get("getVoters4SN"))
	if !callable {
		return nil, fmt.Errorf("jeth.getVoters4SN is not callable")
	}
	return getVoters4SN(goja.Null(), call.Argument(0))
}

func (b *bridge) GetVoteNum4SN(call jsre.Call) (goja.Value, error) {
	getVoteNum4SN, callable := goja.AssertFunction(getJeth(call.VM).Get("getVoteNum4SN"))
	if !callable {
		return nil, fmt.Errorf("jeth.getVoteNum4SN is not callable")
	}
	return getVoteNum4SN(goja.Null(), call.Argument(0))
}

func (b *bridge) GetProxies4Voter(call jsre.Call) (goja.Value, error) {
	getProxies4Voter, callable := goja.AssertFunction(getJeth(call.VM).Get("getProxies4Voter"))
	if !callable {
		return nil, fmt.Errorf("jeth.getProxies4Voter is not callable")
	}
	return getProxies4Voter(goja.Null(), call.Argument(0))
}

func (b *bridge) GetProxiedRecordIDs4Voter(call jsre.Call) (goja.Value, error) {
	getProxiedRecordIDs4Voter, callable := goja.AssertFunction(getJeth(call.VM).Get("getProxiedRecordIDs4Voter"))
	if !callable {
		return nil, fmt.Errorf("jeth.getProxiedRecordIDs4Voter is not callable")
	}
	return getProxiedRecordIDs4Voter(goja.Null(), call.Argument(0))
}

func (b *bridge) GetVoters4Proxy(call jsre.Call) (goja.Value, error) {
	getVoters4Proxy, callable := goja.AssertFunction(getJeth(call.VM).Get("getVoters4Proxy"))
	if !callable {
		return nil, fmt.Errorf("jeth.getVoters4Proxy is not callable")
	}
	return getVoters4Proxy(goja.Null(), call.Argument(0))
}

func (b *bridge) GetVoteNum4Proxy(call jsre.Call) (goja.Value, error) {
	getVoteNum4Proxy, callable := goja.AssertFunction(getJeth(call.VM).Get("getVoteNum4Proxy"))
	if !callable {
		return nil, fmt.Errorf("jeth.getVoteNum4Proxy is not callable")
	}
	return getVoteNum4Proxy(goja.Null(), call.Argument(0))
}

/****************************** proposal ******************************/

func (b *bridge) CreateProposal(call jsre.Call) (goja.Value, error) {
	createProposal, callable := goja.AssertFunction(getJeth(call.VM).Get("createProposal"))
	if !callable {
		return nil, fmt.Errorf("jeth.createProposal is not callable")
	}
	return createProposal(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2), call.Argument(3), call.Argument(4), call.Argument(5), call.Argument(6))
}

func (b *bridge) Vote4Proposal(call jsre.Call) (goja.Value, error) {
	vote4Proposal, callable := goja.AssertFunction(getJeth(call.VM).Get("vote4Proposal"))
	if !callable {
		return nil, fmt.Errorf("jeth.vote4Proposal is not callable")
	}
	return vote4Proposal(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeProposalTitle(call jsre.Call) (goja.Value, error) {
	changeProposalTitle, callable := goja.AssertFunction(getJeth(call.VM).Get("changeProposalTitle"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeProposalTitle is not callable")
	}
	return changeProposalTitle(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeProposalPayAmount(call jsre.Call) (goja.Value, error) {
	changeProposalPayAmount, callable := goja.AssertFunction(getJeth(call.VM).Get("changeProposalPayAmount"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeProposalPayAmount is not callable")
	}
	return changeProposalPayAmount(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeProposalPayTimes(call jsre.Call) (goja.Value, error) {
	changeProposalPayTimes, callable := goja.AssertFunction(getJeth(call.VM).Get("changeProposalPayTimes"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeProposalPayTimes is not callable")
	}
	return changeProposalPayTimes(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeProposalStartPayTime(call jsre.Call) (goja.Value, error) {
	changeProposalStartPayTime, callable := goja.AssertFunction(getJeth(call.VM).Get("changeProposalStartPayTime"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeProposalStartPayTime is not callable")
	}
	return changeProposalStartPayTime(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeProposalEndPayTime(call jsre.Call) (goja.Value, error) {
	changeProposalEndPayTime, callable := goja.AssertFunction(getJeth(call.VM).Get("changeProposalEndPayTime"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeProposalEndPayTime is not callable")
	}
	return changeProposalEndPayTime(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) ChangeProposalDescription(call jsre.Call) (goja.Value, error) {
	changeProposalDescription, callable := goja.AssertFunction(getJeth(call.VM).Get("changeProposalDescription"))
	if !callable {
		return nil, fmt.Errorf("jeth.changeProposalDescription is not callable")
	}
	return changeProposalDescription(goja.Null(), call.Argument(0), call.Argument(1), call.Argument(2))
}

func (b *bridge) GetProposalInfo(call jsre.Call) (goja.Value, error) {
	getProposalInfo, callable := goja.AssertFunction(getJeth(call.VM).Get("getProposalInfo"))
	if !callable {
		return nil, fmt.Errorf("jeth.getProposalInfo is not callable")
	}
	return getProposalInfo(goja.Null(), call.Argument(0))
}

func (b *bridge) GetAllProposals(call jsre.Call) (goja.Value, error) {
	getAllProposals, callable := goja.AssertFunction(getJeth(call.VM).Get("getAllProposals"))
	if !callable {
		return nil, fmt.Errorf("jeth.getAllProposals is not callable")
	}
	return getAllProposals(goja.Null())
}

func (b *bridge) GetMineProposals(call jsre.Call) (goja.Value, error) {
	getMineProposals, callable := goja.AssertFunction(getJeth(call.VM).Get("getMineProposals"))
	if !callable {
		return nil, fmt.Errorf("jeth.getMineProposals is not callable")
	}
	return getMineProposals(goja.Null(), call.Argument(0))
}

func (b *bridge) ExistProposal(call jsre.Call) (goja.Value, error) {
	existProposal, callable := goja.AssertFunction(getJeth(call.VM).Get("existProposal"))
	if !callable {
		return nil, fmt.Errorf("jeth.existProposal is not callable")
	}
	return existProposal(goja.Null(), call.Argument(0))
}
