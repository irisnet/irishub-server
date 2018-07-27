package blockchain

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/modules/logger"
)

var (
	DefaultCtx = context.Background()
	DefaultPage = int64(1)
	DefaultPerPage = int64(10)
	DefaultSorts = "-time"
	Client *commonProtoc.BlockChainServiceClient
)

func init()  {
	var (
		addr = "localhost:9080"
		transport thrift.TTransport
		err error
	)
	protocolFactory := thrift.NewTJSONProtocolFactory()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())


	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		logger.Error.Fatalln(err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)


	Client = commonProtoc.NewBlockChainServiceClient(thrift.NewTStandardClient(iprot, oprot))
}
