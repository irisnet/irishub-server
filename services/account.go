package services

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	"fmt"
	"github.com/irisnet/irishub-server/utils/constants"
	"encoding/json"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/bech32"
	"strconv"
)

type AccountService struct {

}

type AccountRes struct {
	Type string `json:"type"`
	Value Value `json:"value"`
}

type Value struct {
	Coins []*vo.Coin `json:"coins"`
	AccountNum int64 `json:"account_number"`
	Sequence int64 `json:"sequence"`
}

func (s AccountService) GetBalance(reqVO vo.BalanceReqVO) (vo.BalanceResVO, errors.IrisError) {
	var (
		resVO vo.BalanceResVO
		accRes AccountRes
	)

	address, err := bech32.ConvertHexToBech32(reqVO.Address)
	if err != nil {
		return resVO, ConvertBadRequestErr(err)
	}

	uri := fmt.Sprintf(constants.HttpUriGetBalance, address)
	statusCode, resBytes := HttpClientGetData(uri)

	if statusCode == constants.StatusCodeNotContent {
		return resVO, irisErr
	}


	if err := json.Unmarshal(resBytes, &accRes); err != nil {
		return resVO, ConvertSysErr(err)
	}

	resVO = vo.BalanceResVO{
		Data: vo.BalanceResDataVO{
			Coins: accRes.Value.Coins,
		},
	}

	return resVO, irisErr
}

func (s AccountService) GetSequence(reqVO vo.SequenceReqVO) (vo.SequenceResVO, errors.IrisError) {
	var (
		resVO  vo.SequenceResVO
		accRes AccountRes
	)

	address, err := bech32.ConvertHexToBech32(reqVO.Address)
	if err != nil {
		return resVO, ConvertBadRequestErr(err)
	}

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
		Sequence: accRes.Value.Sequence,
		Ext: []byte(strconv.Itoa(int(accRes.Value.AccountNum))),
	}

	return resVO, irisErr
}
