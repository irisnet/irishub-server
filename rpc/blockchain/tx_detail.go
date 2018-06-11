package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"golang.org/x/net/context"
)

type TxDetailController struct {

}

func (c TxDetailController) Handler(ctx context.Context, req *chainModel.TxDetailRequest) (
	*chainModel.TxDetailResponse, error) {
	
	return nil, nil
}

func (c TxDetailController) BuildRequest(req *chainModel.TxDetailRequest) vo.TxDetailReqVO {
	
	reqVO := vo.TxDetailReqVO{
		TxHash: req.GetTxHash(),
	}
	
	return reqVO
}

func (c TxDetailController) BuildResponse(resVO vo.TxDetailResVO) *chainModel.TxDetailResponse {
	resTx := resVO.Tx
	from := rpc.BuildResponseAddress(resTx.From)
	to := rpc.BuildResponseAddress(resTx.To)
	coins := rpc.BuildResponseCoins(resTx.Amount)
	fee := chainModel.TxDetailResponse_Fee{
	
	}
	
	response := chainModel.TxDetailResponse{
		TxHash: resTx.TxHash,
		Time: resTx.Time.String(),
		Height: uint64(resTx.Height),
		From: &from,
		To: &to,
		Amount: coins,
		Type: resTx.Type,
		Status: constants.TxStatusSuccess,
		Fee: &fee,
		Memo: &chainModel.Memo{},
		Ext: []byte(resTx.Candidate.Description.Moniker),
	}
	
	return &response
}
