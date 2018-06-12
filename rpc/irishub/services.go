package irishub

import (
	irisModel "github.com/irisnet/irishub-rpc/codegen/server"
	"golang.org/x/net/context"
)

type IRISHubRPCSERVICES struct {

}


func (s IRISHubRPCSERVICES) GetDelegatorTotalShares(ctx context.Context, req *irisModel.TotalShareRequest) (
	*irisModel.TotalShareResponse, error) {
	
	return nil, nil
}
func (s IRISHubRPCSERVICES) GetCandidateList(ctx context.Context, req *irisModel.CandidateListRequest) (
	*irisModel.CandidateListResponse, error) {
	
	return nil, nil
}
func (s IRISHubRPCSERVICES) GetCandidateDetail(ctx context.Context, req *irisModel.CandidateDetailRequest) (
	*irisModel.CandidateDetailResponse, error) {
	
	return nil, nil
}
func (s IRISHubRPCSERVICES) GetDelegatorCandidateList(ctx context.Context, req *irisModel.DelegatorCandidateListRequest) (
	*irisModel.DelegatorCandidateListResponse, error) {
	
	return nil, nil
}



