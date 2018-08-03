package services

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	"fmt"
	"github.com/irisnet/irishub-server/utils/constants"
	"encoding/json"
	"github.com/irisnet/irishub-server/errors"
	"strconv"
	"github.com/irisnet/irishub-server/utils/helper"
)

type AccountService struct {

}

type AccountRes struct {
	Type string `json:"type"`
	Value Value `json:"value"`
}

type Value struct {
	Coins []coin `json:"coins"`
	AccountNum string `json:"account_number"`
	Sequence string `json:"sequence"`
}

type coin struct {
	Denom  string `json:"denom"`
	Amount string  `json:"amount"`
}

func (s AccountService) GetBalance(reqVO vo.BalanceReqVO) (vo.BalanceResVO, errors.IrisError) {
	var (
		resVO vo.BalanceResVO
		accRes AccountRes
	)

	address := reqVO.Address

	uri := fmt.Sprintf(constants.HttpUriGetBalance, address)
	statusCode, resBytes := HttpClientGetData(uri)

	if statusCode == constants.StatusCodeNotContent {
		return resVO, irisErr
	}


	if err := json.Unmarshal(resBytes, &accRes); err != nil {
		return resVO, ConvertSysErr(err)
	}

	funBuildCoins := func(coins []coin) []*vo.Coin {
		var (
			resCoins []*vo.Coin
		)

		if len(coins) > 0 {
			for _, v := range coins {
				coin := vo.Coin{
					Denom: v.Denom,
					Amount: helper.ConvertStrToInt(v.Amount),
				}

				resCoins = append(resCoins, &coin)
			}
		}

		return resCoins
	}

	resVO = vo.BalanceResVO{
		Data: vo.BalanceResDataVO{
			Coins: funBuildCoins(accRes.Value.Coins),
		},
	}

	return resVO, irisErr
}

func (s AccountService) GetSequence(reqVO vo.SequenceReqVO) (vo.SequenceResVO, errors.IrisError) {
	var (
		resVO  vo.SequenceResVO
		accRes AccountRes
		err error
	)

	address := reqVO.Address

	uri := fmt.Sprintf(constants.HttpUriGetSequence, address)
	statusCode, res := HttpClientGetData(uri)

	// handle nonce is empty
	if statusCode == constants.StatusCodeNotContent {
		resVO.Sequence = 0
		resVO.Ext = []byte(strconv.Itoa(int(0)))
		return resVO, irisErr
	}

	err = json.Unmarshal(res, &accRes)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	resVO = vo.SequenceResVO{
		Sequence: helper.ConvertStrToInt(accRes.Value.Sequence),
		Ext: []byte(strconv.Itoa(int(helper.ConvertStrToInt(accRes.Value.AccountNum)))),
	}

	return resVO, irisErr
}
