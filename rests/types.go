package rests

import (
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/services"
)

var (
	candidateService services.CandidateService
	irisErr          errors.IrisError
)
