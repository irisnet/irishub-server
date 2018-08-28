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
		startTime time.Time
		endTime   time.Time
		err       error
		//valAddrs   []string
		candidates []document.Candidate
		resVO      vo.TxListResVO
		methodName = "GetTxList"
	)

	address := reqVO.Address
	txType := reqVO.Type
	skip, limit := helper.ParseParamPage(int(reqVO.Page), int(reqVO.PerPage))
	sorts := helper.ParseParamSort(reqVO.Sort)
	ext := string(reqVO.Ext)

	if reqVO.StartTime != "" {
		startTime, err = helper.ParseFullTime(reqVO.StartTime)
		if err != nil {
			logger.Error.Printf("%v: err is %v\n", methodName, err)
			return resVO, ConvertBadRequestErr(err)
		}
	}

	if reqVO.EndTime != "" {
		endTime, err = helper.ParseFullTime(reqVO.EndTime)
		if err != nil {
			logger.Error.Printf("%v: err is %v\n", methodName, err)
			return resVO, ConvertBadRequestErr(err)
		}
	}

	commonTxs, err := commonTxModel.GetList(address, txType, startTime, endTime,
		skip, limit, sorts, ext)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
	}

	//for _, commonTx := range commonTxs {
	//	txType := commonTx.Type
	//	if txType == constants.DbTxTypeStakeDelegate ||
	//		txType == constants.DbTxTypeStakeBeginUnBonding ||
	//		txType == constants.DbTxTypeStakeCompleteUnBonding {
	//		valAddrs = append(valAddrs, commonTx.To)
	//	}
	//}

	for i, commonTx := range commonTxs {
		commonTx = s.buildData(commonTx, candidates, address)
		commonTxs[i] = commonTx
	}
	resVO.Txs = commonTxs

	return resVO, irisErr
}

func (s TxService) GetTxGas(reqVO vo.TxGasReqVO) (vo.TxGasResVO, errors.IrisError) {
	var (
		resVO      vo.TxGasResVO
		methodName = "GetTxGas"
	)

	txGas, err := txGasModel.GetTxGas(reqVO.TxType)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
	}

	if txGas.TxType == "" {
		// can't query tx gasUsed and gasPrice by valid txType
		resVO = s.buildDefaultTxGas(reqVO.TxType)
	} else {
		resVO = vo.TxGasResVO{
			TxType:   txGas.TxType,
			Gas:      txGas.GasUsed,
			GasPrice: txGas.GasPrice,
		}
	}

	return resVO, irisErr
}

func (s TxService) buildDefaultTxGas(txType string) vo.TxGasResVO {
	var (
		defaultGasUsed float64
	)
	switch txType {
	case constants.TxTypeCoinSend:
		defaultGasUsed = constants.DefaultTxGasTransfer
		break
	case constants.TxTypeStakeDelegate:
		defaultGasUsed = constants.DefaultTxGasDelegate
		break
	case constants.TxTypeStakeBeginUnBonding:
		defaultGasUsed = constants.DefaultTxGasBeginUbonding
	case constants.TxTypeStakeCompleteUnBonding:
		defaultGasUsed = constants.DefaultTxGasCompleteUnbonding
		break
	default:
		return vo.TxGasResVO{}
	}

	return vo.TxGasResVO{
		TxType: constants.TxTypeFrontMapDb[txType],
		Gas: document.GasUsed{
			MinGasUsed: defaultGasUsed,
			AvgGasUsed: defaultGasUsed,
			MaxGasUsed: defaultGasUsed,
		},
		GasPrice: document.GasPrice{
			Denom:       "iris",
			MinGasPrice: constants.DefaultMinGasPrice,
			AvgGasPrice: constants.DefaultAvgGasPrice,
			MaxGasPrice: constants.DefaultMaxGasPrice,
		},
	}
}

func (s TxService) GetTxDetail(reqVO vo.TxDetailReqVO) (vo.TxDetailResVO, errors.IrisError) {
	var (
		resVO      vo.TxDetailResVO
		pubKeys    []string
		candidates []document.Candidate
	)

	// get tx detail by txHash
	commonTx, err := commonTxModel.GetDetail(reqVO.TxHash)

	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	if commonTx.TxHash != "" {
		txType := commonTx.Type
		if txType == constants.TxTypeStakeDelegate || txType == constants.TxTypeStakeBeginUnBonding {
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
		txType == constants.DbTxTypeStakeBeginUnBonding ||
		txType == constants.DbTxTypeStakeCompleteUnBonding {

		for _, candidate := range candidates {
			if commonTx.To == candidate.Address {
				commonTx.Candidate = candidate
				break
			}
		}
	}

	// display tx type used by front
	switch txType {
	case constants.DbTxTypeTransfer:
		if address == commonTx.From {
			txTypeDisplay = constants.TxTypeCoinSend
		} else if address == commonTx.To {
			txTypeDisplay = constants.TxTypeCoinReceive
		}
		break
	case constants.DbTxTypeStakeDelegate:
		txTypeDisplay = constants.TxTypeStakeDelegate
		break
	case constants.DbTxTypeStakeBeginUnBonding:
		txTypeDisplay = constants.TxTypeStakeBeginUnBonding
		break
	case constants.DbTxTypeStakeCompleteUnBonding:
		txTypeDisplay = constants.TxTypeStakeCompleteUnBonding
	default:
		logger.Info.Println("unsupported tx type")
	}

	commonTx.Type = txTypeDisplay

	return commonTx
}
