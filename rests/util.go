package rests

import (
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/utils/constants"
)

var HttpStatusOk = constants.STATUS_CODE_OK

type BaseResponse struct {
	Status  string      `json:"status_code"`
	ErrCode string      `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func BuildResponse(data interface{}) *BaseResponse {
	return &BaseResponse{
		Status: constants.STATUS_SUCCESS,
		Data:   data,
	}
}

func BuildExpResponse(error errors.IrisError) *BaseResponse {
	return &BaseResponse{
		Status: constants.STATUS_FAIL,
		ErrCode: error.ErrCode,
		ErrMsg: error.ErrMsg,
		Data: struct {
		}{},
	}
}
