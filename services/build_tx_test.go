package services

import (
	"testing"
	
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rpc/vo"
)

func TestBuildTxService_BuildTx(t *testing.T) {
	type args struct {
		vo vo.BuildTxVO
	}
	tests := []struct {
		name    string
		s       BuildTxService
		args    args
	}{
		{
			name: "Test build tx",
			s: BuildTxService{
			},
			args:args{
				vo: vo.BuildTxVO{
					Fees: vo.Fee{
						Denom: "iris",
						Amount: 0,
					},
					Multi: false,
					Sequence: 1,
					From: vo.Address{
						Addr: "1719B45561AE16339CDCDC8D06AF9322B598D3FB",
					},
					To: vo.Address{
						Addr: "BED890EB9DB1309E0884DF8BDD41B16461D8E194",
					},
					Amount: []vo.Coin {
						vo.Coin{
							Denom: "iris",
							Amount: 15,
						},
					},
					Memo:vo.Memo{
					
					},
				},
			},
			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BuildTxService{}
			got, err := s.BuildTx(tt.args.vo)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(got)
		})
	}
}
