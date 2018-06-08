package configs

import (
	"github.com/irisnet/irishub-server/env"
)

type configServer struct {
	AddrNodeServer string
	RpcServerPort  uint64
}

var ServerConfig configServer

func init() {
	var (
		rpcServerPort uint64 = 9080
		addrNodeServer = "http://47.104.155.125:8999"
	)
	
	if env.AddrNodeServer != "" {
		addrNodeServer = env.AddrNodeServer
	}
	
	ServerConfig = configServer{
		AddrNodeServer: addrNodeServer,
		RpcServerPort:  rpcServerPort,
	}
}
