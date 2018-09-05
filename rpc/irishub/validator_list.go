package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type ValidatorListHandler struct {
}

func (c ValidatorListHandler) Handler(ctx context.Context, req *irisProtoc.CandidateListRequest) (
	[]*irisProtoc.Candidate, error) {

	reqVO := c.buildRequest(req)
	resVO, err := validatorService.List(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return c.buildResponse(resVO), nil
}

func (c ValidatorListHandler) buildRequest(req *irisProtoc.CandidateListRequest) vo.ValidatorListReqVO {
	reqVO := vo.ValidatorListReqVO{
		Address: req.GetAddress(),
		Page:    req.GetPage(),
		PerPage: req.GetPerPage(),
		Sort:    req.GetSort(),
		Q:       req.GetQ(),
	}

	return reqVO
}

func (c ValidatorListHandler) buildResponse(resVO vo.ValidatorListResVO) []*irisProtoc.Candidate {
	var (
		response []*irisProtoc.Candidate
	)

	candidates := resVO.Candidates
	if len(candidates) > 0 {
		for _, v := range candidates {
			resCandidate := BuildCandidateResponse(v)
			response = append(response, &resCandidate)
		}
	}
	return response
}
