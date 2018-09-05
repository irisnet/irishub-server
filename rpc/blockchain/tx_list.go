package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type TxListHandler struct {
}

func (c TxListHandler) Handler(ctx context.Context, req *commonProtoc.TxListRequest) (
	[]*commonProtoc.Tx, error) {

	reqVO := c.buildRequest(req)
	resVO, err := txService.GetTxList(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return c.buildResponse(resVO), nil
}

func (c TxListHandler) buildRequest(req *commonProtoc.TxListRequest) vo.TxListReqVO {

	reqVO := vo.TxListReqVO{
		Address:   req.GetAddress(),
		Page:      req.GetPage(),
		PerPage:   req.GetPerPage(),
		Status:    req.Status,
		Type:      req.Type,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Sort:      req.Sort,
		Q:         req.Q,
		Ext:       req.GetExt(),
	}

	return reqVO
}

func (c TxListHandler) buildResponse(resVO vo.TxListResVO) []*commonProtoc.Tx {
	var resTxs []*commonProtoc.Tx

	if len(resVO.Txs) > 0 {
		for _, v := range resVO.Txs {
			var (
				modelCoins []*commonProtoc.Coin
			)
			from := BuildAddressRes(v.From)
			to := BuildAddressRes(v.To)

			modelCoins = BuildCoinsRes(v.Amount)

			resFee, resGasLimit := BuildFeeAndGasLimitRes(v.Fee)
			actualFee := BuildActualFeeRes(v.ActualFee)

			memo := BuildMemoRes(v.Memo)

			resTxListObj := commonProtoc.Tx{
				TxHash:    v.TxHash,
				Time:      v.Time.String(),
				Height:    v.Height,
				Sender:    &from,
				Receiver:  &to,
				Amount:    modelCoins,
				Type:      v.Type,
				Status:    v.Status,
				Ext:       []byte{},
				Fee:       resFee,
				GasLimit:  resGasLimit,
				GasUsed:   float64(v.GasUsed),
				ActualFee: actualFee,
				Memo:      memo,
			}
			resTxs = append(resTxs, &resTxListObj)
		}
	}

	return resTxs
}
