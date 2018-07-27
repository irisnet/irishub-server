package blockchain

import (
	"testing"

	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
)

func TestAccountRouter_GetSequence(t *testing.T) {
	type args struct {
		c commonProtoc.BlockChainServiceClient
	}
	tests := []struct {
		name string
		h    AccountRouter
		args args
	}{
		{
			name: "test get sequence",
			h: AccountRouter{},
			args: args{
				c: *Client,
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountRouter{}
			h.GetSequence(tt.args.c)
		})
	}
}

func TestAccountRouter_GetBalance(t *testing.T) {
	type args struct {
		c commonProtoc.BlockChainServiceClient
	}
	tests := []struct {
		name string
		h    AccountRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountRouter{}
			h.GetBalance(tt.args.c)
		})
	}
}
