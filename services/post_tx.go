package services

import (
	"bytes"
	"fmt"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

type PostTxService struct {
}

func (s PostTxService) PostTx(reqVO vo.PostTxReqVO) ([]byte, errors.IrisError) {
	tx := reqVO.Tx

	reqPostTx := bytes.NewBuffer(tx)
	
	statusCode, res := HttpClientPostJsonData(constants.HttpUriPostTx, reqPostTx)
	
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return nil, NewIrisErr(errors.EC40001, errors.EM40001, fmt.Errorf(string(res)))
	}
	
	return res, irisErr
}

