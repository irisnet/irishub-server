package irishub

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	irisProtoc "github.com/irisnet/irisnet-rpc/irishub/codegen/server/model"
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
		DelAddr: req.DelAddress,
		ValAddr: req.ValAddress,
	}

	return reqVO
}

func (h ValidatorDetailHandler) buildResponse(resVO vo.ValidatorDetailResVO) *irisProtoc.Candidate {
	var (
		response irisProtoc.Candidate
	)

	response = BuildCandidateResponse(resVO.Candidate)

	return &response
}
