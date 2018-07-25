package services

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/modules/bech32"
	"fmt"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	"encoding/json"
	"math"
)

type ShareService struct {

}

func (s ShareService) GetDelegatorTotalShare(reqVO vo.ShareReqVO) (vo.ShareResVO, errors.IrisError) {
	
	var (
		resVO vo.ShareResVO
		totalShares float64
	)
	
	delegatorShares, err := delegatorModel.GetTotalSharesByAddress(reqVO.Address)
	// can't find record by address
	if err != nil {
		return resVO, irisErr
	}

	if len(delegatorShares) > 0 {
		for _, v := range delegatorShares {
			// get exchange rate
			reqVO := vo.ExRateReqVO{
				ValidatorAddress: v.ValidatorAddr,
			}
			res, err := s.GetExRate(reqVO)

			if err.IsNotNull() {
				return resVO, err
			}
			if res.ExRate != "" {
				rate, err2 := helper.ConvertRatStrToFloat(res.ExRate)
				if err2 != nil {
					return resVO, ConvertSysErr(err2)
				}
				totalShares += v.TotalShares * rate
			}
		}
	}
	
	resVO = vo.ShareResVO{
		TotalShare: uint64(math.Floor(totalShares + 0.5)),
	}
	return resVO, irisErr
}

func (s ShareService) GetExRate(reqVO vo.ExRateReqVO) (
	vo.ExRateResVO, errors.IrisError)  {

	var (
		resVO vo.ExRateResVO
	)

	address, err := bech32.ConvertHexToBech32(reqVO.ValidatorAddress)
	if err != nil {
		return resVO, ConvertBadRequestErr(err)
	}

	uri := fmt.Sprintf(constants.HttpUriGetExRate, address)

	statusCode, resBytes := HttpClientGetData(uri)

	// statusCode != 200
	if !helper.SliceContains(constants.SuccessStatusCodes, statusCode) {
		return resVO, ConvertSysErr(fmt.Errorf(string(resBytes)))
	}


	if err := json.Unmarshal(resBytes, &resVO); err != nil {
		return resVO, ConvertSysErr(err)
	}

	return resVO, irisErr
}
