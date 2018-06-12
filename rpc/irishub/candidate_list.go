package irishub

import (
	irisModel "github.com/irisnet/irishub-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type CandidateListController struct {

}

func (c CandidateListController) Handler(ctx context.Context, req *irisModel.CandidateListRequest) (
	*irisModel.CandidateListResponse, error) {
	
	reqVO := c.BuildRequest(req)
	resVO, err := candidateService.List(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.BuildResponse(resVO), nil
}

func (c CandidateListController) BuildRequest(req *irisModel.CandidateListRequest) vo.CandidateListReqVO {
	reqVO := vo.CandidateListReqVO{
		Address: req.GetAddress(),
		Page: req.GetPage(),
		PerPage: req.GetPerPage(),
		Sort: req.GetSort(),
		Q: req.GetQ(),
	}
	
	return reqVO
}

func (c CandidateListController) BuildResponse(resVO vo.CandidateListResVO) *irisModel.CandidateListResponse  {
	var (
		response irisModel.CandidateListResponse
		resCandidate irisModel.Candidate
		resCandidates []*irisModel.Candidate
		resCandidateDescription irisModel.Candidate_Description
		resCandidateDelegator irisModel.Delegator
	)

	candidates := resVO.Candidates
	if len(candidates) > 0 {
		for _, v := range candidates {
			resCandidateDescription = irisModel.Candidate_Description{
				Details: v.Description.Details,
				Identity: v.Description.Identity,
				Moniker: v.Description.Moniker,
				Website: v.Description.Website,
			}
			
			var resCandidateDelegators []*irisModel.Delegator
			
			if len(v.Delegators) > 0 {
				delegator := v.Delegators[0]
				resCandidateDelegator = irisModel.Delegator{
					Address: delegator.Address,
					PubKey: delegator.PubKey,
					Shares: uint64(delegator.Shares),
				}
				resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
			}
			
			
			resCandidate = irisModel.Candidate{
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
	
	response = irisModel.CandidateListResponse{
		Candidates: resCandidates,
	}
	
	return &response
}
