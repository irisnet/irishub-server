package irishub

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	"github.com/irisnet/irishub-server/modules/logger"
)

var (
	DefaultCtx = context.Background()
	DefaultPage = int16(1)
	DefaultPerPage = int16(10)
	DefaultSorts = "-time"

	Client *irisProtoc.IRISHubServiceClient
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


	Client = irisProtoc.NewIRISHubServiceClient(thrift.NewTStandardClient(iprot, oprot))
}
