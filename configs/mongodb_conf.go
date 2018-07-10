package configs

import (
	"strconv"
	
	"github.com/irisnet/irishub-server/env"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
)

type configMongodb struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

var ConfMongodb configMongodb

func init() {
	var (
		host     string
		port     int
		user     string
		password string
		dbName   string
	)

	host = "127.0.0.1"
	if env.DbHost != "" {
		host = env.DbHost
	}

	port = 27217
	if env.DbPort != "" {
		var err error
		port, err = strconv.Atoi(env.DbPort)
		if err != nil {
			logger.Error.Printf("can't convert %v to int",
				constants.ENV_NAME_DB_PORT)
		}
	}

	user = "user"
	if env.DbUser != "" {
		user = env.DbUser
	}

	password = "passwd"
	if env.DbPasswd != "" {
		password = env.DbPasswd
	}

	dbName = "sync_irishub"
	if env.DbDatabase != "" {
		dbName = env.DbDatabase
	}

	ConfMongodb = configMongodb{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}
}
