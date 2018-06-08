package services

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestBalanceService_GetBalance(t *testing.T) {
	type args struct {
		reqVO vo.BalanceReqVO
	}
	tests := []struct {
		name  string
		s     BalanceService
		args  args
	}{
		{
			name: "test get balance",
			s: BalanceService{},
			args: args{
				reqVO: vo.BalanceReqVO{
					Address: "9AB63F3E1633D361915D6217F85D8986FCF5F496",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BalanceService{}
			res, err := s.GetBalance(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
			
		})
	}
}
