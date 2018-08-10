package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
	"golang.org/x/net/context"
)

type ValidatorDetailHandler struct {
}

func (h ValidatorDetailHandler) Handler(ctx context.Context, req *irisProtoc.CandidateDetailRequest) (
	*irisProtoc.Candidate, error) {

	reqVO := h.buildRequest(req)

	resVO, err := validatorService.Detail(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return h.buildResponse(resVO), nil
}

func (h ValidatorDetailHandler) buildRequest(req *irisProtoc.CandidateDetailRequest) vo.ValidatorDetailReqVO {

	reqVO := vo.ValidatorDetailReqVO{
		Address: req.GetAddress(),
		ValAddr: req.GetPubKey(),
	}

	return reqVO
}

func (h ValidatorDetailHandler) buildResponse(resVO vo.ValidatorDetailResVO) *irisProtoc.Candidate {
	var (
		response                irisProtoc.Candidate
		resCandidateDescription irisProtoc.CandidateDescription
		resCandidateDelegator   irisProtoc.Delegator

		resCandidateDelegators []*irisProtoc.Delegator
	)

	candidate := resVO.Candidate

	// description
	resCandidateDescription = irisProtoc.CandidateDescription{
		Details:  candidate.Description.Details,
		Identity: candidate.Description.Identity,
		Moniker:  candidate.Description.Moniker,
		Website:  candidate.Description.Website,
	}

	// delegators
	if len(candidate.Delegators) > 0 {
		delegator := candidate.Delegators[0]
		resCandidateDelegator = irisProtoc.Delegator{
			Address: delegator.Address,
			PubKey:  delegator.ValidatorAddr,
			Shares:  helper.ConvertFloatToInt(delegator.Shares),
		}
		resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
	}

	response = irisProtoc.Candidate{
		Address: candidate.Address,
		PubKey:  candidate.PubKey,
		//Shares: candidate.Shares,
		VotingPower: candidate.VotingPower,
		Description: &resCandidateDescription,
		Delegators:  resCandidateDelegators,
	}

	return &response
}
