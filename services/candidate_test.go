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
		name  string
		s     CandidateService
		args  args
	}{
		{
			name: "test get candidate list",
			s: CandidateService{},
			args: args{
				listVo: vo.CandidateListReqVO{
					Address: "BED890EB9DB1309E0884DF8BDD41B16461D8E194",
					Page: 1,
					PerPage: 50,
					Sort: "-voting_power",
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


