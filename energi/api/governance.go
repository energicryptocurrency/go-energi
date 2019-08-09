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

package api

import (
	"errors"
	"math/big"

	"github.com/pborman/uuid"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	proposalCallGas uint64 = 500000
	upgradeCallGas  uint64 = 5000000
)

type GovernanceAPI struct {
	backend Backend
}

func NewGovernanceAPI(b Backend) *GovernanceAPI {
	return &GovernanceAPI{b}
}

//=============================================================================

func (g *GovernanceAPI) proposal(
	password string,
	owner common.Address,
	proposal common.Address,
) (session *energi_abi.IProposalSession, err error) {
	account := accounts.Account{Address: owner}
	wallet, err := g.backend.AccountManager().Find(account)
	if err != nil {
		return nil, err
	}

	contract, err := energi_abi.NewIProposal(proposal, g.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IProposalSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From: owner,
		},
		TransactOpts: bind.TransactOpts{
			From: owner,
			Signer: func(
				signer types.Signer,
				addr common.Address,
				tx *types.Transaction,
			) (*types.Transaction, error) {
				return wallet.SignTxWithPassphrase(
					account, password, tx, g.backend.ChainConfig().ChainID)
			},
			Value:    common.Big0,
			GasLimit: proposalCallGas,
		},
	}
	return
}

//=============================================================================
// VT-5 Voting API
//=============================================================================

func (g *GovernanceAPI) VoteAccept(
	proposal common.Address,
	mn_owner common.Address,
	password string,
) (txhash common.Hash, err error) {
	contract, err := g.proposal(password, mn_owner, proposal)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	tx, err := contract.VoteAccept()

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (g *GovernanceAPI) VoteReject(
	proposal common.Address,
	mn_owner common.Address,
	password string,
) (txhash common.Hash, err error) {
	contract, err := g.proposal(password, mn_owner, proposal)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	tx, err := contract.VoteReject()

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (g *GovernanceAPI) WithdrawFee(
	proposal common.Address,
	payer common.Address,
	password string,
) (txhash common.Hash, err error) {
	contract, err := g.proposal(password, payer, proposal)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	if res, _ := contract.IsAccepted(); !res {
		err = errors.New("The proposal is not accepted!")
		log.Error("Failed", "err", err)
		return
	}

	tx, err := contract.Withdraw()

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

//=============================================================================
// Generic proposal info
//=============================================================================

type ProposalInfo struct {
	Proposal     common.Address
	Proposer     common.Address
	CreatedBlock uint64
	Deadline     uint64
	QuorumWeight *hexutil.Big
	TotalWeight  *hexutil.Big
	RejectWeight *hexutil.Big
	AcceptWeight *hexutil.Big
	Finished     bool
	Accepted     bool
	Balance      *hexutil.Big
}

func getBalance(backend Backend, address common.Address) *hexutil.Big {
	curr_block := backend.CurrentBlock()

	state, _, err := backend.StateAndHeaderByNumber(
		nil, rpc.BlockNumber(curr_block.Number().Int64()))
	if err != nil {
		log.Error("Failed at state", "err", err)
		return nil
	}

	return (*hexutil.Big)(state.GetBalance(address))
}

func proposalInfo(backend Backend, address common.Address) *ProposalInfo {
	if (address == common.Address{}) {
		return nil
	}

	proposal, err := energi_abi.NewIProposalCaller(
		address, backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed at NewIProposalCaller", "err", err)
		return nil
	}

	call_opts := &bind.CallOpts{}
	if err != nil {
		log.Error("Failed at CallOpts", "err", err)
		return nil
	}

	proposer, err := proposal.FeePayer(call_opts)
	if err != nil {
		log.Error("Failed at FeePayer", "err", err)
		return nil
	}

	block, err := proposal.CreatedBlock(call_opts)
	if err != nil {
		log.Error("Failed at CreatedBlock", "err", err)
		return nil
	}

	deadline, err := proposal.Deadline(call_opts)
	if err != nil {
		log.Error("Failed at Deadline", "err", err)
		return nil
	}

	quorum_w, err := proposal.QuorumWeight(call_opts)
	if err != nil {
		log.Error("Failed at QuorumWeight", "err", err)
		return nil
	}

	total_w, err := proposal.TotalWeight(call_opts)
	if err != nil {
		log.Error("Failed at TotalWeight", "err", err)
		return nil
	}

	rejected_w, err := proposal.RejectedWeight(call_opts)
	if err != nil {
		log.Error("Failed at RejectedWeight", "err", err)
		return nil
	}

	accepted_w, err := proposal.AcceptedWeight(call_opts)
	if err != nil {
		log.Error("Failed at AcceptedWeight", "err", err)
		return nil
	}

	finished, err := proposal.IsFinished(call_opts)
	if err != nil {
		log.Error("Failed at IsFinished", "err", err)
		return nil
	}

	accepted, err := proposal.IsAccepted(call_opts)
	if err != nil {
		log.Error("Failed at IsAccepted", "err", err)
		return nil
	}

	return &ProposalInfo{
		Proposal:     address,
		Proposer:     proposer,
		CreatedBlock: block.Uint64(),
		Deadline:     deadline.Uint64(),
		QuorumWeight: (*hexutil.Big)(quorum_w),
		TotalWeight:  (*hexutil.Big)(total_w),
		RejectWeight: (*hexutil.Big)(rejected_w),
		AcceptWeight: (*hexutil.Big)(accepted_w),
		Finished:     finished,
		Accepted:     accepted,
		Balance:      getBalance(backend, address),
	}
}

//=============================================================================
// SC-15: Upgrade API
//=============================================================================

type UpgradeProposalInfo struct {
	ProposalInfo
	Impl  common.Address
	Proxy common.Address
}

func (g *GovernanceAPI) upgradeProposalInfo(proxy common.Address) []UpgradeProposalInfo {
	proxy_obj, err := energi_abi.NewIGovernedProxyCaller(
		proxy, g.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed NewIGovernedProxyCaller", "err", err)
		return nil
	}

	call_opts := &bind.CallOpts{}
	proposals, err := proxy_obj.ListUpgradeProposals(call_opts)
	if err != nil {
		log.Error("Failed ListUpgradeProposals", "err", err)
		return nil
	}

	ret := make([]UpgradeProposalInfo, len(proposals))
	for i, p := range proposals {
		ret[i].ProposalInfo = *proposalInfo(g.backend, p)
		impl, err := proxy_obj.UpgradeProposalImpl(call_opts, p)
		if err != nil {
			log.Error("Failed UpgradeProposalImpl", "err", err)
			continue
		}
		ret[i].Impl = impl
		ret[i].Proxy = proxy
	}
	return ret
}

func (g *GovernanceAPI) governedProxy(
	password string,
	owner common.Address,
	proxy common.Address,
) (session *energi_abi.IGovernedProxySession, err error) {
	account := accounts.Account{Address: owner}
	wallet, err := g.backend.AccountManager().Find(account)
	if err != nil {
		return nil, err
	}

	contract, err := energi_abi.NewIGovernedProxy(
		proxy, g.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.IGovernedProxySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From: owner,
		},
		TransactOpts: bind.TransactOpts{
			From: owner,
			Signer: func(
				signer types.Signer,
				addr common.Address,
				tx *types.Transaction,
			) (*types.Transaction, error) {
				return wallet.SignTxWithPassphrase(
					account, password, tx, g.backend.ChainConfig().ChainID)
			},
			Value:    common.Big0,
			GasLimit: upgradeCallGas,
		},
	}
	return
}

type UpgradeProposals struct {
	Treasury           []UpgradeProposalInfo
	MasternodeRegistry []UpgradeProposalInfo
	StakerReward       []UpgradeProposalInfo
	BackboneReward     []UpgradeProposalInfo
	SporkRegistry      []UpgradeProposalInfo
	CheckpointRegistry []UpgradeProposalInfo
	BlacklistRegistry  []UpgradeProposalInfo
	MasternodeToken    []UpgradeProposalInfo
}

func (g *GovernanceAPI) UpgradeInfo() *UpgradeProposals {
	ret := new(UpgradeProposals)
	ret.Treasury = g.upgradeProposalInfo(energi_params.Energi_Treasury)
	ret.MasternodeRegistry = g.upgradeProposalInfo(energi_params.Energi_MasternodeRegistry)
	ret.StakerReward = g.upgradeProposalInfo(energi_params.Energi_StakerReward)
	ret.BackboneReward = g.upgradeProposalInfo(energi_params.Energi_BackboneReward)
	ret.SporkRegistry = g.upgradeProposalInfo(energi_params.Energi_SporkRegistry)
	ret.CheckpointRegistry = g.upgradeProposalInfo(energi_params.Energi_CheckpointRegistry)
	ret.BlacklistRegistry = g.upgradeProposalInfo(energi_params.Energi_BlacklistRegistry)
	ret.MasternodeToken = g.upgradeProposalInfo(energi_params.Energi_MasternodeToken)
	return ret
}

func (g *GovernanceAPI) UpgradePropose(
	proxy common.Address,
	new_impl common.Address,
	period uint64,
	fee *hexutil.Big,
	payer common.Address,
	password string,
) (txhash common.Hash, err error) {
	session, err := g.governedProxy(password, payer, proxy)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	session.TransactOpts.Value = fee.ToInt()
	tx, err := session.ProposeUpgrade(new_impl, new(big.Int).SetUint64(period))

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (g *GovernanceAPI) UpgradePerform(
	proxy common.Address,
	proposal common.Address,
	payer common.Address,
	password string,
) (txhash common.Hash, err error) {
	session, err := g.governedProxy(password, payer, proxy)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	tx, err := session.Upgrade(proposal)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the upgrade TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

func (g *GovernanceAPI) UpgradeCollect(
	proxy common.Address,
	proposal common.Address,
	payer common.Address,
	password string,
) (txhash common.Hash, err error) {
	session, err := g.governedProxy(password, payer, proxy)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	tx, err := session.CollectUpgradeProposal(proposal)

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}

//=============================================================================
// GOV-9: Treasury API
//=============================================================================

func (g *GovernanceAPI) treasury(
	password string,
	payer common.Address,
) (session *energi_abi.ITreasurySession, err error) {
	account := accounts.Account{Address: payer}
	wallet, err := g.backend.AccountManager().Find(account)
	if err != nil {
		return nil, err
	}

	contract, err := energi_abi.NewITreasury(
		energi_params.Energi_Treasury, g.backend.(bind.ContractBackend))
	if err != nil {
		return nil, err
	}

	session = &energi_abi.ITreasurySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			From: payer,
		},
		TransactOpts: bind.TransactOpts{
			From: payer,
			Signer: func(
				signer types.Signer,
				addr common.Address,
				tx *types.Transaction,
			) (*types.Transaction, error) {
				return wallet.SignTxWithPassphrase(
					account, password, tx, g.backend.ChainConfig().ChainID)
			},
			Value:    common.Big0,
			GasLimit: proposalCallGas,
		},
	}
	return
}

type BudgetProposalInfo struct {
	ProposalInfo
	ProposedAmount *hexutil.Big
	PaidAmount     *hexutil.Big
	RefUUID        string
}

type BudgetInfo struct {
	Balance   *hexutil.Big
	Proposals []BudgetProposalInfo
}

func (g *GovernanceAPI) BudgetInfo() *BudgetInfo {
	treasury, err := energi_abi.NewITreasuryCaller(
		energi_params.Energi_Treasury, g.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed NewITreasuryCaller", "err", err)
		return nil
	}

	proxy, err := energi_abi.NewIGovernedProxyCaller(
		energi_params.Energi_Treasury, g.backend.(bind.ContractCaller))
	if err != nil {
		log.Error("Failed NewITreasuryCaller", "err", err)
		return nil
	}

	call_opts := &bind.CallOpts{}

	proposals, err := treasury.ListProposals(call_opts)
	if err != nil {
		log.Error("Failed ListProposals", "err", err)
		return nil
	}

	impl, err := proxy.Impl(call_opts)
	if err != nil {
		log.Error("Failed Impl", "err", err)
		return nil
	}

	ret := make([]BudgetProposalInfo, len(proposals))
	for i, p := range proposals {
		ret[i].ProposalInfo = *proposalInfo(g.backend, p)

		budger_proposal, err := energi_abi.NewIBudgetProposalCaller(
			p, g.backend.(bind.ContractCaller))
		if err != nil {
			log.Error("Failed at NewIBudgetProposalCaller", "err", err)
			return nil
		}

		proposed_amount, err := budger_proposal.ProposedAmount(call_opts)
		if err != nil {
			log.Error("Failed ProposedAmount", "err", err)
			continue
		}
		paid_amount, err := budger_proposal.PaidAmount(call_opts)
		if err != nil {
			log.Error("Failed ProposedAmount", "err", err)
			continue
		}
		ref_uuid, err := budger_proposal.RefUuid(call_opts)
		if err != nil {
			log.Error("Failed ProposedAmount", "err", err)
			continue
		}
		ret[i].ProposedAmount = (*hexutil.Big)(proposed_amount)
		ret[i].PaidAmount = (*hexutil.Big)(paid_amount)
		ret[i].RefUUID = uuid.UUID(common.LeftPadBytes(ref_uuid.Bytes(), 16)).String()
	}

	return &BudgetInfo{
		Balance:   getBalance(g.backend, impl),
		Proposals: ret,
	}
}

func (g *GovernanceAPI) BudgetPropose(
	amount *hexutil.Big,
	ref_uuid string,
	period uint64,
	fee *hexutil.Big,
	payer common.Address,
	password string,
) (txhash common.Hash, err error) {
	session, err := g.treasury(password, payer)
	if err != nil {
		log.Error("Failed", "err", err)
		return
	}

	ref_uuid_b := uuid.Parse(ref_uuid)
	if ref_uuid_b == nil {
		err = errors.New("Failed to parse UUID")
		log.Error("Failed", "err", err)
		return
	}

	session.TransactOpts.Value = fee.ToInt()
	tx, err := session.Propose(
		(*big.Int)(amount),
		new(big.Int).SetBytes(ref_uuid_b),
		new(big.Int).SetUint64(period))

	if tx != nil {
		txhash = tx.Hash()
		log.Info("Note: please wait until the proposal TX gets into a block!", "tx", txhash.Hex())
	}

	return
}
