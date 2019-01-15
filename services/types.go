package services

import (
	"bytes"
	"fmt"
	"github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	"github.com/irisnet/irishub-server/utils/http"
	"strings"
)

var (
	candidateModel        document.Candidate
	validatorHistoryModel document.ValidatorHistory
	delegatorModel        document.Delegator
	commonTxModel         document.CommonTx
	txGasModel            document.TxGas
	valUpTimeModel        document.ValidatorUpTime
	irisErr               errors.IrisError
	syncResult            document.SyncResult
)

// get data use http client
func queryFromLCD(uri string) (int, []byte) {
	var reqUrl = fmt.Sprintf("%s%s", configs.ServerConfig.LCDServer, uri)
	resp := http.Get(reqUrl, nil)
	return resp.Code, resp.Data
}

func postTxToLCD(async, simulate bool, data *bytes.Buffer) ([]byte, error) {
	var uri = fmt.Sprintf(constants.HttpUriBroadcastTx, async, simulate)
	var reqUrl = fmt.Sprintf("%s%s", configs.ServerConfig.LCDServer, uri)
	resp := http.Post(reqUrl, constants.HeaderContentTypeJson, data)
	if resp.Error != nil {
		return nil, errors.SysErr(resp.Error.Error())
	}

	if helper.SliceContains(constants.ErrorStatusCodes, resp.Code) {
		return nil, errors.InvalidParamsErr(string(resp.Data))
	}

	if resp.Code == constants.StatusInternalServerError {
		if strings.Contains(string(resp.Data), "Timed out") {
			return nil, errors.TimeoutErr(string(resp.Data))
		}
		return nil, errors.ExtSysUnKnownErr(string(resp.Data))
	}
	return resp.Data, nil
}
