package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type BalanceHandler struct {
}

func (c BalanceHandler) Handler(ctx context.Context, req *commonProtoc.BalanceRequest) (
	*commonProtoc.BalanceResponse, error) {

	reqVO := c.buildRequest(req)
	resVO, err := accountService.GetBalance(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return c.buildResponse(resVO), nil
}

func (c BalanceHandler) buildRequest(req *commonProtoc.BalanceRequest) vo.BalanceReqVO {

	reqVO := vo.BalanceReqVO{
		Address: req.GetAddress(),
	}

	return reqVO
}

func (c BalanceHandler) buildResponse(resVO vo.BalanceResVO) *commonProtoc.BalanceResponse {

	coins := resVO.Data.Coins
	var modelCoins []*commonProtoc.Coin

	if len(coins) > 0 {
		for _, v := range coins {
			modelCoin := commonProtoc.Coin{
				Denom:  v.Denom,
				Amount: float64(v.Amount),
			}
			modelCoins = append(modelCoins, &modelCoin)
		}
	}

	response := commonProtoc.BalanceResponse{
		Coins: modelCoins,
	}

	return &response
}
