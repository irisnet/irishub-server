package services

import (
	"testing"

	"github.com/irisnet/iris-api-server/models/document"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
	"github.com/irisnet/iris-api-server/utils/helper"
)

func TestCandidateService_List(t *testing.T) {
	type args struct {
		listVo vo.CandidateListVo
	}
	baseVo := vo.BaseVO{
		Page:    1,
		PerPage: 10,
	}
	tests := []struct {
		name  string
		s     CandidateService
		args  args
		want  []document.Candidate
		want1 errors.IrisError
	}{
		{
			name: "test candidate list",
			args: args{
				listVo: vo.CandidateListVo{
					BaseVO:  baseVo,
					Sort:    "-voting_power",
					Q:       "",
					Address: "8CD379DAC8B6B7DB578A8E86C2527AE046AFAC0B",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CandidateService{}
			candidates, err := s.List(tt.args.listVo)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(candidates))
		})
	}
}

func TestCandidateService_DelegatorCandidateList(t *testing.T) {
	type args struct {
		listVo vo.DelegatorCandidateListVo
	}
	tests := []struct {
		name string
		s    CandidateService
		args args
	}{
		{
			name: "test delegator candidate list",
			args: args{
				listVo: vo.DelegatorCandidateListVo{
					Address: "8CD379DAC8B6B7DB578A8E86C2527AE046AFAC0B",
					Sort:    "",
					Q:       "",
					BaseVO: vo.BaseVO{
						Page:    1,
						PerPage: 10,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CandidateService{}
			candidates, err := s.DelegatorCandidateList(tt.args.listVo)
			if err.IsNotNull() {
				logger.Error.Panic(err)
			}
			logger.Info.Println(helper.ToJson(candidates))

		})
	}
}

func TestCandidateService_Detail(t *testing.T) {
	type args struct {
		pubKey  string
		address string
	}
	tests := []struct {
		name  string
		s     CandidateService
		args  args
	}{
		{
			name: "get detail of candidate",
			args: args{
				pubKey:"CB103698AC3FB4A181B4C168A0F8B72793990D514D9AB5A7E60389088D3E1C8D",
				address:"8CD379DAC8B6B7DB578A8E86C2527AE046AFAC0B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CandidateService{}
			candidate, err := s.Detail(tt.args.pubKey, tt.args.address)
			if err.IsNotNull() {
				logger.Error.Panicln(err)
			}
			logger.Info.Println(helper.ToJson(candidate))
		})
	}
}
