package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"golang.org/x/net/context"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
)

type AccountNumberHandler struct {

}

func (h AccountNumberHandler) Handler(ctx context.Context, req *irisProtoc.AccountNumRequest) (
	*irisProtoc.AccountNumResponse, error) {

	reqVO := h.BuildRequest(req)

	resVO, err := accountService.GetAccountNum(reqVO)

	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}

	return h.BuildResponse(resVO), nil
}

func (h AccountNumberHandler) BuildRequest(req *irisProtoc.AccountNumRequest) vo.AccountNumReqVO  {
	return vo.AccountNumReqVO{
		Address: req.GetAddress(),
	}
}

func (h AccountNumberHandler) BuildResponse(resVO vo.AccountNumResVO) *irisProtoc.AccountNumResponse {
	response := irisProtoc.AccountNumResponse{
		AccountNum: resVO.AccountNum,
	}

	return &response
}
