package services

import (
	"testing"

	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestAccountService_GetAccountNum(t *testing.T) {
	type args struct {
		reqVO vo.AccountNumReqVO
	}
	tests := []struct {
		name  string
		s     AccountService
		args  args
	}{
		{
			name: "test get account num",
			s: AccountService{},
			args: args{
				reqVO: vo.AccountNumReqVO{
					Address: "cosmosaccaddr16acdgh02w4yqwmu2yluujaymyqynfud56ryvn1",
				},
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AccountService{}
			res, err := s.GetAccountNum(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}

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
					Address: "cosmosaccaddr16acdgh02w4yqwmu2yluujaymyqynfud56ryvn0",
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
					Address: "cosmosaccaddr16acdgh02w4yqwmu2yluujaymyqynfud56ryvn0",
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
		})
	}
}
