package configs

import (
	"github.com/irisnet/irishub-server/env"
)

type configServer struct {
	LCDServer     string
	RpcServerPort uint64
	RainbowServer string
}

var ServerConfig configServer

func init() {
	var (
		rpcServerPort uint64 = 9080
		lcdServer            = "http://35.220.215.42:30417"
		rainbowServer        = "http://192.168.150.7:30101"
	)

	if env.LCDServer != "" {
		lcdServer = env.LCDServer
	}

	if env.RainbowServer != "" {
		rainbowServer = env.RainbowServer
	}

	ServerConfig = configServer{
		LCDServer:     lcdServer,
		RpcServerPort: rpcServerPort,
		RainbowServer: rainbowServer,
	}
}
