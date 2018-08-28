package services

import (
	"bytes"
	"fmt"

	"encoding/json"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

type PostTxService struct {
}

// CheckTx result
type ResultBroadcastTx struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
	Log  string `json:"log"`
	Hash string `json:"hash"`
}

func (s PostTxService) PostTx(reqVO vo.PostTxReqVO) (vo.PostTxResVO, errors.IrisError) {
	var (
		res        ResultBroadcastTx
		resVO      vo.PostTxResVO
		methodName = "PostTx"
	)

	tx := reqVO.Tx

	reqPostTx := bytes.NewBuffer(tx)

	statusCode, resRaw := HttpClientPostJsonData(constants.HttpUriPostTx, reqPostTx)

	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		logger.Error.Printf("%v: statusCode is %v, err is %v\n",
			methodName, statusCode, string(resRaw))
		return resVO, ConvertSysErr(fmt.Errorf(string(resRaw)))
	}

	err := json.Unmarshal(resRaw, &res)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	if res.Code != 0 {
		logger.Error.Printf("%v: err is %v\n", methodName, helper.ToJson(res))
		return resVO, NewIrisErr(res.Code, res.Log, nil)
	}

	resVO = vo.PostTxResVO{
		TxHash: res.Hash,
	}

	return resVO, irisErr
}
