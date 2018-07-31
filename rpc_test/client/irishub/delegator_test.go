package irishub

import (
	"testing"

	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
)

func TestDelegatorRouter_GetDelegatorValidatorsList(t *testing.T) {
	type args struct {
		c irisProtoc.IRISHubServiceClient
	}
	tests := []struct {
		name string
		h    DelegatorRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := DelegatorRouter{}
			h.GetDelegatorValidatorsList(tt.args.c)
		})
	}
}

func TestDelegatorRouter_GetDelegatorTotalShares(t *testing.T) {
	type args struct {
		c irisProtoc.IRISHubServiceClient
	}
	tests := []struct {
		name string
		h    DelegatorRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := DelegatorRouter{}
			h.GetDelegatorTotalShares(tt.args.c)
		})
	}
}
