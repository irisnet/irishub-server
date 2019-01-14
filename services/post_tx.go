package services

import (
	"bytes"
	"encoding/json"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
)

type PostTxService struct {
}

type ResultBroadcastTxCommit struct {
	CheckTx   Response `json:"check_tx,omitempty"`
	DeliverTx Response `json:"deliver_tx,omitempty"`
	Hash      string   `json:"hash,omitempty"`
	Height    string   `json:"height,omitempty"`
}

type Response struct {
	Code      uint32 `json:"code,omitempty"`
	Data      []byte `json:"data,omitempty"`
	Log       string `json:"log,omitempty"`
	Info      string `json:"info,omitempty"`
	GasWanted int64  `json:"gas_wanted,omitempty"`
	GasUsed   int64  `json:"gas_used,omitempty"`
	Codespace string `json:"codespace,omitempty"`
}

func (s PostTxService) PostTx(reqVO vo.PostTxReqVO) (vo.PostTxResVO, errors.IrisError) {
	var resVO vo.PostTxResVO

	tx := reqVO.Tx

	reqPostTx := bytes.NewBuffer(tx)

	hash, err := broadcastTxSync(reqPostTx)
	if err.IsNotNull() {
		return resVO, err
	}

	resVO = vo.PostTxResVO{
		TxHash: hash,
	}

	return resVO, irisErr
}

func broadcastTxSync(requestBody *bytes.Buffer) (hash string, irisErr errors.IrisError) {
	resByte, err := postTxToLCD(false, false, requestBody)
	if err != nil {
		return hash, err.(errors.IrisError)
	}

	var resp ResultBroadcastTxCommit

	er := json.Unmarshal(resByte, &resp)
	if er != nil {
		return resp.Hash, errors.SysErr(err.Error())
	}

	if resp.CheckTx.Code != 0 {
		return resp.Hash, errors.SdkCodeToIrisErr(resp.CheckTx.Code, resp.CheckTx.Log)
	}

	return resp.Hash, irisErr
}
