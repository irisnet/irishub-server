package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"github.com/irisnet/irishub-server/rpc/blockchain"
	"github.com/irisnet/irishub-server/rpc/irishub"
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/modules/logger"
)

func main() {
	var (
		err       error
		transport thrift.TServerTransport
	)

	addr := fmt.Sprintf("%s:%v", "localhost", conf.ServerConfig.RpcServerPort)

	protocolFactory := thrift.NewTJSONProtocolFactory()
	transportFactory := thrift.NewTTransportFactory()

	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		logger.Error.Fatalln(err)
	}

	blockChainHandler := blockchain.BlockChainRPCServices{}
	processor := commonProtoc.NewBlockChainServiceProcessor(blockChainHandler)

	irisHubHandler := irishub.IRISHubRPCServices{}
	processor2 := irisProtoc.NewIRISHubServiceProcessor(irisHubHandler)
	logger.Info.Println(processor2)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	server.Serve()
}
