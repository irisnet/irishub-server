package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"golang.org/x/net/context"
)

type BlockChainRPCServices struct {
}

// get sequence
func (s BlockChainRPCServices) GetSequence(ctx context.Context, req *chainModel.SequenceRequest) (
	*chainModel.SequenceResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*chainModel.SequenceResponse), err
}

// build tx
func (s BlockChainRPCServices) BuildTx(ctx context.Context, req *chainModel.BuildTxRequest) (
	*chainModel.BuildTxResponse, error) {
	res, err := Handler(ctx, req)
	return res.(*chainModel.BuildTxResponse), err
}

// post tx
func (s BlockChainRPCServices) PostTx(ctx context.Context, req *chainModel.PostTxRequest) (
	*chainModel.PostTxResponse, error) {
	res, err := Handler(ctx, req)
	return res.(*chainModel.PostTxResponse), err
}




