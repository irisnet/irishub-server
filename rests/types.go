package rests

import (
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/services"
)

var (
	candidateService services.CandidateService
	delegatorService services.DelegatorService
	stakeTxService   services.StakeTxService
	irisErr          errors.IrisError
)
