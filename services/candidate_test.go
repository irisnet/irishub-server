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
		Page: 1,
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
					BaseVO: baseVo,
					Sort: "-voting_power",
					Q: "",
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
