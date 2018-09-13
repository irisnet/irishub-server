package blockchain

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
	"golang.org/x/net/context"
)

type SequenceHandler struct {
}

func (c SequenceHandler) Handler(ctx context.Context, request *commonProtoc.SequenceRequest) (
	*commonProtoc.SequenceResponse, error) {

	reqVO := c.buildRequest(request)
	resVO, err := accountService.GetSequence(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return c.buildResponse(resVO), nil
}

func (c SequenceHandler) buildRequest(req *commonProtoc.SequenceRequest) vo.SequenceReqVO {
	reqVO := vo.SequenceReqVO{
		Address: req.GetAddress(),
	}

	return reqVO
}

func (c SequenceHandler) buildResponse(resVO vo.SequenceResVO) *commonProtoc.SequenceResponse {
	response := commonProtoc.SequenceResponse{
		Sequence: resVO.Sequence,
		Ext:      resVO.Ext,
	}

	return &response
}
