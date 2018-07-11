package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"golang.org/x/net/context"
)

type IRISHubRPCSERVICES struct {

}

func (s IRISHubRPCSERVICES) GetCandidateList(ctx context.Context, req *irisProtoc.CandidateListRequest) (
	[]*irisProtoc.Candidate, error) {
	
	res, err := Handler(ctx, req)
	return res.([]*irisProtoc.Candidate), err
}

func (s IRISHubRPCSERVICES) GetCandidateDetail(ctx context.Context, req *irisProtoc.CandidateDetailRequest) (
	*irisProtoc.Candidate, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.Candidate), err
}

func (s IRISHubRPCSERVICES) GetDelegatorCandidateList(ctx context.Context, req *irisProtoc.DelegatorCandidateListRequest) (
	[]*irisProtoc.Candidate, error) {
	
	res, err := Handler(ctx, req)
	return res.([]*irisProtoc.Candidate), err
}

func (s IRISHubRPCSERVICES) GetDelegatorTotalShares(ctx context.Context, req *irisProtoc.TotalShareRequest) (
	*irisProtoc.TotalShareResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.TotalShareResponse), err
}



