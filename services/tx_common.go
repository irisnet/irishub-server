package services

import (
	"time"
	
	"github.com/irisnet/iris-api-server/models/document"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
	"github.com/irisnet/iris-api-server/utils/constants"
	"github.com/irisnet/iris-api-server/utils/helper"
)

type CommonTxService struct {
}

func (s CommonTxService) GetList(vo vo.CommonTxListVO) (
	[]document.CommonTx, errors.IrisError) {
	
	var (
		startTime time.Time
		endTime time.Time
		err error
		pubKeys []string
		candidates []document.Candidate
	)
	
	address := vo.Address
	txType := vo.TxType
	skip, limit := helper.ParseParamPage(vo.Page, vo.PerPage)
	sorts := helper.ParseParamSort(vo.Sort)
	
	if vo.StartTime != "" {
		startTime, err = helper.ParseFullTime(vo.StartTime)
		if err != nil {
			return nil, irisErr.New(errors.EC40002, errors.EM40002)
		}
	}
	
	if vo.EndTime != "" {
		endTime, err = helper.ParseFullTime(vo.EndTime)
		if err != nil {
			return nil, irisErr.New(errors.EC40002, errors.EM40002)
		}
	}
	
	commonTxs, err := commonTxModel.GetList(address, txType, startTime, endTime,
		skip, limit, sorts)
	if err != nil {
		return nil, ConvertSysErr(err)
	}
	
	
	for _, commonTx := range commonTxs {
		txType := commonTx.Type
		if txType == constants.TxTypeStakeDelegate || txType == constants.TxTypeStakeUnBond {
			pubKeys = append(pubKeys, commonTx.To)
		}
	}
	
	// remove repetition value in slice
	pubKeys = RemoveRepetitionStrValueFromSlice(pubKeys)
	
	// get candidates by pubKeys
	if pubKeys != nil {
		candidates, err = candidateModel.GetCandidatesListByPubKeys(pubKeys)
		if err != nil {
			return nil, ConvertSysErr(err)
		}
	}
	
	for i, commonTx := range commonTxs {
		commonTx = s.buildData(commonTx, candidates, address)
		commonTxs[i] = commonTx
	}
	
	
	return commonTxs, irisErr
}

func (s CommonTxService) buildData(commonTx document.CommonTx,
	candidates []document.Candidate, address string) document.CommonTx {
	
	var txTypeDisplay string // tx type display for front
	txType := commonTx.Type
	
	// get candidate info
	if txType == constants.DbTxTypeStakeDelegate ||
		txType == constants.DbTxTypeStakeUnBond {
		
		pubKey := commonTx.To
		for _, candidate := range candidates {
			if pubKey == candidate.PubKey {
				commonTx.Candidate = candidate
				break
			}
		}
	}
	
	switch txType {
	case constants.DbTxTypeCoin:
		if address == commonTx.From {
			txTypeDisplay = constants.TxTypeCoinSend
		} else {
			txTypeDisplay = constants.TxTypeCoinReceive
		}
		break
	case constants.DbTxTypeStakeDelegate:
		txTypeDisplay = constants.TxTypeStakeDelegate
		break
	case constants.DbTxTypeStakeUnBond:
		txTypeDisplay = constants.TxTypeStakeUnBond
		commonTx.Amount[0] = CalculateUnBondToken(commonTx.Amount[0])
		break
	default:
		logger.Info.Println("unsupport tx type")
	}
	
	commonTx.Type = txTypeDisplay
	
	return commonTx
}
