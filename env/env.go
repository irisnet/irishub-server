package env

import (
	"os"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
)

var (
	ENV           string
	DbAddr        string
	DbUser        string
	DbPasswd      string
	DbDatabase    string
	LCDServer     string
	RainbowServer string
)

func init() {
	env, found := os.LookupEnv(constants.ENV_NAME_ENV)
	if !found {
		ENV = constants.ENV_DEV
		logger.Info.Printf("Environment variable %v is not set, default set to %v\n",
			constants.ENV_NAME_ENV, ENV)
	} else {
		ENV = env
		logger.Info.Printf("Environment has been set to %v\n", ENV)
	}

	dbAddr, found := os.LookupEnv(constants.ENV_NAME_DB_ADDR)
	if found {
		DbAddr = dbAddr
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

	lcdServer, found := os.LookupEnv(constants.ENV_NAME_LCD_SERVER)
	if found {
		LCDServer = lcdServer
	}

	rainbowServer, found := os.LookupEnv(constants.ENV_NAME_RAINBOW_SERVER)
	if found {
		RainbowServer = rainbowServer
	}

}
