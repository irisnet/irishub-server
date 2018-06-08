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
		name  string
		s     TxService
		args  args
	}{
		{
			name: "Test get tx list",
			s: TxService{},
			args: args{
				reqVO: vo.TxListReqVO{
					Address: "BED890EB9DB1309E0884DF8BDD41B16461D8E194",
					Page: 1,
					PerPage: 20,
					Status: "",
					Type: "receive",
					StartTime: "",
					EndTime: "2018-06-07 00:00:00",
					Sort: "-time",
					Q: "",
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
