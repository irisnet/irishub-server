package rests

import (
	"testing"
	
	"github.com/irisnet/iris-api-server/errors"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/utils/helper"
)

func TestBuildResponse(t *testing.T) {
	type response struct {
		Name string `json:"name"`
		Password string `json:"password"`
	}

	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *BaseResponse
	}{
		{
			name: "test build response",
			args: struct{ data interface{} }{
				data: response{
					Name: "hello",
					Password: "world",
				},
				},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := BuildResponse(tt.args.data)
			logger.Info.Println(helper.ToJson(response))
		})
	}
}

func TestBuildErrResponse(t *testing.T) {
	type args struct {
		error errors.IrisError
	}
	tests := []struct {
		name string
		args args
		want *BaseResponse
	}{
		{
			name: "test build exception response",
			args: struct{ error errors.IrisError }{
				error: struct {
					ErrCode uint32
					ErrMsg  string
				}{
					ErrCode: 40001,
					ErrMsg: "参数缺失",
						}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := BuildExpResponse(tt.args.error)

			logger.Info.Println(helper.ToJson(response))
		})
	}
}
