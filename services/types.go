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
)

func NewIrisErr(errCode uint32, errMsg string, err error) errors.IrisError {
	if err != nil {
		errMsg = errMsg + err.Error()
	}
	return irisErr.New(errCode, errMsg)
}

func RemoveRepetitionStrValueFromSlice(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// calculate unBond token
func CalculateUnBondToken(coin document.Coin) document.Coin {
	token := coin.Amount * GetShareTokenRatio()
	return document.Coin{
		Amount: token,
		Denom:  constants.Denom,
	}
}

// get ratio of share/token
func GetShareTokenRatio() float64 {
	return 1
}

// post json data use http client
func HttpClientPostJsonData(uri string, requestBody *bytes.Buffer) (int, []byte) {
	res, err := http.Post(
		conf.ServerConfig.LCDServer+uri,
		constants.HeaderContentTypeJson,
		requestBody)
	defer res.Body.Close()

	if err != nil {
		logger.Error.Println(err)
	}

	resByte, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logger.Error.Println(err)
	}

	return res.StatusCode, resByte

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
	logger.Info.Println(fmt.Sprintf("request uri:%s", reqUrl))
	res, err := http.Post(reqUrl,
		constants.HeaderContentTypeJson,
		data)
	if err != nil {
		return nil, errors.SysErr(err)
	}

	resByte, err = ioutil.ReadAll(res.Body)
	logger.Info.Println(fmt.Sprintf("hub response data:%s", string(resByte)))

	if err != nil {
		return nil, errors.SysErr(err)
	}

	statusCode := res.StatusCode
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return nil, errors.InvalidParamsErr(err)
	}

	var sdkErr SdkError
	if statusCode == http.StatusInternalServerError {
		jsonByte, err := helper.ParseJson(resByte)
		if err != nil || len(jsonByte) == 0 {
			//TODO
			if strings.Contains(string(resByte), "already exists") {
				return nil, errors.TxExistedErr(fmt.Errorf(string(resByte)))
			} else if strings.Contains(string(resByte), "Timed out") {
				return nil, errors.TxTimeoutErr(fmt.Errorf(string(resByte)))
			}
			return nil, errors.SysErr(err)
		}
		err = json.Unmarshal(jsonByte[0], &sdkErr)
		if err != nil {
			return nil, errors.UnKnownErr(err)
		}
		return nil, errors.SdkCodeToIrisErr(sdkErr.CodeSpace, sdkErr.Code)
	}
	return resByte, irisErr
}
