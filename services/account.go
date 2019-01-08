package services

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	"strconv"
)

type AccountService struct {
}

type AccountRes struct {
	Coins      []string `json:"coins"`
	AccountNum string   `json:"account_number"`
	Sequence   string   `json:"sequence"`
}

type coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func (s AccountService) GetBalance(reqVO vo.BalanceReqVO) (vo.BalanceResVO, errors.IrisError) {
	var (
		resVO      vo.BalanceResVO
		accRes     AccountRes
		methodName = "GetBalance"
	)

	address := reqVO.Address

	uri := fmt.Sprintf(constants.HttpUriGetBalance, address)
	statusCode, resBytes := HttpClientGetData(uri)

	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		logger.Error.Printf("%v: statusCode is %v, err is %v\n",
			methodName, statusCode, string(resBytes))
		return resVO, errors.SysErr(string(resBytes))
	}

	if statusCode == constants.StatusCodeNotContent {
		return resVO, irisErr
	}

	if err := json.Unmarshal(resBytes, &accRes); err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, errors.SysErr(err.Error())
	}

	var coins []*vo.Coin
	for _, str := range accRes.Coins {
		demon, amt, err := helper.ParseCoin(str)
		if err == nil {
			coin := vo.Coin{
				Denom:  demon,
				Amount: helper.ConvertStrToFloat(amt)}
			destCoin := coin.Covert(vo.CoinTypeAtto)
			coins = append(coins, &destCoin)
		}
	}

	resVO = vo.BalanceResVO{
		Data: vo.BalanceResDataVO{
			Coins: coins,
		},
	}

	return resVO, irisErr
}

func (s AccountService) GetSequence(reqVO vo.SequenceReqVO) (vo.SequenceResVO, errors.IrisError) {
	var (
		resVO      vo.SequenceResVO
		accRes     AccountRes
		err        error
		methodName = "GetSequence"
	)

	address := reqVO.Address

	uri := fmt.Sprintf(constants.HttpUriGetSequence, address)
	statusCode, res := HttpClientGetData(uri)

	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		logger.Error.Printf("%v: statusCode is %v, err is %v\n",
			methodName, statusCode, string(res))
		return resVO, errors.SysErr(string(res))
	}

	// handle nonce is empty
	if statusCode == constants.StatusCodeNotContent {
		resVO.Sequence = 0
		resVO.Ext = []byte(strconv.Itoa(int(0)))
		return resVO, irisErr
	}

	err = json.Unmarshal(res, &accRes)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, errors.SysErr(err.Error())
	}

	resVO = vo.SequenceResVO{
		Sequence: helper.ConvertStrToInt(accRes.Sequence),
		Ext:      []byte(strconv.Itoa(int(helper.ConvertStrToInt(accRes.AccountNum)))),
	}

	return resVO, irisErr
}

func (s AccountService) GetRewardInfo(req vo.RewardInfoReqVO) (res vo.RewardInfoResVo, e errors.IrisError) {
	txList := commonTxModel.GetRewardList(req.DelAddr)
	return vo.RewardInfoResVo{
		DelAddr: req.DelAddr,
		Txs:     txList,
	}, irisErr
}

func (s AccountService) QueryWithdrawAddr(delAddr string) (string, errors.IrisError) {
	//查询用户的提现地址
	uri := fmt.Sprintf(constants.HttpUriGetWithdrawAddr, delAddr)
	statusCode, res := HttpClientGetData(uri)
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		logger.Error.Printf("%v: statusCode is %v, err is %v\n",
			"queryWithdrawAddr", statusCode, string(res))
		return "", errors.SysErr(string(res))
	}
	if statusCode == constants.StatusCodeNotContent {
		return delAddr, irisErr
	}
	return string(res), irisErr
}
