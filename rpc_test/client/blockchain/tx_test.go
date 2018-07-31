package blockchain

import (
	"testing"

	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
)

func TestTxRouter_GetTxList(t *testing.T) {
	type args struct {
		c *commonProtoc.BlockChainServiceClient
	}
	tests := []struct {
		name string
		h    TxRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := TxRouter{}
			h.GetTxList(tt.args.c)
		})
	}
}

func TestTxRouter_GetTxDetail(t *testing.T) {
	type args struct {
		c   *commonProtoc.BlockChainServiceClient
		req *commonProtoc.TxDetailRequest
	}
	tests := []struct {
		name string
		h    TxRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := TxRouter{}
			h.GetTxDetail(tt.args.c, tt.args.req)
		})
	}
}
