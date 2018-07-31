package irishub

import (
	"testing"

	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
)

func TestValidatorRouter_GetList(t *testing.T) {
	type args struct {
		c irisProtoc.IRISHubServiceClient
	}
	tests := []struct {
		name string
		h    ValidatorRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := ValidatorRouter{}
			h.GetList(tt.args.c)
		})
	}
}

func TestValidatorRouter_GetValidatorExRate(t *testing.T) {
	type args struct {
		c irisProtoc.IRISHubServiceClient
	}
	tests := []struct {
		name string
		h    ValidatorRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := ValidatorRouter{}
			h.GetValidatorExRate(tt.args.c)
		})
	}
}
