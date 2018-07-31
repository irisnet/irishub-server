package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/helper"
)

type TxRouter struct {
}

func (h TxRouter) buildListReq() *commonProtoc.TxListRequest  {
	req := commonProtoc.TxListRequest{
		Address: "",
		Type: "",
		StartTime: "",
		EndTime: "",
		Ext: []byte(""),

		Page: DefaultPage,
		PerPage: DefaultPerPage,
		Sort: DefaultSorts,
		Q: "",
	}

	return &req
}

func (h TxRouter) buildDetailReq() *commonProtoc.TxDetailRequest  {
	req := commonProtoc.TxDetailRequest{
		TxHash: "",
	}

	return &req
}

func (h TxRouter) GetTxList(c *commonProtoc.BlockChainServiceClient) {
	res, err := c.GetTxList(DefaultCtx, h.buildListReq())

	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println(helper.ToJson(res))
}

func (h TxRouter) GetTxDetail(c *commonProtoc.BlockChainServiceClient, req *commonProtoc.TxDetailRequest)  {
	res, err := c.GetTxDetail(DefaultCtx, h.buildDetailReq())

	if err != nil {
		logger.Error.Fatalln(err)
	}
	logger.Info.Println(helper.ToJson(res))
}
