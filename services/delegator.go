package services

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
)

type DelegatorService struct {
}

func (s DelegatorService) GetTotalShares(address string) (
	document.DelegatorShares, errors.IrisError)  {
	
	delegatorShares, err := delegatorModel.GetTotalSharesByAddress(address)
	// can't find record by address
	if err != nil {
		return document.DelegatorShares{}, irisErr
	}
	
	return delegatorShares, irisErr
}

