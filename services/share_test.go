package services

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestShareService_GetDelegatorTotalShare(t *testing.T) {
	type args struct {
		reqVO vo.ShareReqVO
	}
	tests := []struct {
		name  string
		s     ShareService
		args  args
	}{
		{
			name: "test get delegator total shares",
			s: ShareService{},
			args: args{
				reqVO: vo.ShareReqVO{
					Address: "461B0D58301072D68EB95C54DDFCFBFF7D67DA7C",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShareService{}
			res, err := s.GetDelegatorTotalShare(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
			
		})
	}
}
