package irishub

import (
	"fmt"
	"github.com/irisnet/irishub-server/models/document"
	irisProtoc "github.com/irisnet/irisnet-rpc/irishub/codegen/server/model"
)

const (
	CandidateTypeConsensus = "consensus"
	CandidateTypeCandidate = "candidate"
	CandidateTypeJailed    = "jailed"
)

func BuildCandidateResponse(v document.Candidate) irisProtoc.Candidate {
	var (
		resCandidate            irisProtoc.Candidate
		resCandidateDescription irisProtoc.CandidateDescription
		resCandidateDelegator   irisProtoc.Delegator
		resCandidateDelegators  []*irisProtoc.Delegator
		resDelegatorUbd         irisProtoc.DelegatorUnbondingDelegation
	)

	// description
	resCandidateDescription = irisProtoc.CandidateDescription{
		Details:  v.Description.Details,
		Identity: v.Description.Identity,
		Moniker:  v.Description.Moniker,
		Website:  v.Description.Website,
	}

	// delegator
	if len(v.Delegators) > 0 {
		delegator := v.Delegators[0]
		if balance := delegator.UnbondingDelegation.Balance; len(balance) > 0 {
			resDelegatorUbd = irisProtoc.DelegatorUnbondingDelegation{
				Tokens:  balance[0].Amount,
				MinTime: fmt.Sprintf("%v", delegator.UnbondingDelegation.MinTime),
			}
		}

		resCandidateDelegator = irisProtoc.Delegator{
			Address:             delegator.Address,
			ValAddress:          delegator.ValidatorAddr,
			Shares:              delegator.Shares,
			BondedTokens:        delegator.BondedTokens,
			UnbondingDelegation: &resDelegatorUbd,
		}
		resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
	}

	var cType string
	if v.Jailed {
		cType = CandidateTypeJailed
	} else if v.Status == "Bonded" {
		cType = CandidateTypeConsensus
	} else if v.Status == "Unbonded" || v.Status == "Unbonding" {
		cType = CandidateTypeCandidate
	}

	resCandidate = irisProtoc.Candidate{
		Address:     v.Address,
		PubKey:      v.PubKey,
		UpTime:      v.UpTime,
		VotingPower: v.VotingPower,
		Description: &resCandidateDescription,
		Delegators:  resCandidateDelegators,
		Type:        cType,
		Number:      int8(v.Rank.Number),
		Lift:        int8(v.Rank.Lift),
	}

	return resCandidate
}
