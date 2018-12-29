package document

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestCandidate_GetCandidatesList(t *testing.T) {

	type args struct {
		q     string
		skip  int
		limit int
		sorts []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get candidate list",
			args: args{
				q:     "",
				skip:  0,
				limit: 10,
				sorts: []string{"-voting_power"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Candidate{}
			candidates, err := d.GetCandidatesList(tt.args.q, tt.args.sorts, tt.args.skip, tt.args.limit)
			if err != nil {
				logger.Error.Panicln(err)
			}
			logger.Info.Println(helper.ToJson(candidates))
		})
	}
}

func TestCandidate_GetTotalShares(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test get totalShares",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Candidate{}
			totalShares, err := d.GetTotalShares()
			if err != nil {
				logger.Error.Panicln(err)
			}
			logger.Info.Println(totalShares)
		})
	}
}

func TestCandidate_GetCandidateDetail(t *testing.T) {
	type args struct {
		pubKey string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get detail of candidate",
			args: args{
				pubKey: "CB103698AC3FB4A181B4C168A0F8B72793990D514D9AB5A7E60389088D3E1C8D",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Candidate{}
			candidate, err := d.GetCandidateDetail(tt.args.pubKey)
			if err != nil {
				logger.Error.Panicln(err)
			}
			logger.Info.Println(helper.ToJson(candidate))
		})
	}
}
