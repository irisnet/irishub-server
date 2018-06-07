package services

import (
	"encoding/json"
	"fmt"
	
	"github.com/irisnet/iris-api-server/errors"
	"github.com/irisnet/iris-api-server/rpc/vo"
	"github.com/irisnet/iris-api-server/utils/constants"
	"github.com/irisnet/iris-api-server/utils/helper"
)

type SequenceService struct {
}

var (
	resVO vo.SequenceResVO
)

func (c SequenceService) GetSequence(reqVO vo.SequenceReqVO) (vo.SequenceResVO, errors.IrisError) {
	address := reqVO.Address
	uri := fmt.Sprintf(constants.HttpUriGetSequence, address)
	statusCode, res := HttpClientGetData(uri)
	
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return resVO, NewIrisErr(errors.EC40001, errors.EM40001 + string(res))
	}
	
	err := json.Unmarshal(res, &resVO)
	if err != nil {
		return resVO, NewIrisErr(errors.EC40002, errors.EM40002 + err.Error())
	}
	
	return resVO, irisErr
}
