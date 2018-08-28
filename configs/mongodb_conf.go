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

	host = "192.168.150.7"
	if env.DbHost != "" {
		host = env.DbHost
	}

	port = "30000"
	if env.DbPort != "" {
		port = env.DbPort
	}

	user = "iris"
	if env.DbUser != "" {
		user = env.DbUser
	}

	password = "irispassword"
	if env.DbPasswd != "" {
		password = env.DbPasswd
	}

	dbName = "sync-iris"
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
