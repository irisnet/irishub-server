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
		rpcServerPort  uint64 = 9080
		addrNodeServer        = "http://192.168.150.7:1317"
	)

	if env.AddrNodeServer != "" {
		addrNodeServer = env.AddrNodeServer
	}

	ServerConfig = configServer{
		AddrNodeServer: addrNodeServer,
		RpcServerPort:  rpcServerPort,
	}
}
