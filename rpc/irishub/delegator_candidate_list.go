package irishub

import (
	"fmt"
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
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

			resCandidate = irisProtoc.Candidate{
				Address:     v.Address,
				PubKey:      v.PubKey,
				VotingPower: v.VotingPower,
				Description: &resCandidateDescription,
				Delegators:  resCandidateDelegators,
			}

			response = append(response, &resCandidate)
		}
	}

	return response
}
