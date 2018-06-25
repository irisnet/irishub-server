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
	resVO, err := buildTxService.BuildTx(buildTxVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return c.buildResponse(resVO), nil
}

func (c BuildTxHandler) buildRequest(req *commonProtoc.BuildTxRequest) (vo.BuildTxReqVO) {
	var coins []vo.Coin
	for _, amount := range req.Amount {
		coin := vo.Coin{
			Denom: amount.GetDenom(),
			Amount: int64(amount.GetAmount()),
		}
		coins = append(coins, coin)
	}
	
	
	reqVO := vo.BuildTxReqVO{
		Fees: vo.Fee{
			Denom:  req.Fee.Denom,
			Amount: int64(req.Fee.Amount),
		},
		Multi:    false,
		Sequence: req.Sequence,
		From: vo.Address{
			Chain: req.Sender.GetChain(),
			App:   req.Sender.GetApp(),
			Addr:  req.Sender.GetAddr(),
		},
		To: vo.Address{
			Chain: req.Receiver.GetChain(),
			App:   req.Receiver.GetApp(),
			Addr:  req.Receiver.GetAddr(),
		},
		Amount: coins,
		TxType: req.TxType,
	}
	
	if req.Memo != nil {
		reqVO.Memo = vo.Memo{
			Id: req.Memo.ID,
			Text: req.Memo.Text,
		}
	}
	
	return reqVO
}

func (c BuildTxHandler) buildResponse(resVO vo.BuildTxResVO) (*commonProtoc.BuildTxResponse) {
	return &commonProtoc.BuildTxResponse{
		Data: resVO.Data,
		Ext: resVO.Ext,
	}
}
