package main

import (
	"context"
	"net/http"
	
	"git.apache.org/thrift.git/lib/go/thrift"
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/blockchain"
	"github.com/irisnet/irishub-server/rpc/irishub"
)

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":" + string(conf.ServerConfig.RpcServerPort), nil)
	if err != nil {
		logger.Error.Fatalln("ListenAndServe: ", err)
	}
}

func Handler(w http.ResponseWriter, req *http.Request) {
	var (
		bodyContent []byte
	)
	
	_, err := req.Body.Read(bodyContent)
	if err != nil {
		logger.Error.Println(err)
	}
	
	uri := req.RequestURI
	
	out := thriftRequest(bodyContent, uri)
	println(string(out))
	w.WriteHeader(200)
	w.Write(out)
}

func thriftRequest(input []byte, uri string) []byte {
	var (
		inProtocol *thrift.TJSONProtocol
		outProtocol *thrift.TJSONProtocol
		inBuffer thrift.TTransport
		outBuffer thrift.TTransport
	)
	
	
	inBuffer = thrift.NewTMemoryBuffer()
	inBuffer.Write(input)
	if inBuffer != nil {
		defer inBuffer.Close()
	}
	
	outBuffer = thrift.NewTMemoryBuffer()
	if outBuffer != nil {
		defer outBuffer.Close()
	}
	
	inProtocol = thrift.NewTJSONProtocol(inBuffer)
	outProtocol = thrift.NewTJSONProtocol(outBuffer)
	
	if uri != "" {
		switch uri {
		case "blockchain":
			var (
				service blockchain.BlockChainRPCServices
			)
			process := commonProtoc.NewBlockChainServiceProcessor(service)
			process.Process(context.Background(), inProtocol, outProtocol)
			break
		case "irishub":
			var (
				service irishub.IRISHubRPCSERVICES
			)
			process := irisProtoc.NewIRISHubServiceProcessor(service)
			process.Process(context.Background(), inProtocol, outProtocol)
			break
		}
	}
	
	out := make([]byte, outBuffer.RemainingBytes())
	outBuffer.Read(out)
	return out
	
}