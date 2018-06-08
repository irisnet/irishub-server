package services

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rests/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestStakeTxService_GetList(t *testing.T) {
	type args struct {
		vo vo.StakeTxListVO
	}
	tests := []struct {
		name  string
		s     StakeTxService
		args  args
	}{
		{
			name: "Test get list of stake tx",
			args: args{
				vo: vo.StakeTxListVO{
					BaseVO: vo.BaseVO{
						Page: 1,
						PerPage: 10,
					},
					Address: "D4C9FEA4BEBD5600878EC90E1F87B5F07A9DB00A",
					PubKey: "01EFE2106DAC707FF21B47C03BABC4CB1EF10F28289B142B0E5017CC5B71721A",
					TxType: "delegate",
					StartTime: "2018-05-09 00:00:00",
					EndTime: "",
					Sort: "-time",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StakeTxService{}
			got, err := s.GetList(tt.args.vo)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(got))
			
		})
	}
}
