package blockchain

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
	"golang.org/x/net/context"
)

type SimulateTxHandler struct {
}

func (c SimulateTxHandler) Handler(ctx context.Context, req *commonProtoc.SimulateTxRequest) (
	*commonProtoc.SimulateTxResponse, error) {

	reqVO := c.buildRequest(req)
	res, err := simulateTxService.SimulateTx(reqVO)
	if err.IsNotNull() {
		return nil, BuildException(err)
	}
	return c.buildResponse(res), nil
}

func (c SimulateTxHandler) buildRequest(request *commonProtoc.SimulateTxRequest) vo.SimulateTxReqVO {
	reqVO := vo.SimulateTxReqVO{
		Tx: request.GetTx(),
	}

	return reqVO
}

func (c SimulateTxHandler) buildResponse(resVO vo.SimulateTxResVO) *commonProtoc.SimulateTxResponse {
	var details = make([]*commonProtoc.RewardDetail, len(resVO.Records))
	for i, record := range resVO.Records {
		details[i] = &commonProtoc.RewardDetail{
			ValAddress: record.ValAddress,
			Name:       record.Name,
			Amount: &commonProtoc.Coin{
				Amount: record.Amount.Amount,
				Denom:  record.Amount.Denom,
			},
		}
	}

	return &commonProtoc.SimulateTxResponse{
		Details: details,
		Gas:     resVO.Gas,
	}
}
