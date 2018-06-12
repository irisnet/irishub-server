package irishub

import (
	irisModel "github.com/irisnet/irishub-rpc/codegen/server"
	"golang.org/x/net/context"
)

type IRISHubRPCSERVICES struct {

}


func (s IRISHubRPCSERVICES) GetDelegatorTotalShares(ctx context.Context, req *irisModel.TotalShareRequest) (
	*irisModel.TotalShareResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisModel.TotalShareResponse), err
}

func (s IRISHubRPCSERVICES) GetCandidateList(ctx context.Context, req *irisModel.CandidateListRequest) (
	*irisModel.CandidateListResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisModel.CandidateListResponse), err
}

func (s IRISHubRPCSERVICES) GetCandidateDetail(ctx context.Context, req *irisModel.CandidateDetailRequest) (
	*irisModel.CandidateDetailResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisModel.CandidateDetailResponse), err
}

func (s IRISHubRPCSERVICES) GetDelegatorCandidateList(ctx context.Context, req *irisModel.DelegatorCandidateListRequest) (
	*irisModel.DelegatorCandidateListResponse, error) {
	
	res, err := Handler(ctx, req)
	return res.(*irisModel.DelegatorCandidateListResponse), err
}



