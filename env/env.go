package env

import (
	"os"
	
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
)

var (
	ENV        string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPasswd   string
	DbDatabase string
	AddrNodeServer string
)

func init()  {
	env, found := os.LookupEnv(constants.ENV_NAME_ENV)
	if !found {
		ENV = constants.ENV_DEV
		logger.Info.Printf("Environment variable %v is not set, default set to %v\n",
			constants.ENV_NAME_ENV, ENV)
	} else {
		ENV = env
		logger.Info.Printf("Environment has been set to %v\n", ENV)
	}
	
	dbHost, found := os.LookupEnv(constants.ENV_NAME_DB_HOST)
	if found {
		DbHost = dbHost
	}
	
	dbPort, found := os.LookupEnv(constants.ENV_NAME_DB_PORT)
	if found {
		DbPort = dbPort
	}

	dbUser, found := os.LookupEnv(constants.ENV_NAME_DB_User)
	if found {
		DbUser = dbUser
	}

	dbPasswd, found := os.LookupEnv(constants.ENV_NAME_DB_Passwd)
	if found {
		DbPasswd = dbPasswd
	}

	dbDatabase, found := os.LookupEnv(constants.ENV_NAME_DB_DATABASE)
	if found {
		DbDatabase = dbDatabase
	}
	
	addrNodeServer, found := os.LookupEnv(constants.ENV_NAME_ADDR_NODE_SERVER)
	if found {
		AddrNodeServer = addrNodeServer
	}
	
	
	
}