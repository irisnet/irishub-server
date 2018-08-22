package document

import (
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
	"testing"
)

func TestValidatorUpTime_GetUpTime(t *testing.T) {
	type args struct {
		valAddress []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get validator up time",
			args: args{
				valAddress: []string{
					"44662D2F27C3700ACDD4072E017EA48157F36939",
					"D99216FC43FB55EBB1802FDFE50A5B9585B9BFB9",
					"DFCDDB6906A05F4A84D33159B92FECADFD4E86A5",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := ValidatorUpTime{}
			res, err := d.GetUpTime(tt.args.valAddress)
			if err != nil {
				t.Error(err)
			} else {
				logger.Info.Println(helper.ToJson(res))
			}
		})
	}
}
