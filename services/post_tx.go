package services

import (
	"bytes"
	"encoding/json"
	
	"github.com/irisnet/iris-api-server/errors"
	"github.com/irisnet/iris-api-server/rpc/vo"
	"github.com/irisnet/iris-api-server/utils/constants"
	"github.com/irisnet/iris-api-server/utils/helper"
)

type PostTxService struct {
}

func (s PostTxService) PostTx(vo vo.PostTxReqVO) ([]byte, errors.IrisError) {
	requestBody, err := json.Marshal(vo)
	if err != nil {
		return nil, NewIrisErr(errors.EC40002, errors.EM40002 + err.Error())
	}
	
	reqPostTx := bytes.NewBuffer([]byte(requestBody))
	
	statusCode, res := HttpClientPostJsonData(constants.HttpUriPostTx, reqPostTx)
	
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return nil, NewIrisErr(errors.EC40001, errors.EM40001 + err.Error())
	}
	
	return res, irisErr
}

