package services

import (
	"encoding/json"
	"fmt"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

type SequenceService struct {
}


func (c SequenceService) GetSequence(reqVO vo.SequenceReqVO) (vo.SequenceResVO, errors.IrisError) {
	address := reqVO.Address
	var (
		resVO vo.SequenceResVO
	)
	
	uri := fmt.Sprintf(constants.HttpUriGetSequence, address)
	statusCode, res := HttpClientGetData(uri)
	
	if helper.SliceContains(constants.ErrorStatusCodes, statusCode) {
		return resVO, NewIrisErr(errors.EC40001, errors.EM40001 + string(res), nil)
	}
	
	err := json.Unmarshal(res, &resVO)
	if err != nil {
		return resVO, NewIrisErr(errors.EC40002, errors.EM40002, err)
	}
	
	return resVO, irisErr
}
