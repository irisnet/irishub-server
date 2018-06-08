package rests

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/utils/constants"
)

var OK = constants.STATUS_CODE_OK

type BaseResponse struct {
	Status  string      `json:"status_code"`
	ErrCode uint32      `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func BuildResponse(data interface{}) BaseResponse {
	return BaseResponse{
		Status: constants.STATUS_SUCCESS,
		Data:   data,
	}
}

func BuildExpResponse(error errors.IrisError) BaseResponse {
	return BaseResponse{
		Status: constants.STATUS_FAIL,
		ErrCode: error.ErrCode,
		ErrMsg: error.ErrMsg,
		Data: struct {
		}{},
	}
}
