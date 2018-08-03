package bech32

import (
	"testing"
	"github.com/irisnet/irishub-server/modules/logger"
)

func TestConvertHexToBech32(t *testing.T) {
	var (
		addr = "A0446775F9B6245EA8C19A006E8B29A7EC4BDE16"
		//addrBech32 = "cosmosaccaddr1xesgljj8yumjtf3s94tnttlkpe4xp2kzhf3rnr"
	)
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
	}{
		{
			name: "test convert hex to bech32",
			args: args{
				addr: addr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ConvertHexToBech32(tt.args.addr)
			if err != nil {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(res)
			//if res == addrBech32 {
			//	logger.Info.Println("Convert success")
			//} else {
			//	logger.Info.Println("Convert failed")
			//}
		})
	}
}
