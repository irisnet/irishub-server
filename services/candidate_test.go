package services

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestCandidateService_List(t *testing.T) {
	type args struct {
		listVo vo.CandidateListReqVO
	}
	tests := []struct {
		name string
		s    CandidateService
		args args
	}{
		{
			name: "test get candidate list",
			s:    CandidateService{},
			args: args{
				listVo: vo.CandidateListReqVO{
					Address: "BED890EB9DB1309E0884DF8BDD41B16461D8E194",
					Page:    1,
					PerPage: 50,
					Sort:    "-voting_power",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CandidateService{}
			res, err := s.List(tt.args.listVo)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}

func TestCandidateService_Detail(t *testing.T) {
	type args struct {
		reqVO vo.CandidateDetailReqVO
	}
	tests := []struct {
		name string
		s    CandidateService
		args args
	}{
		{
			name: "get candidate detail",
			s:    CandidateService{},
			args: args{
				reqVO: vo.CandidateDetailReqVO{
					Address: "BED890EB9DB1309E0884DF8BDD41B16461D8E194",
					PubKey:  "EFF0C056C8F1602DC6F61F87C6EE8ACCF1855BEFD8AFD8B4B2C90397312D768AB",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CandidateService{}
			res, err := s.Detail(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}

			logger.Info.Println(helper.ToJson(res))
		})
	}
}

func TestCandidateService_DelegatorCandidateList(t *testing.T) {
	type args struct {
		reqVO vo.DelegatorCandidateListReqVO
	}
	tests := []struct {
		name  string
		s     CandidateService
		args  args
	}{
		{
			name: "test get delegator candidate list",
			s: CandidateService{},
			args:args{
				reqVO: vo.DelegatorCandidateListReqVO{
					Address: "BED890EB9DB1309E0884DF8BDD41B16461D8E194",
					Page: 1,
					PerPage: 20,
					Sort: "-time",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CandidateService{}
			res, err := s.DelegatorCandidateList(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}
