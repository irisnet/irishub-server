package services

import (
	"testing"

	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rpc/vo"
)

func TestPostTxService_PostTx(t *testing.T) {
	type args struct {
		vo vo.PostTxVO
	}
	tests := []struct {
		name  string
		s     PostTxService
		args  args
	}{
		{
			name: "Test post tx",
			s: PostTxService{},
			args: args{
				vo: vo.PostTxVO{
				
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostTxService{}
			res, err := s.PostTx(tt.args.vo)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(res)
			
		})
	}
}
