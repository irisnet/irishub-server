package configs

import (
	"github.com/irisnet/iris-api-server/env"
	"github.com/irisnet/iris-api-server/utils/constants"
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
		host = "127.0.0.1"
		port = 27117
		dbName = "sync_iris"
	case constants.ENV_PRO:
		host = "127.0.0.1"
		port = 27117
		dbName = "sync_iris"
	}

	ConfMongodb = configMongodb{
		Host:     host,
		User:     user,
		Port:     port,
		DbName:   dbName,
		Password: password,
	}
}
