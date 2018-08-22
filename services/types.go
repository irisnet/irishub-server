package services

import (
	"bytes"
	"io/ioutil"
	"net/http"

	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
)

var (
	candidateModel document.Candidate
	delegatorModel document.Delegator
	commonTxModel  document.CommonTx
	txGasModel     document.TxGas
	valUpTimeModel document.ValidatorUpTime
	irisErr        errors.IrisError
)

func ConvertSysErr(err error) errors.IrisError {
	return irisErr.New(errors.EC50001, errors.EM50001+err.Error())
}

func ConvertBadRequestErr(err error) errors.IrisError {
	return irisErr.New(errors.EC40001, errors.EM40001+err.Error())
}

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
		conf.ServerConfig.AddrNodeServer+uri,
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
	res, err := http.Get(conf.ServerConfig.AddrNodeServer + uri)
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
