package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

type ValidatorRouter struct {
	
}

func (h ValidatorRouter) buildListReq() *irisProtoc.CandidateListRequest  {
	req := irisProtoc.CandidateListRequest{
		Address: "",

		Page: DefaultPage,
		PerPage: DefaultPerPage,
		Sort: DefaultSorts,
		Q: "",
	}

	return &req
}

func (h ValidatorRouter) buildExRateReq() *irisProtoc.ExRateRequest  {
	req := irisProtoc.ExRateRequest{
		ValidatorAddress: "",
	}

	return &req
}

func (h ValidatorRouter) GetList(c irisProtoc.IRISHubServiceClient)  {
	res, err := c.GetCandidateList(DefaultCtx, h.buildListReq())
	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println(helper.ToJson(res))
}

func (h ValidatorRouter) GetValidatorExRate(c irisProtoc.IRISHubServiceClient)  {
	c.GetExRate(DefaultCtx, h.buildExRateReq())
}
