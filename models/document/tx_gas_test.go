package document

import (
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	"testing"
)

func TestTxGas_GetTxGas(t *testing.T) {
	type args struct {
		txType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get tx gas",
			args: args{
				txType: constants.TxTypeStakeDelegate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := TxGas{}
			res, err := d.GetTxGas(tt.args.txType)
			if err != nil {
				t.Error(err)
			} else {
				logger.Info.Println(helper.ToJson(res))
			}
		})
	}
}
