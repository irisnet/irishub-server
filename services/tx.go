package services

import (
	"time"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

type TxService struct {

}

func (s TxService) GetTxList(reqVO vo.TxListReqVO) (vo.TxListResVO, errors.IrisError) {
	var (
		startTime  time.Time
		endTime    time.Time
		err        error
		valAddrs   []string
		candidates []document.Candidate
		resVO      vo.TxListResVO
	)
	
	address := reqVO.Address
	txType := reqVO.Type
	skip, limit := helper.ParseParamPage(int(reqVO.Page), int(reqVO.PerPage))
	sorts := helper.ParseParamSort(reqVO.Sort)
	
	if reqVO.StartTime != "" {
		startTime, err = helper.ParseFullTime(reqVO.StartTime)
		if err != nil {
			return resVO, NewIrisErr(errors.EC40002, errors.EM40002, err)
		}
	}
	
	if reqVO.EndTime != "" {
		endTime, err = helper.ParseFullTime(reqVO.EndTime)
		if err != nil {
			return resVO, NewIrisErr(errors.EC40002, errors.EM40002, err)
		}
	}
	
	commonTxs, err := commonTxModel.GetList(address, txType, startTime, endTime,
		skip, limit, sorts)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}
	
	
	for _, commonTx := range commonTxs {
		txType := commonTx.Type
		if txType == constants.TxTypeStakeDelegate || txType == constants.TxTypeStakeUnBond {
			valAddrs = append(valAddrs, commonTx.To)
		}
	}
	
	// remove repetition value in slice
	valAddrs = RemoveRepetitionStrValueFromSlice(valAddrs)
	
	// get candidates by valAddrs
	if valAddrs != nil {
		candidates, err = candidateModel.GetCandidatesListByValidatorAddrs(valAddrs)
		if err != nil {
			return resVO, ConvertSysErr(err)
		}
	}
	
	for i, commonTx := range commonTxs {
		commonTx = s.buildData(commonTx, candidates, address)
		commonTxs[i] = commonTx
	}
	resVO.Txs = commonTxs
	
	
	return resVO, irisErr
}

func (s TxService) GetTxDetail(reqVO vo.TxDetailReqVO) (vo.TxDetailResVO, errors.IrisError) {
	var (
		resVO vo.TxDetailResVO
		pubKeys []string
		candidates []document.Candidate
	)
	
	// get tx detail by txHash
	commonTx, err := commonTxModel.GetDetail(reqVO.TxHash)
	
	if err != nil {
		return resVO, ConvertSysErr(err)
	}
	
	if commonTx.TxHash != "" {
		txType := commonTx.Type
		if txType == constants.TxTypeStakeDelegate || txType == constants.TxTypeStakeUnBond {
			pubKeys = append(pubKeys, commonTx.To)
		}
	}
	
	// get candidates by pubKeys
	if pubKeys != nil {
		candidates, err = candidateModel.GetCandidatesListByValidatorAddrs(pubKeys)
		if err != nil {
			return resVO, ConvertSysErr(err)
		}
	}
	
	commonTx = s.buildData(commonTx, candidates, "")
	resVO.Tx = commonTx
	
	return resVO, irisErr
}

func (s TxService) buildData(commonTx document.CommonTx,
	candidates []document.Candidate, address string) document.CommonTx {
	
	var (
		txTypeDisplay string // tx type display in front
	)
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
		} else if address == commonTx.To {
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
