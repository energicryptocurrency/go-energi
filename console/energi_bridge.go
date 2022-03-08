// Copyright 2020 The Energi Core Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package console

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/energicryptocurrency/energi/internal/jsre"
	"github.com/energicryptocurrency/energi/internal/web3ext"
	"github.com/energicryptocurrency/energi/log"

	"github.com/robertkrimen/otto"
)

type passwordPos int

const (
	passwordMask  = `*****`
	passwordRegex = `("|')(%v)("|')`

	// Maximum time the jsre can take to evaluate a command.
	runtimeTimeout = 2 * time.Second

	first passwordPos = iota
	last
)

// maskingJSRE uses a javascript runtime environment auto-loaded with web3.js and
// bignumber.js scripts so as to help resolve the functions passed as command
// arguments before obtaining the actual password in the specified cmd command.
type maskingJSRE struct {
	vmInstance *otto.Otto
	vmQuit     chan struct{}
}

type info struct {
	parent string
	index  passwordPos
}

var passwordIndexMap = map[string]info{
	"blacklistCollect":       {"energi", last},
	"blacklistDrain":         {"energi", last},
	"blacklistEnforce":       {"energi", last},
	"blacklistRevoke":        {"energi", last},
	"budgetPropose":          {"energi", last},
	"checkpointPropose":      {"energi", last},
	"claimGen2CoinsCombined": {"energi", first},
	"claimGen2CoinsDirect":   {"energi", first},
	"claimGen2CoinsImport":   {"energi", first},
	"upgradeCollect":         {"energi", last},
	"upgradePerform":         {"energi", last},
	"upgradePropose":         {"energi", last},
	"voteAccept":             {"energi", last},
	"voteReject":             {"energi", last},
	"withdrawFee":            {"energi", last},
	"announce":               {"masternode", last},
	"denounce":               {"masternode", last},
	"depositCollateral":      {"masternode", last},
	"withdrawCollateral":     {"masternode", last},
}

func autoCompileAndRun(vm *otto.Otto, filename string, src interface{}) (otto.Value, error) {
	script, err := vm.Compile(filename, src)
	if err != nil {
		return otto.Value{}, err
	}
	return vm.Run(script)
}

// firstPasswordIndex fetches the first argument that always the password.
func (m *maskingJSRE) firstPasswordIndex(call otto.FunctionCall) otto.Value {
	result, _ := m.vmInstance.ToValue(call.Argument(0).String())
	return result
}

// lastPasswordIndex fetches the last argument that always the password.
func (m *maskingJSRE) lastPasswordIndex(call otto.FunctionCall) otto.Value {
	c := len(call.ArgumentList)
	result, _ := m.vmInstance.ToValue(call.Argument(c - 1).String())
	return result
}

// newMaskingJSRE returns a new masking javascript runtime instance.
func newMaskingJSRE(
	send func(call otto.FunctionCall) otto.Value,
) (runtime *maskingJSRE, err error) {
	vm := otto.New()
	vm.Interrupt = make(chan func(), 1) // The buffer prevents blocking
	vm.Set("mask", struct{}{})

	maskObj, _ := vm.Get("mask")
	maskObj.Object().Set("send", send)
	maskObj.Object().Set("sendAsync", send)

	if _, err = autoCompileAndRun(vm, "bignumber.js", jsre.BigNumber_JS); err != nil {
		return
	}
	if _, err = autoCompileAndRun(vm, "web3.js", jsre.Web3_JS); err != nil {
		return
	}

	flatten := `var Web3 = require('web3'); var web3 = new Web3(mask); `
	if _, err = vm.Run(flatten); err != nil {
		return
	}

	flatten = ""
	for _, api := range []string{"eth", "admin", "personal"} {
		if file, ok := web3ext.Modules[api]; ok {
			if _, err = autoCompileAndRun(vm, fmt.Sprintf("%s.js", api), file); err != nil {
				return nil, err
			}
			flatten += fmt.Sprintf("var %s = web3.%s; ", api, api)
		} else if obj, err := vm.Run("web3." + api); err == nil && obj.IsObject() {
			flatten += fmt.Sprintf("var %s = web3.%s; ", api, api)
		}
	}

	if _, err = vm.Run(flatten); err != nil {
		return
	}

	energiObj, err := vm.Object("energi = {}")
	if err != nil {
		return
	}

	mnObj, err := vm.Object("masternode = {}")
	if err != nil {
		return
	}

	runtime = &maskingJSRE{
		vmInstance: vm,
		vmQuit:     make(chan struct{}),
	}

	// Set handler functions
	for key, info := range passwordIndexMap {
		funcCall := runtime.lastPasswordIndex
		if info.index == first {
			funcCall = runtime.firstPasswordIndex
		}
		switch info.parent {
		case "energi":
			energiObj.Set(key, funcCall)
		case "masternode":
			mnObj.Set(key, funcCall)
		default:
			err = fmt.Errorf("Unknown parent object instance: %v", info.parent)
			return
		}
	}
	return runtime, nil
}

// MaskPassword replaces the password in the provided command with a default
// mask password. This helps to protect the user info logged in the history file.
func (m *maskingJSRE) MaskPassword(command string) (string, error) {
	return m.replacePassword(command, passwordMask)
}

// UnMaskPassword replaces the mask password with the actual password.
func (m *maskingJSRE) UnMaskPassword(maskedCommand, pass string) (string, error) {
	// replace quotes if found.
	rg := regexp.MustCompile(`('|")`)
	pass = rg.ReplaceAllString(pass, "")

	return m.replacePassword(maskedCommand, pass)
}

func (m *maskingJSRE) vmRun(command string) (otto.Value, error) {
	go func() {
		select {
		case <-time.After(runtimeTimeout): // Stop after runtimeTimeout
			m.vmInstance.Interrupt <- func() {
				log.Debug("masking JSRE timeout on command: " + command)
			}
		case <-m.vmQuit: // Successfully completed command evaluation.
		}
	}()
	return m.vmInstance.Run(command)
}

// replacePassword puts the placeholder at the position where the password should be.
func (m *maskingJSRE) replacePassword(command, placeholder string) (string, error) {
	for key, val := range passwordIndexMap {
		searchTerm := fmt.Sprintf("%s.%s", val.parent, key)

		if strings.Contains(command, searchTerm) {

			// Use the JSRE to evaluate the actual password in the command.
			realPass, err := m.vmRun(command)
			if err != nil || !realPass.IsDefined() {
				return command, fmt.Errorf("Unable to locate passphrase in %s: %v", command, err)
			}

			m.vmQuit <- struct{}{}

			// escape any special characters.
			escapedPass := regexp.QuoteMeta(realPass.String())
			var re = regexp.MustCompile(fmt.Sprintf(passwordRegex, escapedPass))
			command = re.ReplaceAllString(command, `${1}`+placeholder+`${3}`)
			break
		}
	}
	return command, nil
}

// IsPasswordMasked checks if the password in the provided command has already been masked.
func (m *maskingJSRE) IsPasswordMasked(command string) bool {
	return strings.Contains(command, passwordMask)
}
