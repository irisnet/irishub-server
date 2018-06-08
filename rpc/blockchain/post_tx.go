package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type PostTxController struct {
}

func (c PostTxController) Handler(ctx context.Context, req *chainModel.PostTxRequest) (
	*chainModel.PostTxResponse, error) {
	
	postTxVO := c.buildRequest(req)
	res, err := postTxService.PostTx(postTxVO)
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return c.buildResponse(res), nil
}


func (c PostTxController) buildRequest(request *chainModel.PostTxRequest) vo.PostTxReqVO {
	postTxVO := vo.PostTxReqVO{
		Tx: request.GetTx(),
	}
	
	return postTxVO
}

func (c PostTxController) buildResponse(res []byte) (*chainModel.PostTxResponse) {
	return &chainModel.PostTxResponse{
		TxHash: string(res),
	}
}
