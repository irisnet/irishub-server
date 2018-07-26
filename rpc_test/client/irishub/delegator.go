package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

type DelegatorRouter struct {

}

func (h DelegatorRouter) buildListReq() *irisProtoc.DelegatorCandidateListRequest  {
	req := irisProtoc.DelegatorCandidateListRequest{
		Address: "",

		Page: DefaultPage,
		PerPage: DefaultPerPage,
		Sort: DefaultSorts,
		Q: "",
	}

	return &req
}

func (h DelegatorRouter) buildTotalSharesReq() *irisProtoc.TotalShareRequest  {
	req := irisProtoc.TotalShareRequest{
		Address: "",
	}

	return &req
}

func (h DelegatorRouter) GetDelegatorValidatorsList(c irisProtoc.IRISHubServiceClient)  {
	res, err := c.GetDelegatorCandidateList(DefaultCtx, h.buildListReq())
	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println(helper.ToJson(res))
}

func (h DelegatorRouter) GetDelegatorTotalShares(c irisProtoc.IRISHubServiceClient)  {
	res, err := c.GetDelegatorTotalShares(DefaultCtx, h.buildTotalSharesReq())
	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println(helper.ToJson(res))
}
