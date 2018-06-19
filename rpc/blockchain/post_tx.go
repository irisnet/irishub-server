package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type PostTxHandler struct {
}

func (c PostTxHandler) Handler(ctx context.Context, req *commonProtoc.PostTxRequest) (
	*commonProtoc.PostTxResponse, error) {
	
	postTxVO := c.buildRequest(req)
	res, err := postTxService.PostTx(postTxVO)
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return c.buildResponse(res), nil
}


func (c PostTxHandler) buildRequest(request *commonProtoc.PostTxRequest) vo.PostTxReqVO {
	postTxVO := vo.PostTxReqVO{
		Tx: request.GetTx(),
	}
	
	return postTxVO
}

func (c PostTxHandler) buildResponse(res []byte) (*commonProtoc.PostTxResponse) {
	return &commonProtoc.PostTxResponse{
		TxHash: string(res),
	}
}
