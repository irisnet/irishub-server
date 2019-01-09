package configs

import (
	"github.com/irisnet/irishub-server/env"
)

type configServer struct {
	LCDServer     string
	RpcServerPort uint64
}

var ServerConfig configServer

func init() {
	var (
		rpcServerPort uint64 = 9080
		lcdServer            = "http://127.0.0.1:1317"
	)

	if env.LCDServer != "" {
		lcdServer = env.LCDServer
	}

	ServerConfig = configServer{
		LCDServer:     lcdServer,
		RpcServerPort: rpcServerPort,
	}
}
