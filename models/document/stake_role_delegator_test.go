package document

import (
	"testing"

	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/utils/helper"
)

func TestDelegator_GetDelegatorList(t *testing.T) {
	type args struct {
		address string
		pubKey  []string
	}
	tests := []struct {
		name    string
		args    args
	}{
		{
			name: "test delegator list",
			args: args{
				address:"8CD379DAC8B6B7DB578A8E86C2527AE046AFAC0B",
				pubKey:[]string{
					"CB103698AC3FB4A181B4C168A0F8B72793990D514D9AB5A7E60389088D3E1C8D",
				},
				},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delegator{
			}
			delegator, err := d.GetDelegatorListByAddressAndPubKeys(tt.args.address, tt.args.pubKey)
			if err != nil {
				logger.Error.Fatalln(err)
			}

			logger.Info.Println(helper.ToJson(delegator))
		})
	}
}
