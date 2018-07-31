package services

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestTxService_GetTxList(t *testing.T) {
	type args struct {
		reqVO vo.TxListReqVO
	}
	tests := []struct {
		name string
		s    TxService
		args args
	}{
		{
			name: "Test get tx list",
			s:    TxService{},
			args: args{
				reqVO: vo.TxListReqVO{
					Address:   "9A3FCB5662DAFDE94D4DE2E3114A7462A1B55F5C",
					Page:      1,
					PerPage:   20,
					Status:    "",
					Type:      "unbond",
					StartTime: "",
					EndTime:   "",
					Sort:      "-time",
					Q:         "",
					Ext: []byte("39F8BDAAFA6DF322313B60333317E9D4F18E70A6"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TxService{}
			res, err := s.GetTxList(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))

		})
	}
}

func TestTxService_getTxDetail(t *testing.T) {
	type args struct {
		reqVO vo.TxDetailReqVO
	}
	tests := []struct {
		name  string
		s     TxService
		args  args
	}{
		{
			name: "test get tx detail",
			s: TxService{},
			args: args{
				reqVO: vo.TxDetailReqVO{
					TxHash: "5289539FE1FE03E5B427F9ACADB0FC185B66EB54",
				},
			},
			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TxService{}
			res, err := s.GetTxDetail(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}
