package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type CandidateListHandler struct {

}

func (c CandidateListHandler) Handler(ctx context.Context, req *irisProtoc.CandidateListRequest) (
	*irisProtoc.CandidateListResponse, error) {
	
	reqVO := c.BuildRequest(req)
	resVO, err := candidateService.List(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.BuildResponse(resVO), nil
}

func (c CandidateListHandler) BuildRequest(req *irisProtoc.CandidateListRequest) vo.CandidateListReqVO {
	reqVO := vo.CandidateListReqVO{
		Address: req.GetAddress(),
		Page: req.GetPage(),
		PerPage: req.GetPerPage(),
		Sort: req.GetSort(),
		Q: req.GetQ(),
	}
	
	return reqVO
}

func (c CandidateListHandler) BuildResponse(resVO vo.CandidateListResVO) *irisProtoc.CandidateListResponse  {
	var (
		response                irisProtoc.CandidateListResponse
		resCandidate            irisProtoc.Candidate
		resCandidates           []*irisProtoc.Candidate
		resCandidateDescription irisProtoc.Candidate_Description
		resCandidateDelegator   irisProtoc.Delegator
	)

	candidates := resVO.Candidates
	if len(candidates) > 0 {
		for _, v := range candidates {
			// description
			resCandidateDescription = irisProtoc.Candidate_Description{
				Details: v.Description.Details,
				Identity: v.Description.Identity,
				Moniker: v.Description.Moniker,
				Website: v.Description.Website,
			}
			
			// delegator
			var resCandidateDelegators []*irisProtoc.Delegator
			
			if len(v.Delegators) > 0 {
				delegator := v.Delegators[0]
				resCandidateDelegator = irisProtoc.Delegator{
					Address: delegator.Address,
					PubKey: delegator.PubKey,
					Shares: uint64(delegator.Shares),
				}
				resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
			}
			
			
			resCandidate = irisProtoc.Candidate{
				Address: v.Address,
				PubKey: v.PubKey,
				Shares: uint64(v.Shares),
				VotingPower: v.VotingPower,
				Description: &resCandidateDescription,
				Delegators: resCandidateDelegators,
			}
			
			resCandidates = append(resCandidates, &resCandidate)
		}
	}
	
	response = irisProtoc.CandidateListResponse{
		Candidates: resCandidates,
	}
	
	return &response
}
