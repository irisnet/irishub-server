package blockchain

import (
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"golang.org/x/net/context"
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
)

type TxListController struct {

}

func (c TxListController) Handler(ctx context.Context, req *chainModel.TxListRequest) (
	*chainModel.TxListResponse, error) {
	
	reqVO := c.buildRequest(req)
	resVO, err := txService.GetTxList(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.buildResponse(resVO), nil
}

func (c TxListController) buildRequest(req *chainModel.TxListRequest) vo.TxListReqVO {
	
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

func (c TxListController) buildResponse(resVO vo.TxListResVO) *chainModel.TxListResponse {
	response := chainModel.TxListResponse{}
	var resTxs []*chainModel.TxListObject
	
	if len(resVO.Txs) > 0 {
		for _, v := range resVO.Txs {
			from := rpc.BuildResponseAddress(v.From)
			to := rpc.BuildResponseAddress(v.To)
			
			
			var modelCoins []*chainModel.Coin
			modelCoins = rpc.BuildResponseCoins(v.Amount)
			
			resTxListObj := chainModel.TxListObject{
				TxHash: v.TxHash,
				Time: v.Time.String(),
				Height: uint64(v.Height),
				From: &from,
				To: &to,
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
