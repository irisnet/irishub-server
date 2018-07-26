package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
	"github.com/irisnet/irishub-server/utils/helper"
)

type DelegatorTotalSharesHandler struct {

}

func (c DelegatorTotalSharesHandler) Handler(ctx context.Context, req *irisProtoc.TotalShareRequest) (
	*irisProtoc.TotalShareResponse, error) {
	
	reqVO := c.BuildRequest(req)
	
	resVO, err := delegatorService.GetDelegatorTotalShare(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return c.BuildResponse(resVO), nil
}

func (c DelegatorTotalSharesHandler) BuildRequest(req *irisProtoc.TotalShareRequest) vo.TotalShareReqVO {
	
	reqVO := vo.TotalShareReqVO{
		Address: req.GetAddress(),
	}
	
	return reqVO
}

func (c DelegatorTotalSharesHandler) BuildResponse(resVO vo.TotalShareResVO) *irisProtoc.TotalShareResponse {
	
	response := irisProtoc.TotalShareResponse{
		TotalShares: helper.ConvertFloatToInt(resVO.TotalShare),
	}
	
	return &response
}