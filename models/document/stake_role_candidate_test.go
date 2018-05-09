package document

import (
	"testing"

	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/utils/helper"
)


func TestCandidate_GetCandidatesList(t *testing.T) {

	type args struct {
		skip  int
		limit int
		sorts []string
	}
	tests := []struct {
		name   string
		args   args
	}{
		{
			name: "test get candidate list",
			args: args{
				skip: 0,
				limit: 10,
				sorts: []string{"-voting_power"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Candidate{
			}
			candidates, err := d.GetCandidatesList(tt.args.sorts, tt.args.skip, tt.args.limit)
			if err != nil {
				logger.Error.Panicln(err)
			}
			logger.Info.Println(helper.ToJson(candidates))
		})
	}
}
