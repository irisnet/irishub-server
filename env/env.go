package env

import (
	"os"
	
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/utils/constants"
)

var (
	ENV        string
	DbHost     string
	DbPort     string
	ServerPort string
)

func init()  {
	ENV = os.Getenv(constants.ENV_NAME_ENV)
	DbHost = os.Getenv(constants.ENV_NAME_DB_HOST)
	DbPort = os.Getenv(constants.ENV_NAME_DB_PORT)
	ServerPort = os.Getenv(constants.ENV_NAME_SERVER_PORT)
	
	if ENV == "" {
		logger.Error.Printf("Environment variable %v is not set, default set to dev\n",
			constants.ENV_NAME_ENV)
		ENV = constants.ENV_DEV
	} else {
		logger.Info.Printf("Environment has been set to %v\n", ENV)
	}
	
}