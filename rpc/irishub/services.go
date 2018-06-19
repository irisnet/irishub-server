package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"golang.org/x/net/context"
)

type IRISHubRPCSERVICES struct {

}


func (s IRISHubRPCSERVICES) GetDelegatorTotalShares(ctx context.Context, req *irisProtoc.TotalShareRequest) (
	*irisProtoc.TotalShareResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.TotalShareResponse), err
}

func (s IRISHubRPCSERVICES) GetCandidateList(ctx context.Context, req *irisProtoc.CandidateListRequest) (
	*irisProtoc.CandidateListResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.CandidateListResponse), err
}

func (s IRISHubRPCSERVICES) GetCandidateDetail(ctx context.Context, req *irisProtoc.CandidateDetailRequest) (
	*irisProtoc.CandidateDetailResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.CandidateDetailResponse), err
}

func (s IRISHubRPCSERVICES) GetDelegatorCandidateList(ctx context.Context, req *irisProtoc.DelegatorCandidateListRequest) (
	*irisProtoc.DelegatorCandidateListResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisProtoc.DelegatorCandidateListResponse), err
}



