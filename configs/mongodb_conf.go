package configs

import (
	"github.com/irisnet/irishub-server/env"
)

type configMongodb struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

var ConfMongodb configMongodb

func init() {
	var (
		host     string
		port     string
		user     string
		password string
		dbName   string
	)

	host = "116.62.62.39"
	if env.DbHost != "" {
		host = env.DbHost
	}

	port = "27217"
	if env.DbPort != "" {
		port = env.DbPort
	}

	user = "irishub"
	if env.DbUser != "" {
		user = env.DbUser
	}

	password = "bianjie.ai"
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
