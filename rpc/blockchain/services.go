package blockchain

import (
	
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"golang.org/x/net/context"
)

type BlockChainRPCServices struct {
}

// get sequence
func (s BlockChainRPCServices) GetSequence(ctx context.Context, req *commonProtoc.SequenceRequest) (
	*commonProtoc.SequenceResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*commonProtoc.SequenceResponse), err
}

// build tx
func (s BlockChainRPCServices) BuildTx(ctx context.Context, req *commonProtoc.BuildTxRequest) (
	*commonProtoc.BuildTxResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*commonProtoc.BuildTxResponse), err
}

// post tx
func (s BlockChainRPCServices) PostTx(ctx context.Context, req *commonProtoc.PostTxRequest) (
	*commonProtoc.PostTxResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*commonProtoc.PostTxResponse), err
}

// get balance
func (s BlockChainRPCServices) GetBalance(ctx context.Context, req *commonProtoc.BalanceRequest) (
	*commonProtoc.BalanceResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*commonProtoc.BalanceResponse), err
}

// get tx list
func (s BlockChainRPCServices) GetTxList(ctx context.Context, req *commonProtoc.TxListRequest) (
	*commonProtoc.TxListResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*commonProtoc.TxListResponse), err
}


// get tx detail
func (s BlockChainRPCServices) GetTxDetail(ctx context.Context, req *commonProtoc.TxDetailRequest) (
	*commonProtoc.TxDetailResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*commonProtoc.TxDetailResponse), err
}




