package services

import (
	"bytes"
	"encoding/json"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
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
	var resVO vo.PostTxResVO

	tx := reqVO.Tx

	reqPostTx := bytes.NewBuffer(tx)

	hash, err := postTx(reqPostTx)
	if err.IsNotNull() {
		return resVO, err
	}

	resVO = vo.PostTxResVO{
		TxHash: hash,
	}

	return resVO, irisErr
}

func postTx(requestBody *bytes.Buffer) (hash string, irisErr errors.IrisError) {
	resByte, err := broadcastTx(false, requestBody)
	if err.IsNotNull() {
		return hash, err
	}

	var resp ResultBroadcastTxCommit

	er := json.Unmarshal(resByte, &resp)
	if er != nil {
		return hash, ConvertSysErr(err)
	}

	if resp.CheckTx.Code != 0 {
		logger.Error.Printf("%v: err is %v\n", "PostTx", resp.CheckTx.Log)
		return hash, NewIrisErr(resp.CheckTx.Code, resp.CheckTx.Log, nil)
	}

	if resp.DeliverTx.Code != 0 {
		logger.Error.Printf("%v: err is %v\n", "PostTx", resp.CheckTx.Log)
		return hash, NewIrisErr(resp.DeliverTx.Code, resp.DeliverTx.Log, nil)
	}

	return resp.Hash, irisErr
}

func postTxAsync(requestBody *bytes.Buffer) (hash string, irisErr errors.IrisError) {
	resByte, err := broadcastTx(true, requestBody)
	if err.IsNotNull() {
		return hash, err
	}

	var resp ResultBroadcastTx

	er := json.Unmarshal(resByte, &resp)
	if er != nil {
		return hash, ConvertSysErr(err)
	}

	if resp.Code != 0 {
		logger.Error.Printf("%v: err is %v\n", "PostTxAsync", resp.Log)
		return hash, NewIrisErr(resp.Code, resp.Log, nil)
	}

	return resp.Hash, irisErr
}
