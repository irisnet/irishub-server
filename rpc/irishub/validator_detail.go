package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
	"github.com/irisnet/irishub-server/utils/helper"
)

type ValidatorDetailHandler struct {

}

func (h ValidatorDetailHandler) Handler(ctx context.Context, req *irisProtoc.CandidateDetailRequest) (
	*irisProtoc.Candidate, error) {
	
	reqVO := h.BuildRequest(req)
	
	resVO, err := validatorService.Detail(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return h.BuildResponse(resVO), nil
}

func (h ValidatorDetailHandler) BuildRequest(req *irisProtoc.CandidateDetailRequest) vo.ValidatorDetailReqVO {
	
	reqVO := vo.ValidatorDetailReqVO{
		Address: req.GetAddress(),
		ValAddr: req.GetPubKey(),
	}
	
	return reqVO
}

func (h ValidatorDetailHandler) BuildResponse(resVO vo.ValidatorDetailResVO) *irisProtoc.Candidate {
	var (
		response irisProtoc.Candidate
		resCandidateDescription irisProtoc.CandidateDescription
		resCandidateDelegator irisProtoc.Delegator
		
		resCandidateDelegators []*irisProtoc.Delegator
	)
	
	candidate := resVO.Candidate
	
	// description
	resCandidateDescription = irisProtoc.CandidateDescription{
		Details: candidate.Description.Details,
		Identity: candidate.Description.Identity,
		Moniker: candidate.Description.Moniker,
		Website: candidate.Description.Website,
	}
	
	// delegators
	if len(candidate.Delegators) > 0 {
		delegator := candidate.Delegators[0]
		resCandidateDelegator = irisProtoc.Delegator{
			Address: delegator.Address,
			PubKey: delegator.ValidatorAddr,
			Shares: helper.ConvertFloatToInt(delegator.Shares),
		}
		resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
	}
	
	
	response = irisProtoc.Candidate{
		Address: candidate.Address,
		PubKey: candidate.PubKey,
		//Shares: candidate.Shares,
		VotingPower: candidate.VotingPower,
		Description: &resCandidateDescription,
		Delegators: resCandidateDelegators,
		
	}
	
	return &response
}