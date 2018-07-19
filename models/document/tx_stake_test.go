package document

import (
	"testing"
	"time"
	
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestStakeTx_GetStakeTxList(t *testing.T) {
	var (
		startTime time.Time
		endTime time.Time
		err error
	)
	
	startTime, err = helper.ParseTime(constants.TIME_LAYOUT_FULL, "2018-05-08 00:00:00")
	if err != nil {
		logger.Error.Fatalln(err)
	}
	endTime, err = helper.ParseTime(constants.TIME_LAYOUT_FULL, "2018-05-09 00:00:00")
	if err != nil {
		logger.Error.Fatalln(err)
	}
	
	type args struct {
		address   string
		pubKey    string
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
			name: "test get delegator list",
			args:args{
				address: "D4C9FEA4BEBD5600878EC90E1F87B5F07A9DB00A",
				startTime: startTime,
				endTime: endTime,
				skip: 0,
				limit: 10,
				sorts:[]string{"time"},
			},
		},
		{
			name: "test get delegator tx list on a pubKey",
			args: args{
				address: "D4C9FEA4BEBD5600878EC90E1F87B5F07A9DB00A",
				pubKey: "01EFE2106DAC707FF21B47C03BABC4CB1EF10F28289B142B0E5017CC5B71721A",
				txType: "delegate",
				startTime: time.Time{},
				endTime: time.Time{},
				skip: 0,
				limit: 10,
				sorts: []string{"-time"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := StakeTx{
			}
			got, err := m.GetStakeTxList(
				tt.args.address, tt.args.pubKey, tt.args.txType,
				tt.args.startTime, tt.args.endTime, tt.args.skip, tt.args.limit, tt.args.sorts)
			if err != nil {
				logger.Error.Fatalln(err.Error())
			}
			logger.Info.Println(helper.ToJson(got))
		})
	}
}
