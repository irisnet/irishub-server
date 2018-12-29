package irishub

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/services"
	irisProtoc "github.com/irisnet/irisnet-rpc/irishub/codegen/server/model"
	"golang.org/x/net/context"
)

var (
	validatorListHandler   ValidatorListHandler
	validatorDetailHandler ValidatorDetailHandler
	validatorExRateHandle  ValidatorExRateHandler

	delegatorCandidateListHandler DelegatorCandidateListHandler
	delegatorTotalSharesHandler   DelegatorTotalSharesHandler

	withdrawInfoHandler WithdrawInfoHandler

	validatorService services.ValidatorService
	delegatorService services.DelegatorService
	accountService   services.AccountService
)

func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)

	switch req.(type) {
	case *irisProtoc.CandidateListRequest:
		res, err = validatorListHandler.Handler(ctx, req.(*irisProtoc.CandidateListRequest))
		break
	case *irisProtoc.CandidateDetailRequest:
		res, err = validatorDetailHandler.Handler(ctx, req.(*irisProtoc.CandidateDetailRequest))
		break
	case *irisProtoc.ValidatorExRateRequest:
		res, err = validatorExRateHandle.Handle(ctx, req.(*irisProtoc.ValidatorExRateRequest))
		break

	case *irisProtoc.DelegatorCandidateListRequest:
		res, err = delegatorCandidateListHandler.Handler(ctx, req.(*irisProtoc.DelegatorCandidateListRequest))
		break
	case *irisProtoc.TotalShareRequest:
		res, err = delegatorTotalSharesHandler.Handler(ctx, req.(*irisProtoc.TotalShareRequest))
		break
	case *irisProtoc.WithdrawAddrRequest:
		res, err = withdrawInfoHandler.Handler(ctx, req.(*irisProtoc.WithdrawAddrRequest))
		break
	}

	return res, err
}

func BuildException(err errors.IrisError) error {
	var (
		exception irisProtoc.Exception
	)
	exception.ErrCode = int32(err.ErrCode)
	exception.ErrMsg = err.ErrMsg
	return &exception
}
