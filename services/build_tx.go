package services

import (
	"bytes"
	"encoding/json"
	
	"github.com/irisnet/iris-api-server/errors"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rpc/vo"
	"github.com/irisnet/iris-api-server/utils/constants"
	"github.com/irisnet/iris-api-server/utils/helper"
)

type BuildTxService struct {

}

func (s BuildTxService) BuildTx(vo vo.BuildTxVO) ([]byte, errors.IrisError) {
	requestBody, err := json.Marshal(vo)
	if err != nil {
		return nil, NewIrisErr(errors.EC40002, errors.EM40002 + err.Error())
	}
	
	reqBuildTx := bytes.NewBuffer([]byte(requestBody))
	statusCode, resBuildTx := HttpClientPostJsonData(constants.HttpUriBuildTx, reqBuildTx)
	
	// http status code isn't ok
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return nil, NewIrisErr(errors.EC40001, errors.EM40001 + string(resBuildTx))
	}
	
	reqByteTx := "{\"tx\": " + string(resBuildTx) + "}"
	logger.Info.Println(reqByteTx)
	reqByteTxData := bytes.NewBuffer([]byte(reqByteTx))
	statusCode, resByteTx := HttpClientPostJsonData(constants.HttpUriByteTx, reqByteTxData)
	
	// http status code isn't success
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return nil, NewIrisErr(errors.EC40001, errors.EM40001 + string(resBuildTx))
	}
	
	return resByteTx, irisErr
}
