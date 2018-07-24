package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/services"
	"golang.org/x/net/context"
)

var (
	exRateHandle ExRateHandler
	shareHandler ShareHandler
	shareService services.ShareService
	
	candidateListHandler CandidateListHandler
	
	candidateDetailHandler CandidateDetailHandler
	delegatorCandidateListHandler DelegatorCandidateListHandler

	candidateService services.CandidateService
)


func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)
	
	switch req.(type) {
	case *irisProtoc.TotalShareRequest:
		res, err = shareHandler.Handler(ctx, req.(*irisProtoc.TotalShareRequest))
		break
	case *irisProtoc.CandidateListRequest:
		res, err = candidateListHandler.Handler(ctx, req.(*irisProtoc.CandidateListRequest))
		break
	case *irisProtoc.CandidateDetailRequest:
		res, err = candidateDetailHandler.Handler(ctx, req.(*irisProtoc.CandidateDetailRequest))
		break
	case *irisProtoc.DelegatorCandidateListRequest:
		res, err = delegatorCandidateListHandler.Handler(ctx, req.(*irisProtoc.DelegatorCandidateListRequest))
		break
	case *irisProtoc.ExRateRequest:
		res, err = exRateHandle.Handle(ctx, req.(*irisProtoc.ExRateRequest))
		break

	}
	
	return res, err
}