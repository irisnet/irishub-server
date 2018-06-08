package document

import (
	"testing"
	"time"
	
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestCommonTx_GetList(t *testing.T) {
	startTime, _ := helper.ParseFullTime("2018-05-08 00:00:00")
	endTime, _ := helper.ParseFullTime("2018-09-08 00:00:00")
	
	type args struct {
		address   string
		txType    string
		startTime time.Time
		endTime   time.Time
		skip      int
		limit     int
		sorts     []string
	}
	tests := []struct {
		name    string
		args    args
	}{
		{
			name: "test get list",
			args: args{
				address: "D4C9FEA4BEBD5600878EC90E1F87B5F07A9DB00A",
				txType: "unBond",
				startTime: startTime,
				endTime: endTime,
				skip: 0,
				limit: 10,
				sorts: []string{"-time"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CommonTx{
			}
			got, err := d.GetList(tt.args.address, tt.args.txType, tt.args.startTime, tt.args.endTime, tt.args.skip, tt.args.limit, tt.args.sorts)
			if err != nil {
				logger.Error.Fatalln(err.Error())
			}
			logger.Info.Println(helper.ToJson(got))
		})
	}
}