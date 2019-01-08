package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	h "github.com/irisnet/irishub-server/utils/http"
	"io/ioutil"
	"net/http"
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
	syncResult			  document.SyncResult
)

func NewIrisErr(errCode uint32, errMsg string, err error) errors.IrisError {
	if err != nil {
		errMsg = errMsg + err.Error()
	}
	return irisErr.New(errCode, errMsg)
}

// get data use http client
func HttpClientGetData(uri string) (int, []byte) {
	var reqUrl = fmt.Sprintf("%s%s", conf.ServerConfig.LCDServer, uri)
	res, err := http.Get(reqUrl)
	logger.Info.Println(fmt.Sprintf("request uri:%s", reqUrl))
	defer res.Body.Close()

	if err != nil {
		logger.Error.Println(err)
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error.Println(err)
	}
	logger.Info.Println(fmt.Sprintf("hub response data:%s", string(resByte)))

	return res.StatusCode, resByte
}

type SdkError struct {
	CodeSpace string `json:"codespace"`
	Code      uint16 `json:"code"`
	Message   string `json:"message"`
}

func broadcastTx(async, simulate bool, data *bytes.Buffer) (resByte []byte, irisErr errors.IrisError) {
	var uri = fmt.Sprintf(constants.HttpUriPostTxAsync, async, simulate)
	var reqUrl = fmt.Sprintf("%s%s", conf.ServerConfig.LCDServer, uri)
	resp := h.Post(reqUrl, constants.HeaderContentTypeJson, data)
	if resp.Error != nil {
		return nil, errors.SysErr(resp.Error.Error())
	}

	if helper.SliceContains(constants.ErrorStatusCodes, resp.Code) {
		return nil, errors.InvalidParamsErr(resp.Error.Error())
	}

	resByte = resp.Data

	var sdkErr SdkError
	if resp.Code == http.StatusInternalServerError {
		jsonByte, err := helper.ParseJson(resByte)
		if err != nil || len(jsonByte) == 0 {
			//TODO
			if strings.Contains(string(resByte), "already exists") {
				return nil, errors.TxExistedErr(string(resByte))
			} else if strings.Contains(string(resByte), "Timed out") {
				return nil, errors.TxTimeoutErr(string(resByte))
			}
			return nil, errors.SysErr(err.Error())
		}
		err = json.Unmarshal(jsonByte[0], &sdkErr)
		if err != nil {
			return nil, errors.UnKnownErr(err.Error())
		}
		return nil, errors.SdkCodeToIrisErr(sdkErr.CodeSpace, sdkErr.Code, sdkErr.Message)
	}
	return resByte, irisErr
}
