package rests

import (
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/utils/constants"
)

var HTTP_OK = constants.STATUS_CODE_OK

type BaseResponse struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Data          interface{} `json:"data"`
}

type ExceptionResponse struct {
	BaseResponse
	Data struct {
		err_code    string
		err_message string
	} `json:"data"`
}

func BuildResponse(data interface{}) *BaseResponse {
	return &BaseResponse{
		StatusCode:    constants.STATUS_CODE_OK,
		StatusMessage: constants.STATUS_CODE_OK_MESSAGE,
		Data:          data,
	}
}

func BuildExceptionResponse(irisErr errors.IrisError) *ExceptionResponse {
	return &ExceptionResponse{
		BaseResponse: BaseResponse{
			StatusCode:    constants.STATUS_CODE_FAILED,
			StatusMessage: constants.STATUS_CODE_FAILED_MESSAGE,
		},
		Data: struct {
			err_code    string
			err_message string
		}{err_code: irisErr.ErrCode, err_message: irisErr.ErrMsg},
	}
}
