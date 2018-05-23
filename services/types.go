package services

import (
	"github.com/irisnet/iris-api-server/models/document"
	"github.com/irisnet/iris-api-server/rests/errors"
)

var (
	candidateModel document.Candidate
	delegatorModel document.Delegator
	stakeTxModel   document.StakeTx
	irisErr        errors.IrisError
)

func ConvertSysErr(err error) errors.IrisError  {
	return irisErr.New(errors.EC50001, err.Error())
}