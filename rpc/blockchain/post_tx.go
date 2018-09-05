package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type PostTxHandler struct {
}

func (c PostTxHandler) Handler(ctx context.Context, req *commonProtoc.PostTxRequest) (
	*commonProtoc.PostTxResponse, error) {

	reqVO := c.buildRequest(req)
	res, err := postTxService.PostTx(reqVO)
	if err.IsNotNull() {
		return nil, BuildException(err)
	}
	return c.buildResponse(res), nil
}

func (c PostTxHandler) buildRequest(request *commonProtoc.PostTxRequest) vo.PostTxReqVO {
	reqVO := vo.PostTxReqVO{
		Tx: request.GetTx(),
	}

	return reqVO
}

func (c PostTxHandler) buildResponse(resVO vo.PostTxResVO) *commonProtoc.PostTxResponse {
	return &commonProtoc.PostTxResponse{
		TxHash: resVO.TxHash,
	}
}
