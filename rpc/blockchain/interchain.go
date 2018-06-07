package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"golang.org/x/net/context"
)

type BlockChainRPC struct {
}

func (s BlockChainRPC) BuildTx(ctx context.Context, request *chainModel.BuildTxRequest) (
	*chainModel.BuildTxResponse, error) {
	res, err := Handler(ctx, request)
	return res.(*chainModel.BuildTxResponse), err
}

func (s BlockChainRPC) PostTx(ctx context.Context, request *chainModel.PostTxRequest) (
	*chainModel.PostTxResponse, error) {
	res, err := Handler(ctx, request)
	return res.(*chainModel.PostTxResponse), err
}

