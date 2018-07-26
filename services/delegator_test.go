package services

import (
	"testing"

	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestDelegatorService_DelegatorCandidateList(t *testing.T) {
	type args struct {
		reqVO vo.DelegatorCandidateListReqVO
	}
	tests := []struct {
		name  string
		args  args
	}{
		{
			name: "test get delegator candidate list",
			args: args{
				reqVO: vo.DelegatorCandidateListReqVO{
					Address: "461B0D58301072D68EB95C54DDFCFBFF7D67DA7C",

					Page: 1,
					PerPage: 10,
					Sort: "-time",
					Q: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DelegatorService{}
			res, err := s.DelegatorCandidateList(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}

func TestDelegatorService_GetDelegatorTotalShare(t *testing.T) {
	type args struct {
		reqVO vo.TotalShareReqVO
	}
	tests := []struct {
		name  string
		args  args
	}{
		{
			name: "test get delegator total share",
			args: args{
				reqVO: vo.TotalShareReqVO{
					Address: "461B0D58301072D68EB95C54DDFCFBFF7D67DA7C",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DelegatorService{}
			res, err := s.GetDelegatorTotalShare(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))

		})
	}
}
