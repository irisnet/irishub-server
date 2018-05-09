package services

import (
	"github.com/irisnet/iris-api-server/models/document"
	"github.com/irisnet/iris-api-server/rests/errors"
)

var (
	candidateModel document.Candidate
	delegatorModel document.Delegator
	irisErr        errors.IrisError
)
