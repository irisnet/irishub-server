package services

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	"fmt"
	"github.com/irisnet/irishub-server/utils/constants"
	"encoding/json"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/utils/helper"
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

func (s AccountService) GetAccountNum(reqVO vo.AccountNumReqVO) (vo.AccountNumResVO, errors.IrisError) {
	address := reqVO.Address

	var (
		resVO vo.AccountNumResVO
		accRes AccountRes
	)

	uri := fmt.Sprintf(constants.HttpUriGetAccountNum, address)
	statusCode, resBytes := HttpClientGetData(uri)

	if statusCode == constants.StatusCodeBadRequest {
		return resVO, NewIrisErr(errors.EC40001, errors.EM40001, fmt.Errorf(string(resBytes)))
	}


	if err := json.Unmarshal(resBytes, &accRes); err != nil {
		return resVO, ConvertSysErr(err)
	}

	resVO.AccountNum = accRes.Value.AccountNum

	return resVO, irisErr
}

func (s AccountService) GetBalance(reqVO vo.BalanceReqVO) (vo.BalanceResVO, errors.IrisError) {
	address := reqVO.Address
	var (
		resVO vo.BalanceResVO
		accRes AccountRes
	)

	uri := fmt.Sprintf(constants.HttpUriGetBalance, address)
	statusCode, resBytes := HttpClientGetData(uri)

	if statusCode == constants.StatusCodeBadRequest {
		return resVO, NewIrisErr(errors.EC40001, errors.EM40001, fmt.Errorf(string(resBytes)))
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
	address := reqVO.Address
	var (
		resVO  vo.SequenceResVO
		accRes AccountRes
	)

	uri := fmt.Sprintf(constants.HttpUriGetSequence, address)
	statusCode, res := HttpClientGetData(uri)

	// handle nonce is empty
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		resVO.Sequence = 0
		return resVO, irisErr
	}

	err := json.Unmarshal(res, &accRes)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	resVO = vo.SequenceResVO{
		Sequence: accRes.Value.Sequence,
	}

	return resVO, irisErr
}
