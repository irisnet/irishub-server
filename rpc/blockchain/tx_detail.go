package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"golang.org/x/net/context"
)

type TxDetailHandler struct {

}

func (c TxDetailHandler) Handler(ctx context.Context, req *commonProtoc.TxDetailRequest) (
	*commonProtoc.TxDetailResponse, error) {
	
	reqVO := c.BuildRequest(req)
	
	resVO, err := txService.GetTxDetail(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.BuildResponse(resVO), nil
}

func (c TxDetailHandler) BuildRequest(req *commonProtoc.TxDetailRequest) vo.TxDetailReqVO {
	
	reqVO := vo.TxDetailReqVO{
		TxHash: req.GetTxHash(),
	}
	
	return reqVO
}

func (c TxDetailHandler) BuildResponse(resVO vo.TxDetailResVO) *commonProtoc.TxDetailResponse {
	resTx := resVO.Tx
	from := rpc.BuildResponseAddress(resTx.From)
	to := rpc.BuildResponseAddress(resTx.To)
	coins := rpc.BuildResponseCoins(resTx.Amount)
	fee := commonProtoc.FeeUsed{
	}
	
	response := commonProtoc.TxDetailResponse{
		TxHash: resTx.TxHash,
		Time: resTx.Time.String(),
		Height: resTx.Height,
		Sender: &from,
		Receiver: &to,
		Amount: coins,
		Type: resTx.Type,
		Status: constants.TxStatusSuccess,
		Fee: &fee,
		Memo: &commonProtoc.Memo{},
		Ext: []byte(resTx.Candidate.Description.Moniker),
	}
	
	return &response
}
