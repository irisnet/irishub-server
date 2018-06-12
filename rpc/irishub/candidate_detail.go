package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type CandidateDetailHandler struct {

}

func (h CandidateDetailHandler) Handler(ctx context.Context, req *irisProtoc.CandidateDetailRequest) (
	*irisProtoc.CandidateDetailResponse, error) {
	
	reqVO := h.BuildRequest(req)
	
	resVO, err := candidateService.Detail(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return h.BuildResponse(resVO), nil
}

func (h CandidateDetailHandler) BuildRequest(req *irisProtoc.CandidateDetailRequest) vo.CandidateDetailReqVO {
	
	reqVO := vo.CandidateDetailReqVO{
		Address: req.GetAddress(),
		PubKey: req.GetPubKey(),
	}
	
	return reqVO
}

func (h CandidateDetailHandler) BuildResponse(resVO vo.CandidateDetailResVO) *irisProtoc.CandidateDetailResponse {
	var (
		response irisProtoc.CandidateDetailResponse
		resCandidate irisProtoc.Candidate
		resCandidateDescription irisProtoc.Candidate_Description
		resCandidateDelegator irisProtoc.Delegator
		
		resCandidateDelegators []*irisProtoc.Delegator
	)
	
	candidate := resVO.Candidate
	
	// description
	resCandidateDescription = irisProtoc.Candidate_Description{
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
			PubKey: delegator.PubKey,
			Shares: uint64(delegator.Shares),
		}
		resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
	}
	
	
	resCandidate = irisProtoc.Candidate{
		Address: candidate.Address,
		PubKey: candidate.PubKey,
		Shares: uint64(candidate.Shares),
		VotingPower: candidate.VotingPower,
		Description: &resCandidateDescription,
		Delegators: resCandidateDelegators,
		
	}
	
	response = irisProtoc.CandidateDetailResponse{
		Candidate: &resCandidate,
	}
	
	return &response
}