package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"golang.org/x/net/context"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/rpc"
)

type ExRateHandler struct {
	
}

func (h ExRateHandler) Handle(ctx context.Context, req *irisProtoc.ExRateRequest) (
	*irisProtoc.ExRateResponse, error) {

	reqVO := h.BuildRequest(req)

	resVO, err := shareService.GetExRate(reqVO)

	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}

	return h.BuildResponse(resVO), nil
}

func (h ExRateHandler) BuildRequest(req *irisProtoc.ExRateRequest) vo.ExRateReqVO {
	
	reqVO := vo.ExRateReqVO{
		ValidatorAddress: req.GetValidatorAddress(),
	}
	
	return reqVO
}

func (h ExRateHandler) BuildResponse(resVO vo.ExRateResVO) *irisProtoc.ExRateResponse {
	var (
		res irisProtoc.ExRateResponse
	)

	res = irisProtoc.ExRateResponse{
		TokenSharesRate: resVO.ExRate,
	}

	return &res
}
