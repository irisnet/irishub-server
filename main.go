package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"

	"git.apache.org/thrift.git/lib/go/thrift"
	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/blockchain"
	"github.com/irisnet/irishub-server/rpc/irishub"
	"github.com/irisnet/irishub-server/utils/constants"
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
	irisProtoc "github.com/irisnet/irisnet-rpc/irishub/codegen/server/model"

	"encoding/json"
	"fmt"
	"github.com/irisnet/irishub-server/services"
	"github.com/rs/cors"
	"net/http/pprof"
	"regexp"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.HandleFunc("/sync", syncHandler)
	handler := cors.Default().Handler(mux)

	port := strconv.Itoa(int(conf.ServerConfig.RpcServerPort))
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		logger.Error.Fatalln("ListenAndServe: ", err)
	}
}

func syncHandler(w http.ResponseWriter, rep *http.Request) {
	result := services.SyncService{}.GetCurrentSyncResult()

	str, _ := json.Marshal(result)
	w.WriteHeader(constants.STATUS_CODE_OK)
	fmt.Fprintln(w, string(str))
}

func Handler(w http.ResponseWriter, req *http.Request) {
	var (
		uri  string
		body []byte
	)
	uri = req.RequestURI

	body, err := ioutil.ReadAll(req.Body)
	body = convertReqBody(body)

	if err != nil {
		logger.Error.Println(err)
		return
	}

	out := process(body, uri)
	w.WriteHeader(constants.STATUS_CODE_OK)
	w.Write(out)
}

func process(input []byte, uri string) []byte {
	var (
		inProtocol  *thrift.TJSONProtocol
		outProtocol *thrift.TJSONProtocol
		inBuffer    thrift.TTransport
		outBuffer   thrift.TTransport
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
			service irishub.IRISHubRPCServices
		)
		process := irisProtoc.NewIRISHubServiceProcessor(service)
		process.Process(context.Background(), inProtocol, outProtocol)
		break
	default:
		return []byte("unsupported uri")
	}

	out := make([]byte, outBuffer.RemainingBytes())
	outBuffer.Read(out)
	return out
}

func convertReqBody(body []byte) []byte {
	reg1 := regexp.MustCompile("\\\\\"")
	reg2 := regexp.MustCompile("\"\"")
	reg3 := regexp.MustCompile("\"{")
	reg4 := regexp.MustCompile("}\"")
	reg5 := regexp.MustCompile("\"\\[")
	reg6 := regexp.MustCompile("]\"")
	for reg1.Find(body) != nil {
		body = reg1.ReplaceAll(body, []byte("\""))
		body = reg2.ReplaceAll(body, []byte("\""))
	}
	body = reg3.ReplaceAll(body, []byte("{"))
	body = reg4.ReplaceAll(body, []byte("}"))
	body = reg5.ReplaceAll(body, []byte("["))
	body = reg6.ReplaceAll(body, []byte("]"))

	return body
}
