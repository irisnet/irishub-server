package services

import (
	"testing"

	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestAccountService_GetBalance(t *testing.T) {
	type args struct {
		reqVO vo.BalanceReqVO
	}
	tests := []struct {
		name  string
		s     AccountService
		args  args
	}{
		{
			name: "test get balance",
			s: AccountService{},
			args: args{
				reqVO: vo.BalanceReqVO{
					Address: "ADA14398D8FA297E29AB7BA241C7B955F2680C46",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AccountService{}
			res, err := s.GetBalance(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}

func TestAccountService_GetSequence(t *testing.T) {
	type args struct {
		reqVO vo.SequenceReqVO
	}
	tests := []struct {
		name  string
		s     AccountService
		args  args
	}{
		{
			name: "test get sequence",
			s: AccountService{},
			args: args{
				reqVO: vo.SequenceReqVO{
					Address: "ADA14398D8FA297E29AB7BA241C7B955F2680C46",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AccountService{}
			res, err := s.GetSequence(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
			logger.Info.Println(string(res.Ext))
		})
	}
}
