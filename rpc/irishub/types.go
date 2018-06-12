package irishub

import (
	irisModel "github.com/irisnet/irishub-rpc/codegen/server"
	"github.com/irisnet/irishub-server/services"
	"golang.org/x/net/context"
)

var (
	shareController ShareController
	shareService services.ShareService
	
	candidateListController CandidateListController
	candidateService services.CandidateService
)


func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)
	
	switch req.(type) {
	case *irisModel.TotalShareRequest:
		res, err = shareController.Handler(ctx, req.(*irisModel.TotalShareRequest))
		break
	case *irisModel.CandidateListRequest:
		res, err = candidateListController.Handler(ctx, req.(*irisModel.CandidateListRequest))
		break
	}
	
	return res, err
}