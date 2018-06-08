package rests

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/services"
)

var (
	candidateService services.CandidateService
	delegatorService services.DelegatorService
	stakeTxService   services.StakeTxService
	commonTxService  services.CommonTxService
	irisErr          errors.IrisError
)
