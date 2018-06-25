package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

type BuildTxService struct {

}

func (s BuildTxService) BuildTx(reqVO vo.BuildTxReqVO) (vo.BuildTxResVO, errors.IrisError) {
	var (
		err error
		resBuildTx []byte
		resVO vo.BuildTxResVO
	)
	
	buildTxFunc := func(uri string, reqVO interface{}) ([]byte, error) {
		requestBody, err := json.Marshal(reqVO)
		
		if err != nil {
			return nil, err
		}
		
		reqBuildTx := bytes.NewBuffer([]byte(requestBody))
		statusCode, res := HttpClientPostJsonData(uri, reqBuildTx)
		
		// http status code isn't ok
		if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
			return nil, fmt.Errorf(string(res))
		}
		
		return res, nil
	}
	
	txType := reqVO.TxType
	switch txType {
	case constants.DbTxTypeCoin:
		resBuildTx, err = buildTxFunc(constants.HttpUriBuildCoinTx, reqVO)
		break
	case constants.DbTxTypeStakeDelegate:
		delegateReqVO := vo.BuildDelegateTxReqVO{
			Fees: reqVO.Fees,
			Sequence: reqVO.Sequence,
			From: reqVO.From,
			PubKey: vo.PubKey{
				Type: "ed25519",
				Data: reqVO.To.Addr,
			},
			Amount: reqVO.Amount[0],
		}
		resBuildTx, err = buildTxFunc(constants.HttpUriBuildDelegateTx, delegateReqVO)
		break
	case constants.DbTxTypeStakeUnBond:
		unBondReqVO := vo.BuildUnBondTxReqVO{
			Fees: reqVO.Fees,
			Sequence: reqVO.Sequence,
			From: reqVO.From,
			PubKey: vo.PubKey{
				Type: "ed25519",
				Data: reqVO.To.Addr,
			},
			Amount: reqVO.Amount[0].Amount,
		}
		resBuildTx, err = buildTxFunc(constants.HttpUriBuildUnBondTx, unBondReqVO)
		break
	default:
		logger.Info.Println("unkoow tx type")
	}
	
	if err != nil {
		return resVO, NewIrisErr(errors.EC40001, errors.EM40001, err)
	}
	
	reqByteTx := "{\"tx\": " + string(resBuildTx) + "}"
	reqByteTxData := bytes.NewBuffer([]byte(reqByteTx))
	statusCode, resByteTx := HttpClientPostJsonData(constants.HttpUriByteTx, reqByteTxData)
	
	// http status code isn't success
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return resVO, NewIrisErr(errors.EC40001, errors.EM40001 + string(resBuildTx), nil)
	}
	resVO.Data = resByteTx
	resVO.Ext = resBuildTx
	return resVO, irisErr
}
