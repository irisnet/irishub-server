package configs

import (
	"github.com/irisnet/iris-api-server/env"
	"github.com/irisnet/iris-api-server/utils/constants"
)

type configServer struct {
	Host string
}

var ServerConfig configServer

func init() {
	var (
		host string
	)

	switch env.ENV {
	case constants.ENV_DEV:
		host = "0.0.0.0:9080"
	case constants.ENV_PRO:
		host = "0.0.0.0:80"
	}

	ServerConfig = configServer{
		Host: host,
	}
}
