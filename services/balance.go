package services

import (
	"encoding/json"
	"fmt"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
)

type BalanceService struct {

}

func (s BalanceService) GetBalance(reqVO vo.BalanceReqVO) (vo.BalanceResVO, errors.IrisError) {
	address := reqVO.Address
	var (
		resVO vo.BalanceResVO
	)
	
	uri := fmt.Sprintf(constants.HttpUriGetBalance, address)
	statusCode, resBytes := HttpClientGetData(uri)
	
	if statusCode == constants.StatusCodeBadRequest {
		return resVO, irisErr
	}
	
	
	if err := json.Unmarshal(resBytes, &resVO); err != nil {
		return resVO, NewIrisErr(errors.EC50001, errors.EM50001 + err.Error())
	}
	
	return resVO, irisErr
}
