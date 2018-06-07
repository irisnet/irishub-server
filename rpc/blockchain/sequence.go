package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/iris-api-server/rpc"
	"github.com/irisnet/iris-api-server/rpc/vo"
	"golang.org/x/net/context"
)

type SequenceController struct {

}

func (c SequenceController) Handler(ctx context.Context, request *chainModel.SequenceRequest) (
	*chainModel.SequenceResponse, error) {
	
	reqVO := c.buildRequest(request)
	resVO, err := sequenceService.GetSequence(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.buildResponse(resVO), nil
}

func (c SequenceController) buildRequest(req *chainModel.SequenceRequest) vo.SequenceReqVO {
	reqVO := vo.SequenceReqVO{
		Address: req.GetAddress(),
	}
	
	return reqVO
}

func (c SequenceController) buildResponse(resVO vo.SequenceResVO) *chainModel.SequenceResponse {
	response := chainModel.SequenceResponse{
		Sequence: resVO.Sequence,
		Height: resVO.Height,
	}
	
	return &response
}