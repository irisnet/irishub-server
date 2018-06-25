package document

import (
"testing"


"github.com/irisnet/irishub-server/modules/logger"
"github.com/irisnet/irishub-server/utils/helper"

)

func TestDelegator_GetDelegatorList(t *testing.T) {
	type args struct {
		address string
		pubKey  []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test delegator list",
			args: args{
				address: "8CD379DAC8B6B7DB578A8E86C2527AE046AFAC0B",
				pubKey: []string{
					"CB103698AC3FB4A181B4C168A0F8B72793990D514D9AB5A7E60389088D3E1C8D",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delegator{}
			delegator, err := d.GetDelegatorListByAddressAndPubKeys(tt.args.address, tt.args.pubKey)
			if err != nil {
				logger.Error.Fatalln(err)
			}

			logger.Info.Println(helper.ToJson(delegator))
		})
	}
}

func TestDelegator_GetTotalTokenByAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
	}{
		{
			name: "test get total token by address",
			args: args{
				address: "E7113444CC960FD64D289CD3BD8A7FC123607ECE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delegator{
			}
			got, err := d.GetTotalSharesByAddress(tt.args.address)
			if err != nil {
				logger.Error.Fatalln(err.Error())
			}
			logger.Info.Println(helper.ToJson(got))
		})
	}
}
