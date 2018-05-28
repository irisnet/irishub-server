package services

import (
	"github.com/irisnet/iris-api-server/models/document"
	"github.com/irisnet/iris-api-server/rests/errors"
)

var (
	candidateModel document.Candidate
	delegatorModel document.Delegator
	stakeTxModel   document.StakeTx
	commonTxModel  document.CommonTx
	irisErr        errors.IrisError
)

func ConvertSysErr(err error) errors.IrisError  {
	return irisErr.New(errors.EC50001, err.Error())
}

func RemoveRepetitionStrValueFromSlice(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	
	return list
}