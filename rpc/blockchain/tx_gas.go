package blockchain

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
	"golang.org/x/net/context"
	"math"
)

type TxGasHandler struct {
}

func (h TxGasHandler) Handler(ctx context.Context, req *commonProtoc.TxGasRequest) (
	*commonProtoc.TxGasResponse, error) {
	reqVO := h.buildReq(req)

	resVO, err := txService.GetTxGas(reqVO)
	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return h.buildRes(resVO), nil
}

func (h TxGasHandler) buildReq(req *commonProtoc.TxGasRequest) vo.TxGasReqVO {
	reqVO := vo.TxGasReqVO{
		TxType: req.GetTxType(),
	}

	return reqVO
}

func (h TxGasHandler) buildRes(resVO vo.TxGasResVO) *commonProtoc.TxGasResponse {
	var (
		response    commonProtoc.TxGasResponse
		resGasPrice commonProtoc.GasPrice
	)

	resGasPrice = commonProtoc.GasPrice{
		Denom:       resVO.GasPrice.Denom,
		MinGasPrice: math.Ceil(resVO.GasPrice.MinGasPrice),
		MaxGasPrice: math.Ceil(resVO.GasPrice.MaxGasPrice),
		AvgGasPrice: math.Ceil(resVO.GasPrice.AvgGasPrice),
	}
	resGasLimit := int64(resVO.Gas.MaxGasUsed * 1.5)
	response = commonProtoc.TxGasResponse{
		TxType:   resVO.TxType,
		GasLimit: float64(resGasLimit),
		GasPrice: &resGasPrice,
	}

	return &response
}
