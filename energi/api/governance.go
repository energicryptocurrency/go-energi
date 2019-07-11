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
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"energi.world/core/gen3/internal/ethapi"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	proposalCallGas uint64 = 500000
)

type GovernanceAPI struct {
	backend ethapi.Backend
}

func NewGovernanceAPI(b ethapi.Backend) *GovernanceAPI {
	return &GovernanceAPI{b}
}

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

func (g *GovernanceAPI) VoteAccept(
	proposal common.Address,
	mn_owner common.Address,
	password string,
) error {
	contract, err := g.proposal(password, mn_owner, proposal)
	if err != nil {
		log.Error("Failed", "err", err)
		return err
	}

	tx, err := contract.VoteAccept()

	log.Info("Note: please wait until proposal TX gets into a block!", "tx", tx.Hash())

	return err
}

func (g *GovernanceAPI) VoteReject(
	proposal common.Address,
	mn_owner common.Address,
	password string,
) error {
	contract, err := g.proposal(password, mn_owner, proposal)
	if err != nil {
		log.Error("Failed", "err", err)
		return err
	}

	tx, err := contract.VoteReject()

	log.Info("Note: please wait until proposal TX gets into a block!", "tx", tx.Hash())

	return err
}

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

func proposalInfo(backend ethapi.Backend, address common.Address) *ProposalInfo {
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

	curr_block := backend.CurrentBlock()
	if err != nil {
		log.Error("Failed at current block", "err", err)
		return nil
	}

	state, _, err := backend.StateAndHeaderByNumber(
		nil, rpc.BlockNumber(curr_block.Number().Int64()))
	if err != nil {
		log.Error("Failed at state", "err", err)
		return nil
	}

	balance := state.GetBalance(address)

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
		Balance:      (*hexutil.Big)(balance),
	}
}

type UpgradeProposalInfo struct {
	ProposalInfo
	Impl common.Address
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
	}
	return ret
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

func (g *GovernanceAPI) ListUpgrades() *UpgradeProposals {
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
