package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/iris-api-server/rpc"
	"golang.org/x/net/context"
)

type BlockChainRPC struct {
}

func (s BlockChainRPC) BuildTx(context context.Context, request *chainModel.BuildTxRequest) (
	*chainModel.BuildTxResponse, error) {
	buildTxVO := TransformBuildTxRequest(request)
	res, err := buildTxService.BuildTx(buildTxVO)
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	response := TransformBuildTxResponse(res)
	return response, nil
}

