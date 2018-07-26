package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"golang.org/x/net/context"
)

type IRISHubRPCServices struct {

}

func (s IRISHubRPCServices) GetCandidateList(ctx context.Context, req *irisProtoc.CandidateListRequest) (
	[]*irisProtoc.Candidate, error) {
	
	res, err := Handler(ctx, req)
	return res.([]*irisProtoc.Candidate), err
}

func (s IRISHubRPCServices) GetCandidateDetail(ctx context.Context, req *irisProtoc.CandidateDetailRequest) (
	*irisProtoc.Candidate, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.Candidate), err
}

func (s IRISHubRPCServices) GetDelegatorCandidateList(ctx context.Context, req *irisProtoc.DelegatorCandidateListRequest) (
	[]*irisProtoc.Candidate, error) {
	
	res, err := Handler(ctx, req)
	return res.([]*irisProtoc.Candidate), err
}

func (s IRISHubRPCServices) GetDelegatorTotalShares(ctx context.Context, req *irisProtoc.TotalShareRequest) (
	*irisProtoc.TotalShareResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.TotalShareResponse), err
}

func (s IRISHubRPCServices)GetValidatorExRate(ctx context.Context, req *irisProtoc.ValidatorExRateRequest) (
	r *irisProtoc.ValidatorExRateResponse, err error) {

	res, err := Handler(ctx, req)
	return res.(*irisProtoc.ValidatorExRateResponse), err
}



