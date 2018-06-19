package rpc

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertIrisErrToGRPCErr(err errors.IrisError) error {
	return status.Error(codes.Code(err.ErrCode), err.ErrMsg)
}

func BuildResponseAddress(address string) chainModel.Address {
	return chainModel.Address{
		Chain: "",
		App: "",
		Addr: address,
	}
}

func BuildResponseCoins(coins document.Coins) []*chainModel.Coin {
	var (
		modelCoins []*chainModel.Coin
	)
	
	
	if len(coins) > 0 {
		for _, v := range coins {
			modelCoin := chainModel.Coin{
				Denom: v.Denom,
				Amount: float64(v.Amount),
			}
			modelCoins = append(modelCoins, &modelCoin)
		}
	}
	
	return modelCoins
}
