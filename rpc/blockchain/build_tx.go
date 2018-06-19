package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
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
		Sequence: request.Sequence,
		From: vo.Address{
			Chain: request.Sender.GetChain(),
			App: request.Sender.GetApp(),
			Addr: request.Sender.GetAddr(),
		},
		To: vo.Address{
			Chain: request.Receiver.GetChain(),
			App: request.Receiver.GetApp(),
			Addr: request.Receiver.GetAddr(),
		},
		Amount: coins,
		Memo:vo.Memo{
			Id: request.Memo.ID,
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
