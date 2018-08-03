package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"golang.org/x/net/context"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/rpc"
)

type ValidatorExRateHandler struct {
	
}

func (h ValidatorExRateHandler) Handle(ctx context.Context, req *irisProtoc.ValidatorExRateRequest) (
	*irisProtoc.ValidatorExRateResponse, error) {

	reqVO := h.BuildRequest(req)

	resVO, err := validatorService.GetValidatorExRate(reqVO)

	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}

	return h.BuildResponse(resVO), nil
}

func (h ValidatorExRateHandler) BuildRequest(req *irisProtoc.ValidatorExRateRequest) vo.ValidatorExRateReqVO {
	
	reqVO := vo.ValidatorExRateReqVO{
		ValidatorAddress: req.GetValidatorAddress(),
	}
	
	return reqVO
}

func (h ValidatorExRateHandler) BuildResponse(resVO vo.ValidatorExRateResVO) *irisProtoc.ValidatorExRateResponse {
	var (
		res irisProtoc.ValidatorExRateResponse
	)

	res = irisProtoc.ValidatorExRateResponse{
		TokenSharesRate: resVO.ExRate,
	}

	return &res
}