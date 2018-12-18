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

type PostTxService struct {
}

// CheckTx result
type ResultBroadcastTx struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
	Log  string `json:"log"`
	Hash string `json:"hash"`
}

type ResultBroadcastTxCommit struct {
	CheckTx   Response `json:"check_tx"`
	DeliverTx Response `json:"deliver_tx"`
	Hash      string   `json:"hash"`
	Height    string   `json:"height"`
}

type Response struct {
	Code      uint32 `json:"code"`
	Data      []byte `json:"data"`
	Log       string `json:"log"`
	Info      string `json:"info"`
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
}

func (s PostTxService) PostTx(reqVO vo.PostTxReqVO) (vo.PostTxResVO, errors.IrisError) {
	var (
		res        ResultBroadcastTxCommit
		resVO      vo.PostTxResVO
		methodName = "PostTx"
	)

	tx := reqVO.Tx

	reqPostTx := bytes.NewBuffer(tx)

	resRaw, err := PostTx(constants.HttpUriPostTx, reqPostTx)
	if err.IsNotNull() {
		return resVO, err
	}

	er := json.Unmarshal(resRaw, &res)
	if er != nil {
		return resVO, ConvertSysErr(err)
	}

	if res.CheckTx.Code != 0 {
		logger.Error.Printf("%v: err is %v\n", methodName, helper.ToJson(res))
		return resVO, NewIrisErr(res.CheckTx.Code, res.CheckTx.Log, nil)
	}

	if res.DeliverTx.Code != 0 {
		logger.Error.Printf("%v: err is %v\n", methodName, helper.ToJson(res))
		return resVO, NewIrisErr(res.DeliverTx.Code, res.DeliverTx.Log, nil)
	}

	resVO = vo.PostTxResVO{
		TxHash: res.Hash,
	}

	return resVO, irisErr
}
