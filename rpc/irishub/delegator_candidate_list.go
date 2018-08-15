package irishub

import (
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
			resCandidate := BuildCandidateResponse(v)
			response = append(response, &resCandidate)
		}
	}

	return response
}
