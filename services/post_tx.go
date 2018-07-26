package services

import (
	"bytes"
	"fmt"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	"encoding/json"
)

type PostTxService struct {
}

// CheckTx result
type ResultBroadcastTx struct {
	Code uint32       `json:"code"`
	Data string       `json:"data"`
	Log  string       `json:"log"`
	Hash string       `json:"hash"`
}

func (s PostTxService) PostTx(reqVO vo.PostTxReqVO) (vo.PostTxResVO, errors.IrisError) {
	var (
		res ResultBroadcastTx
		resVO vo.PostTxResVO
	)

	tx := reqVO.Tx

	reqPostTx := bytes.NewBuffer(tx)
	
	statusCode, resRaw := HttpClientPostJsonData(constants.HttpUriPostTx, reqPostTx)
	
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return resVO, ConvertSysErr(fmt.Errorf(string(resRaw)))
	}

	err := json.Unmarshal(resRaw, &res)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	if res.Code != 0 {
		return resVO, ConvertSysErr(fmt.Errorf(string(resRaw)))
	}

	resVO = vo.PostTxResVO{
		TxHash: res.Hash,
	}
	
	return resVO, irisErr
}

