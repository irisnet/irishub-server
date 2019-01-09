package services

import (
	"fmt"
	"strings"
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
)

func TestPostTxService_PostTx(t *testing.T) {
	type args struct {
		vo vo.PostTxReqVO
	}
	tests := []struct {
		name string
		s    PostTxService
		args args
	}{
		{
			name: "Test post tx",
			s:    PostTxService{},
			args: args{
				vo: vo.PostTxReqVO{},
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

func TestTrimStr(t *testing.T) {
	key := "withdraw-reward-from-validator-fva1pmazgz5652uy9v3esana6rq54e7tuup7fv6evk"
	//value := "1466.7412168003iris-atto"

	if strings.HasPrefix(key, "withdraw-reward-from-validator-") {
		valAddr := strings.TrimPrefix(key, "withdraw-reward-from-validator-")
		fmt.Println(valAddr)
	}
}
