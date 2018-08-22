package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/models/document"
)

func BuildResAddress(address string) commonProtoc.Address {
	return commonProtoc.Address{
		Chain: "",
		App:   "",
		Addr:  address,
	}
}

func BuildResCoins(coins document.Coins) []*commonProtoc.Coin {
	var (
		modelCoins []*commonProtoc.Coin
	)

	if len(coins) > 0 {
		for _, v := range coins {
			modelCoin := commonProtoc.Coin{
				Denom:  v.Denom,
				Amount: float64(v.Amount),
			}
			modelCoins = append(modelCoins, &modelCoin)
		}
	}

	return modelCoins
}

func BuildResFeeAndGasLimit(fee document.Fee) (*commonProtoc.Fee, float64) {
	var (
		resFee      commonProtoc.Fee
		resGasLimit float64
	)
	resGasLimit = float64(fee.Gas)

	if len(fee.Amount) > 0 {
		feeAmount := fee.Amount[0]
		resFee = commonProtoc.Fee{
			Amount: feeAmount.Amount,
			Denom:  feeAmount.Denom,
		}
	}

	return &resFee, resGasLimit
}

func BuildResActualFee(actualFee document.ActualFee) *commonProtoc.Fee {
	var (
		resActualFee commonProtoc.Fee
	)

	resActualFee = commonProtoc.Fee{
		Amount: actualFee.Amount,
		Denom:  actualFee.Denom,
	}

	return &resActualFee
}
