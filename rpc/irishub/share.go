package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type ShareHandler struct {

}

func (c ShareHandler) Handler(ctx context.Context, req *irisProtoc.TotalShareRequest) (
	*irisProtoc.TotalShareResponse, error) {
	
	reqVO := c.BuildRequest(req)
	
	resVO, err := shareService.GetDelegatorTotalShare(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return c.BuildResponse(resVO), nil
}

func (c ShareHandler) BuildRequest(req *irisProtoc.TotalShareRequest) vo.ShareReqVO  {
	
	reqVO := vo.ShareReqVO{
		Address: req.GetAddress(),
	}
	
	return reqVO
}

func (c ShareHandler) BuildResponse(resVO vo.ShareResVO) *irisProtoc.TotalShareResponse {
	
	response := irisProtoc.TotalShareResponse{
		TotalShares: int64(resVO.TotalShare),
	}
	
	return &response
}