package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type BuildTxHandler struct {
}

// Deprecated: no longer used
func (c BuildTxHandler) Handler(ctx context.Context, request *commonProtoc.BuildTxRequest) (
	*commonProtoc.BuildTxResponse, error) {

	buildTxVO := c.buildRequest(request)
	resVO, err := buildTxService.BuildTx(buildTxVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}
	return c.buildResponse(resVO), nil
}

func (c BuildTxHandler) buildRequest(req *commonProtoc.BuildTxRequest) vo.BuildTxReqVO {
	reqTx := req.GetTx()
	var coins []vo.Coin
	for _, amount := range reqTx.Amount {
		coin := vo.Coin{
			Denom:  amount.GetDenom(),
			Amount: amount.GetAmount(),
		}
		coins = append(coins, coin)
	}

	reqVO := vo.BuildTxReqVO{
		Fees: vo.Fee{
			Denom:  reqTx.Fee.Denom,
			Amount: reqTx.Fee.Amount,
		},
		Multi:    false,
		Sequence: reqTx.Sequence,
		From: vo.Address{
			Chain: reqTx.Sender.GetChain(),
			App:   reqTx.Sender.GetApp(),
			Addr:  reqTx.Sender.GetAddr(),
		},
		To: vo.Address{
			Chain: reqTx.Receiver.GetChain(),
			App:   reqTx.Receiver.GetApp(),
			Addr:  reqTx.Receiver.GetAddr(),
		},
		Amount: coins,
		TxType: reqTx.Type,
	}

	if reqTx.Memo != nil {
		reqVO.Memo = vo.Memo{
			Id:   reqTx.Memo.ID,
			Text: reqTx.Memo.Text,
		}
	}

	return reqVO
}

func (c BuildTxHandler) buildResponse(resVO vo.BuildTxResVO) *commonProtoc.BuildTxResponse {
	return &commonProtoc.BuildTxResponse{
		Data: resVO.Data,
		Ext:  resVO.Ext,
	}
}
