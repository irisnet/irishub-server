package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
	"golang.org/x/net/context"
)

type DelegatorCandidateListHandler struct {
}

func (h DelegatorCandidateListHandler) Handler(ctx context.Context, req *irisProtoc.DelegatorCandidateListRequest) (
	[]*irisProtoc.Candidate, error) {

	reqVO := h.buildRequest(req)

	resVO, err := delegatorService.DelegatorCandidateList(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}
	return h.buildResponse(resVO), nil
}

func (h DelegatorCandidateListHandler) buildRequest(req *irisProtoc.DelegatorCandidateListRequest) vo.DelegatorCandidateListReqVO {

	reqVO := vo.DelegatorCandidateListReqVO{
		Address: req.GetAddress(),
		Page:    req.GetPage(),
		PerPage: req.GetPerPage(),
		Q:       req.GetQ(),
	}

	return reqVO
}

func (h DelegatorCandidateListHandler) buildResponse(resVO vo.DelegatorCandidateListResVO) []*irisProtoc.Candidate {
	var (
		response []*irisProtoc.Candidate
	)

	candidates := resVO.Candidates
	if len(candidates) > 0 {
		for _, v := range candidates {
			var (
				resCandidate            irisProtoc.Candidate
				resCandidateDescription irisProtoc.CandidateDescription
				resCandidateDelegator   irisProtoc.Delegator
				resCandidateDelegators  []*irisProtoc.Delegator
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
				resCandidateDelegator = irisProtoc.Delegator{
					Address: delegator.Address,
					PubKey:  delegator.ValidatorAddr,
					Shares:  helper.ConvertFloatToInt(delegator.Shares),
				}
				resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
			}

			resCandidate = irisProtoc.Candidate{
				Address: v.Address,
				PubKey:  v.PubKey,
				//Shares: v.Shares,
				VotingPower: v.VotingPower,
				Description: &resCandidateDescription,
				Delegators:  resCandidateDelegators,
			}

			response = append(response, &resCandidate)
		}
	}

	return response
}
