package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	vo "github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type BuildTxHandler struct {
}


func (c BuildTxHandler) Handler(ctx context.Context, request *commonProtoc.BuildTxRequest) (
	*commonProtoc.BuildTxResponse, error) {
	
	buildTxVO := c.buildRequest(request)
	res, err := buildTxService.BuildTx(buildTxVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return c.buildResponse(res), nil
}

// transform common request to suitable request
//
// buildTxRequest is common model,
// every api server of chain may need transform them before handle these data
func (c BuildTxHandler) buildRequest(request *commonProtoc.BuildTxRequest) (vo.BuildTxReqVO) {
	var coins []vo.Coin
	for _, amount := range request.Amount {
		coin := vo.Coin{
			Denom: amount.GetDenom(),
			Amount: int64(amount.GetAmount()),
		}
		coins = append(coins, coin)
	}
	
	
	buildTxVO := vo.BuildTxReqVO{
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
func (c BuildTxHandler) buildResponse(res []byte) (*commonProtoc.BuildTxResponse) {
	return &commonProtoc.BuildTxResponse{
		Data: res,
	}
}
