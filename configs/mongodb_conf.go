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
		host = "116.62.62.39"
		user = "postgres"
		port = 27117
		dbName = "sync_iris"
		password = "admino0o0oo0"
	case constants.ENV_PRO:
		host = "127.0.0.1"
		user = "postgres"
		port = 27117
		dbName = "sync_iris"
		password = "admino0o0oo0"
	}

	ConfMongodb = configMongodb{
		Host:     host,
		User:     user,
		Port:     port,
		DbName:   dbName,
		Password: password,
	}
}
