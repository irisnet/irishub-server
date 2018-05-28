package services

import (
	"testing"

	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rests/vo"
	"github.com/irisnet/iris-api-server/utils/helper"
)

func TestCommonTxService_GetList(t *testing.T) {
	type args struct {
		vo vo.CommonTxListVO
	}
	tests := []struct {
		name  string
		s     CommonTxService
		args  args
	}{
		{
			name: "test get common tx list",
			args: args{
				vo: vo.CommonTxListVO{
					BaseVO: vo.BaseVO{
						Page: 1,
						PerPage: 50,
					},
					Address: "D4C9FEA4BEBD5600878EC90E1F87B5F07A9DB00A",
					TxType: "send",
					StartTime: "2018-05-27 00:00:00",
					EndTime: "",
					Sort: "-time",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CommonTxService{}
			got, err := s.GetList(tt.args.vo)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(got))
		})
	}
}
