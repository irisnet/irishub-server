package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	
	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/rpc/blockchain"
	
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.ServerConfig.RpcServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	grpcServer := grpc.NewServer()
	
	var (
		blockChainRPCServices blockchain.BlockChainRPCServices
	)
	chainModel.RegisterBlockChainServiceServer(grpcServer, blockChainRPCServices)
	
	grpcServer.Serve(lis)
	fmt.Println("This is test")
}