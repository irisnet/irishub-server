package blockchain

import (
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
	"golang.org/x/net/context"
)

type BlockChainRPCServices struct {
}

func (s BlockChainRPCServices) GetTxGas(ctx context.Context, req *commonProtoc.TxGasRequest) (
	r *commonProtoc.TxGasResponse, err error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.TxGasResponse{}, err
	}
	return res.(*commonProtoc.TxGasResponse), err
}

// get sequence
func (s BlockChainRPCServices) GetSequence(ctx context.Context, req *commonProtoc.SequenceRequest) (
	*commonProtoc.SequenceResponse, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.SequenceResponse{}, err
	}
	return res.(*commonProtoc.SequenceResponse), err
}

// build tx
func (s BlockChainRPCServices) BuildTx(ctx context.Context, req *commonProtoc.BuildTxRequest) (
	*commonProtoc.BuildTxResponse, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.BuildTxResponse{}, err
	}
	return res.(*commonProtoc.BuildTxResponse), err
}

// post tx
func (s BlockChainRPCServices) PostTx(ctx context.Context, req *commonProtoc.PostTxRequest) (
	*commonProtoc.PostTxResponse, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.PostTxResponse{}, err
	}
	return res.(*commonProtoc.PostTxResponse), err
}

// post tx
func (s BlockChainRPCServices) SimulateTx(ctx context.Context, req *commonProtoc.SimulateTxRequest) (
	*commonProtoc.SimulateTxResponse, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.SimulateTxResponse{}, err
	}
	return res.(*commonProtoc.SimulateTxResponse), err
}

// get balance
func (s BlockChainRPCServices) GetBalance(ctx context.Context, req *commonProtoc.BalanceRequest) (
	*commonProtoc.BalanceResponse, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.BalanceResponse{}, err
	}
	return res.(*commonProtoc.BalanceResponse), err
}

// get tx list
func (s BlockChainRPCServices) GetTxList(ctx context.Context, req *commonProtoc.TxListRequest) (
	[]*commonProtoc.Tx, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return []*commonProtoc.Tx{}, err
	}
	return res.([]*commonProtoc.Tx), err
}

// get tx detail
func (s BlockChainRPCServices) GetTxDetail(ctx context.Context, req *commonProtoc.TxDetailRequest) (
	*commonProtoc.Tx, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.Tx{}, err
	}
	return res.(*commonProtoc.Tx), err
}

// GetRewardInfo
func (s BlockChainRPCServices) GetRewardInfo(ctx context.Context, req *commonProtoc.RewardInfoRequest) (
	*commonProtoc.RewardInfoResponse, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &commonProtoc.RewardInfoResponse{}, err
	}
	return res.(*commonProtoc.RewardInfoResponse), err
}
