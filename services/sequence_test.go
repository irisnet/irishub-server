package services

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestSequenceService_GetSequence(t *testing.T) {
	type args struct {
		reqVO vo.SequenceReqVO
	}
	tests := []struct {
		name  string
		c     SequenceService
		args  args
	}{
		{
			name: "Test get sequence",
			c: SequenceService{},
			args: args{
				reqVO: vo.SequenceReqVO{
					Address: "1719B45561AE16339CDCDC8D06AF9322B598D3FB",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := SequenceService{}
			res, err := c.GetSequence(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}
