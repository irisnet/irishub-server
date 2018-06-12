package irishub

import (
	
	irisModel "github.com/irisnet/irishub-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type ShareController struct {

}

func (c ShareController) Handler(ctx context.Context, req *irisModel.TotalShareRequest) (
	*irisModel.TotalShareResponse, error) {
	
	reqVO := c.BuildRequest(req)
	
	resVO, err := shareService.GetDelegatorTotalShare(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return c.BuildResponse(resVO), nil
}

func (c ShareController) BuildRequest(req *irisModel.TotalShareRequest) vo.ShareReqVO  {
	
	reqVO := vo.ShareReqVO{
		Address: req.GetAddress(),
	}
	
	return reqVO
}

func (c ShareController) BuildResponse(resVO vo.ShareResVO) *irisModel.TotalShareResponse {
	
	response := irisModel.TotalShareResponse{
		TotalShares: resVO.TotalShare,
	}
	
	return &response
}