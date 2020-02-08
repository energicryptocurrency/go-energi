package console

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/robertkrimen/otto"
)

// TestPasswordMasking tests the password masking and unmasking functionality.
func TestPasswordMasking(t *testing.T) {
	type testData struct {
		name     string // test name
		original string // UnMasked command
		masked   string
	}

	var td = []testData{
		{
			"announce",
			` masternode.announce('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 'enode://40865e11060ba79c3df7137e695ca452e4dbc5@127.0.0.1:49797', 'password') 
			`,
			` masternode.announce('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 'enode://40865e11060ba79c3df7137e695ca452e4dbc5@127.0.0.1:49797', '*****') 
			`,
		}, {
			"announce",
			` masternode.announce(eth.accounts[0], admin.nodeInfo.enode, 'password')`,
			` masternode.announce(eth.accounts[0], admin.nodeInfo.enode, '*****')`,
		}, {
			"depositCollateral",
			`masternode.depositCollateral('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', web3.toWei('1000', 'ether'), 'password')`,
			`masternode.depositCollateral('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', web3.toWei('1000', 'ether'), '*****')`,
		}, {
			"depositCollateral",
			`masternode.depositCollateral(eth.accounts[0], web3.toWei('1000', 'ether'), 'password')`,
			`masternode.depositCollateral(eth.accounts[0], web3.toWei('1000', 'ether'), '*****')`,
		}, {
			"withdrawCollateral",
			`masternode.withdrawCollateral('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', web3.toWei('1000', 'ether'), 'password')`,
			`masternode.withdrawCollateral('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', web3.toWei('1000', 'ether'), '*****')`,
		}, {
			"denounce",
			`masternode.denounce('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 'password')`,
			`masternode.denounce('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '*****')`,
		}, {
			"blacklistCollect",
			`energi.blacklistCollect('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x0000000000000000000067236482647824682', 'password')`,
			`energi.blacklistCollect('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x0000000000000000000067236482647824682', '*****')`,
		}, {
			"blacklistDrain",
			`energi.blacklistDrain('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 10000, '0x0000000000000000000067236482647824682', 'password')`,
			`energi.blacklistDrain('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 10000, '0x0000000000000000000067236482647824682', '*****')`,
		}, {
			"blacklistEnforce",
			`energi.blacklistEnforce('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 10000, '0x0000000000000000000067236482647824682', 'password')`,
			`energi.blacklistEnforce('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 10000, '0x0000000000000000000067236482647824682', '*****')`,
		}, {
			"blacklistRevoke",
			`energi.blacklistRevoke('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 10000, '0x0000000000000000000067236482647824682', 'password')`,
			`energi.blacklistRevoke('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 10000, '0x0000000000000000000067236482647824682', '*****')`,
		}, {
			"budgetPropose",
			`energi.budgetPropose(10000000000667700, '063b5fb2-7e2a-4292-acc7-8ece5dd1c530', 36000, '0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', 'password')`,
			`energi.budgetPropose(10000000000667700, '063b5fb2-7e2a-4292-acc7-8ece5dd1c530', 36000, '0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '*****')`,
		}, {
			"checkpointPropose",
			`energi.checkpointPropose(3467578, '0x2D0bc327d0843CAF6Fd9ae1eFaB0bF7196Fc2FC8', 'password')`,
			`energi.checkpointPropose(3467578, '0x2D0bc327d0843CAF6Fd9ae1eFaB0bF7196Fc2FC8', '*****')`,
		}, {
			"claimGen2CoinsCombined",
			`energi.claimGen2CoinsCombined('password', '0x2D0bc327d0843CAF6Fd9ae1eFaB0bF7196Fc2FC8', '/tmp/file/path.energi')`,
			`energi.claimGen2CoinsCombined('*****', '0x2D0bc327d0843CAF6Fd9ae1eFaB0bF7196Fc2FC8', '/tmp/file/path.energi')`,
		}, {
			"claimGen2CoinsDirect",
			`energi.claimGen2CoinsDirect('password', '0x2D0bc327d0843CAF6Fd9ae1eFaB0bF7196Fc2FC8', '@#1234556')`,
			`energi.claimGen2CoinsDirect('*****', '0x2D0bc327d0843CAF6Fd9ae1eFaB0bF7196Fc2FC8', '@#1234556')`,
		}, {
			"claimGen2CoinsImport",
			`energi.claimGen2CoinsImport('password', '/tmp/file/path.energi')`,
			`energi.claimGen2CoinsImport('*****', '/tmp/file/path.energi')`,
		}, {
			"upgradeCollect",
			`energi.upgradeCollect('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x00000000000000000000000000000000000', '0x0000000000000000000067236482647824682', 'password')`,
			`energi.upgradeCollect('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x00000000000000000000000000000000000', '0x0000000000000000000067236482647824682', '*****')`,
		}, {
			"upgradePropose",
			`energi.upgradePropose('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd', 2600, 1000000, '0x00000000000000000000000000000000000', '0x0000000000000000000067236482647824682', 'password')`,
			`energi.upgradePropose('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd', 2600, 1000000, '0x00000000000000000000000000000000000', '0x0000000000000000000067236482647824682', '*****')`,
		}, {
			"upgradePerform",
			`energi.upgradePerform('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd', 'password')`,
			`energi.upgradePerform('0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd', '*****')`,
		}, {
			"voteAccept",
			`energi.voteAccept('0x00000000000000000000000000000000000', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd','password')`,
			`energi.voteAccept('0x00000000000000000000000000000000000', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd','*****')`,
		}, {
			"voteReject",
			`energi.voteReject('0x00000000000000000000000000000000000', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd','password')`,
			`energi.voteReject('0x00000000000000000000000000000000000', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd','*****')`,
		}, {
			"withdrawFee",
			`energi.withdrawFee('0x00000000000000000000000000000000000', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd','password')`,
			`energi.withdrawFee('0x00000000000000000000000000000000000', '0x73E286b244c17F030F72e98C57FC83015a3C53Fd','*****')`,
		},
	}

	mockRPCresponse := map[string]string{
		"admin_nodeInfo": `{"enode": "enode://40865e11060ba79c3df7137e695ca452e4dbc5@127.0.0.1:49797"}`,
		"eth_accounts":   `["0x01d4A3FdEc411aBcfc1dece7013890C6ec2629f9"]`,
	}

	if len(td) < len(passwordIndexMap) {
		t.Fatalf("Some regex test case scenarios were skipped")
	}

	var mockSendFunc = func(call otto.FunctionCall) otto.Value {
		var req jsonrpcCall
		JSON, _ := call.Otto.Object("JSON")
		reqVal, _ := JSON.Call("stringify", call.Argument(0))
		json.NewDecoder(strings.NewReader(reqVal.String())).Decode(&req)

		resp, _ := call.Otto.Object(`({"jsonrpc":"2.0"})`)
		resp.Set("id", req.ID)
		resultVal, _ := JSON.Call("parse", mockRPCresponse[req.Method])
		resp.Set("result", resultVal)
		return resp.Value()
	}

	instance, err := newMaskingJSRE(mockSendFunc)
	if err != nil {
		t.Fatalf("expected no error but found: %v", err)
	}

	for _, data := range td {
		info, ok := passwordIndexMap[data.name]
		searchTerm := fmt.Sprintf("%s.%s", info.parent, data.name)
		if !ok {
			t.Fatalf("test for search term (%s) is not accounted for", searchTerm)
		}

		t.Run(searchTerm, func(t *testing.T) {
			if instance.IsPasswordMasked(data.original) {
				t.Fatalf("expected the command not to be password masked but it was")
			}

			expected, err := instance.MaskPassword(data.original)
			if err != nil {
				t.Fatalf("expected no error from MaskPassword() but found: %v", err)
			}

			if expected != data.masked {
				t.Fatalf("masking password command failed, required =%v found =%v", data.masked, expected)
			}

			if !instance.IsPasswordMasked(expected) {
				t.Fatalf("expected the command to be password masked but it wasn't")
			}

			expected, err = instance.UnMaskPassword(data.masked, passwordMask)
			if err != nil {
				t.Fatalf("expected no error from UnMaskPassword() but found: %v", err)
			}

			if expected != data.masked {
				t.Fatalf("unmasking password command failed, required =%v found =%v", data.masked, expected)
			}
		})
	}
}
