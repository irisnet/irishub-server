package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"golang.org/x/net/context"
)

type TxListHandler struct {

}

func (c TxListHandler) Handler(ctx context.Context, req *commonProtoc.TxListRequest) (
	*commonProtoc.TxListResponse, error) {
	
	reqVO := c.buildRequest(req)
	resVO, err := txService.GetTxList(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.buildResponse(resVO), nil
}

func (c TxListHandler) buildRequest(req *commonProtoc.TxListRequest) vo.TxListReqVO {
	
	reqVO := vo.TxListReqVO{
		Address: req.GetAddress(),
		Page: req.GetPage(),
		PerPage: req.GetPerPage(),
		Status: req.Status,
		Type: req.Type,
		StartTime: req.StartTime,
		EndTime: req.EndTime,
		Sort: req.Sort,
		Q: req.Q,
	}
	
	return reqVO
}

func (c TxListHandler) buildResponse(resVO vo.TxListResVO) *commonProtoc.TxListResponse {
	response := commonProtoc.TxListResponse{}
	var resTxs []*commonProtoc.TxListObject
	
	if len(resVO.Txs) > 0 {
		for _, v := range resVO.Txs {
			from := rpc.BuildResponseAddress(v.From)
			to := rpc.BuildResponseAddress(v.To)
			
			
			var modelCoins []*commonProtoc.Coin
			modelCoins = rpc.BuildResponseCoins(v.Amount)
			
			resTxListObj := commonProtoc.TxListObject{
				TxHash: v.TxHash,
				Time: v.Time.String(),
				Height: v.Height,
				Sender: &from,
				Receiver: &to,
				Amount: modelCoins,
				Type: v.Type,
				Status: constants.TxStatusSuccess,
				Ext: "",
			}
			resTxs = append(resTxs, &resTxListObj)
		}
	}
	
	return &response
}
