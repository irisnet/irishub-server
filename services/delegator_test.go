package services

import (
	"testing"

	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/utils/helper"
)

func TestDelegatorService_GetTotalShares(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name  string
		s     DelegatorService
		args  args
	}{
		{
			name: "test get total token",
			args: args{
				address: "3F9BB6D8A3938CD8E3A52584CC1F066027C522FE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DelegatorService{}
			got, err := s.GetTotalShares(tt.args.address)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(got))
		})
	}
}
