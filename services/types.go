package services

import (
	"bytes"
	"io/ioutil"
	"net/http"
	
	conf "github.com/irisnet/iris-api-server/configs"
	"github.com/irisnet/iris-api-server/models/document"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rests/errors"
	irisErrRPC "github.com/irisnet/iris-api-server/rpc/errors"
	"github.com/irisnet/iris-api-server/utils/constants"
)

var (
	candidateModel document.Candidate
	delegatorModel document.Delegator
	stakeTxModel   document.StakeTx
	commonTxModel  document.CommonTx
	irisErr        errors.IrisError
	irisError      irisErrRPC.IrisError
)

func ConvertSysErr(err error) errors.IrisError  {
	return irisErr.New(errors.EC50001, err.Error())
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
		Denom: constants.Denom,
	}
}

// get ratio of share/token
func GetShareTokenRatio() int64 {
	return 1
}

// post json data use http client
func HttpClientPostJsonData(uri string, requestBody *bytes.Buffer) (int, []byte) {
	res, err := http.Post(
		conf.ServerConfig.AddrNodeServer + uri,
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