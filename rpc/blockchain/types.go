package blockchain

import (
	"github.com/irisnet/irishub-server/models/document"
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
)

func BuildAddressRes(address string) commonProtoc.Address {
	return commonProtoc.Address{
		Chain: "",
		App:   "",
		Addr:  address,
	}
}

func BuildCoinsRes(coins document.Coins) []*commonProtoc.Coin {
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

func AddCoin(coinA commonProtoc.Coin, coinB commonProtoc.Coin) (coin commonProtoc.Coin) {
	if coinA.Denom == "" {
		coinA.Denom = coinB.Denom
	}
	if coinB.Denom == "" {
		coinB.Denom = coinA.Denom
	}
	if coinA.Denom == coinB.Denom {
		coin.Amount = coinA.Amount + coinB.Amount
		coin.Denom = coinA.Denom
	}
	return coin
}

func BuildFeeAndGasLimitRes(fee document.Fee) (*commonProtoc.Fee, float64) {
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

func BuildActualFeeRes(actualFee document.ActualFee) *commonProtoc.Fee {
	var (
		resActualFee commonProtoc.Fee
	)

	resActualFee = commonProtoc.Fee{
		Amount: actualFee.Amount,
		Denom:  actualFee.Denom,
	}

	return &resActualFee
}

func BuildMemoRes(memo string) *commonProtoc.Memo {
	var (
		resMemo commonProtoc.Memo
	)
	resMemo = commonProtoc.Memo{
		ID:   0,
		Text: []byte(memo),
	}

	return &resMemo
}
