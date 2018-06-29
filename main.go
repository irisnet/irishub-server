package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	
	"git.apache.org/thrift.git/lib/go/thrift"
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server/model"
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server/model"
	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/blockchain"
	"github.com/irisnet/irishub-server/rpc/irishub"
	"github.com/irisnet/irishub-server/utils/constants"
	
	"github.com/rs/cors"
	"regexp"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	handler := cors.Default().Handler(mux)
	
	port := strconv.Itoa(int(conf.ServerConfig.RpcServerPort))
	if err := http.ListenAndServe(":" + port, handler); err != nil {
		logger.Error.Fatalln("ListenAndServe: ", err)
	}
}

func Handler(w http.ResponseWriter, req *http.Request) {
	var (
		bodyContent []byte
	)
	bodyContent, err := ioutil.ReadAll(req.Body)
	reg := regexp.MustCompile("\\\\\"")
	reg1 := regexp.MustCompile("\"\"")
	reg2 := regexp.MustCompile("\"{")
	reg3 := regexp.MustCompile("}\"")
	for reg.Find(bodyContent) != nil {
		bodyContent = reg.ReplaceAll(bodyContent, []byte("\""))
		bodyContent = reg1.ReplaceAll(bodyContent, []byte("\""))
	}
	bodyContent = reg2.ReplaceAll(bodyContent, []byte("{"))
	bodyContent = reg3.ReplaceAll(bodyContent, []byte("}"))
	if err != nil {
		// TODO: Handle exception
		logger.Error.Println(err)
		return
	}
	uri := req.RequestURI
	logger.Info.Println(uri)
	
	out := thriftRequest(bodyContent, uri)
	w.WriteHeader(constants.STATUS_CODE_OK)
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
	
	switch uri {
	case constants.UriBlockChainRPC:
		var (
			service blockchain.BlockChainRPCServices
		)
		process := commonProtoc.NewBlockChainServiceProcessor(service)
		process.Process(context.Background(), inProtocol, outProtocol)
		break
	case constants.UriIrisHubRpc:
		var (
			service irishub.IRISHubRPCSERVICES
		)
		process := irisProtoc.NewIRISHubServiceProcessor(service)
		process.Process(context.Background(), inProtocol, outProtocol)
		break
	}
	
	out := make([]byte, outBuffer.RemainingBytes())
	outBuffer.Read(out)
	return out
}