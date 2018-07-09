package configs

import (
	"strconv"
	
	"github.com/irisnet/irishub-server/env"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
)

type configMongodb struct {
	Host     string
	User     string
	Port     int
	DbName   string
	Password string
}

var ConfMongodb configMongodb

func init() {
	var (
		host     string
		user     string
		port     int
		dbName   string
		password string
	)

	switch env.ENV {
	case constants.ENV_DEV:
		host = "116.62.62.39"
		if env.DbHost != "" {
			host = env.DbHost
		}
		
		port = 27117
		if env.DbPort != "" {
			var err error
			port, err = strconv.Atoi(env.DbPort)
			if err != nil {
				logger.Error.Printf("can't convert %v to int",
					constants.ENV_NAME_DB_PORT)
			}
		} 
		
		dbName = "sync-iris-dev"
		break
	case constants.ENV_STAGE:
		host = "127.0.0.1"
		if env.DbHost != "" {
			host = env.DbHost
		}
		
		port = 27117
		if env.DbPort != "" {
			var err error
			port, err = strconv.Atoi(env.DbPort)
			if err != nil {
				logger.Error.Printf("can't convert %v to int",
					constants.ENV_NAME_DB_PORT)
			}
		} 
		
		dbName = "sync_iris"
		break
	case constants.ENV_PRO:
		host = "127.0.0.1"
		if env.DbHost != "" {
			host = env.DbHost
		}
		
		port = 27117
		if env.DbPort != "" {
			var err error
			port, err = strconv.Atoi(env.DbPort)
			if err != nil {
				logger.Error.Printf("can't convert %v to int",
					constants.ENV_NAME_DB_PORT)
			}
		}
		dbName = "sync_iris"
		break
	}

	ConfMongodb = configMongodb{
		Host:     host,
		User:     user,
		Port:     port,
		DbName:   dbName,
		Password: password,
	}
}
