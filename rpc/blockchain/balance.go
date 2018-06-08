package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type BalanceController struct {

}

func (c BalanceController) Handler(ctx context.Context, req *chainModel.BalanceRequest) (
	*chainModel.BalanceResponse, error) {
	
	reqVO := c.buildRequest(req)
	resVO, err := balanceService.GetBalance(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.buildResponse(resVO), nil
}

func (c BalanceController) buildRequest(req *chainModel.BalanceRequest) vo.BalanceReqVO  {
	
	reqVO := vo.BalanceReqVO{
		Address: req.GetAddress(),
	}
	
	return reqVO
}

func (c BalanceController) buildResponse(resVO vo.BalanceResVO) *chainModel.BalanceResponse {
	
	coins := resVO.Data.Coins
	var modelCoins []*chainModel.Coin
	
	if len(coins) > 0 {
		for _, v := range coins {
			modelCoin := chainModel.Coin{
				Denom: v.Denom,
				Amount: float64(v.Amount),
			}
			modelCoins = append(modelCoins, &modelCoin)
		}
	}
	
	response := chainModel.BalanceResponse{
		Coins: modelCoins,
	}
	
	return &response
}
