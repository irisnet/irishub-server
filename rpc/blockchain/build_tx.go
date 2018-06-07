package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	vo "github.com/irisnet/iris-api-server/rpc/vo"
)

// transform common request to suitable request
//
// buildTxRequest is common model,
// every api server of chain may need transform them before handle these data
func TransformBuildTxRequest(request *chainModel.BuildTxRequest) (vo.BuildTxVO) {
	var coins []vo.Coin
	for _, amount := range request.Amount {
		coin := vo.Coin{
			Denom: amount.GetDenom(),
			Amount: int64(amount.GetAmount()),
		}
		coins = append(coins, coin)
	}
	
	
	buildTxVO := vo.BuildTxVO{
		Fees: vo.Fee{
			Denom: request.Fee.Denom,
			Amount: int64(request.Fee.Amount),
		},
		Multi: false,
		Sequence: request.Nonce,
		From: vo.Address{
			Chain: request.From.GetChain(),
			App: request.From.GetApp(),
			Addr: request.From.GetAddr(),
		},
		To: vo.Address{
			Chain: request.To.GetChain(),
			App: request.To.GetApp(),
			Addr: request.To.GetAddr(),
		},
		Amount: coins,
		Memo:vo.Memo{
			Id: request.Memo.Id,
			Text: request.Memo.GetText(),
		},
	}
	
	return buildTxVO
}

// transform service result to common response
func TransformBuildTxResponse(res []byte) (*chainModel.BuildTxResponse) {
	return &chainModel.BuildTxResponse{
		Data: res,
	}
}
