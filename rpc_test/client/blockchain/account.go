package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

type AccountRouter struct {

}

func (h AccountRouter) buildSeqReq() *commonProtoc.SequenceRequest {
	req := commonProtoc.SequenceRequest{
		Address: "1CBD57F3B5DC9F76C7A23EDBEA05BE7A6EB2BBEA",
	}

	return &req
}

func (h AccountRouter) buildBalanceReq() *commonProtoc.BalanceRequest {
	req := commonProtoc.BalanceRequest{
		Address: "",
	}

	return &req
}

func (h AccountRouter) GetSequence(c commonProtoc.BlockChainServiceClient) {
	res, err := c.GetSequence(DefaultCtx, h.buildSeqReq())
	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println(helper.ToJson(res))
}

func (h AccountRouter) GetBalance(c commonProtoc.BlockChainServiceClient) {
	res, err := c.GetBalance(DefaultCtx, h.buildBalanceReq())
	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println(helper.ToJson(res))
}
