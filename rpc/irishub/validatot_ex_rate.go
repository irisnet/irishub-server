package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type ValidatorExRateHandler struct {
}

func (h ValidatorExRateHandler) Handle(ctx context.Context, req *irisProtoc.ValidatorExRateRequest) (
	*irisProtoc.ValidatorExRateResponse, error) {

	reqVO := h.buildRequest(req)

	resVO, err := validatorService.GetValidatorExRate(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return h.buildResponse(resVO), nil
}

func (h ValidatorExRateHandler) buildRequest(req *irisProtoc.ValidatorExRateRequest) vo.ValidatorExRateReqVO {

	reqVO := vo.ValidatorExRateReqVO{
		ValidatorAddress: req.GetValidatorAddress(),
	}

	return reqVO
}

func (h ValidatorExRateHandler) buildResponse(resVO vo.ValidatorExRateResVO) *irisProtoc.ValidatorExRateResponse {
	var (
		res irisProtoc.ValidatorExRateResponse
	)

	res = irisProtoc.ValidatorExRateResponse{
		TokenSharesRate: resVO.ExRate,
	}

	return &res
}
