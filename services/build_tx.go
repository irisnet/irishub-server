package services

import (
	"bytes"
	"encoding/json"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

type BuildTxService struct {

}

func (s BuildTxService) BuildTx(vo vo.BuildTxReqVO) ([]byte, errors.IrisError) {
	requestBody, err := json.Marshal(vo)
	if err != nil {
		return nil, NewIrisErr(errors.EC40002, errors.EM40002, err)
	}
	
	reqBuildTx := bytes.NewBuffer([]byte(requestBody))
	statusCode, resBuildTx := HttpClientPostJsonData(constants.HttpUriBuildTx, reqBuildTx)
	
	// http status code isn't ok
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return nil, NewIrisErr(errors.EC40001, errors.EM40001, err)
	}
	
	reqByteTx := "{\"tx\": " + string(resBuildTx) + "}"
	logger.Info.Println(reqByteTx)
	reqByteTxData := bytes.NewBuffer([]byte(reqByteTx))
	statusCode, resByteTx := HttpClientPostJsonData(constants.HttpUriByteTx, reqByteTxData)
	
	// http status code isn't success
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return nil, NewIrisErr(errors.EC40001, errors.EM40001 + string(resBuildTx), nil)
	}
	
	return resByteTx, irisErr
}
