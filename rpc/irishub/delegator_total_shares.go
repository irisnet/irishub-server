package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type DelegatorTotalSharesHandler struct {
}

func (c DelegatorTotalSharesHandler) Handler(ctx context.Context, req *irisProtoc.TotalShareRequest) (
	*irisProtoc.TotalShareResponse, error) {

	reqVO := c.buildRequest(req)

	resVO, err := delegatorService.GetDelegatorTotalShare(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}
	return c.buildResponse(resVO), nil
}

func (c DelegatorTotalSharesHandler) buildRequest(req *irisProtoc.TotalShareRequest) vo.DelegatorTotalShareReqVO {

	reqVO := vo.DelegatorTotalShareReqVO{
		Address: req.GetAddress(),
	}

	return reqVO
}

func (c DelegatorTotalSharesHandler) buildResponse(resVO vo.DelegatorTotalShareResVO) *irisProtoc.TotalShareResponse {

	response := irisProtoc.TotalShareResponse{
		TotalShares:     resVO.TotalShares,
		BondedTokens:    resVO.ToTalBondedTokens,
		UnbondingTokens: resVO.ToTalUnbondingTokens,
	}

	return &response
}
