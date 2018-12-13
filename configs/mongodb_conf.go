package configs

import (
	"github.com/irisnet/irishub-server/env"
)

type configMongodb struct {
	Addr     string
	User     string
	Password string
	DbName   string
}

var ConfMongodb configMongodb

func init() {
	var (
		addr     string
		user     string
		password string
		dbName   string
	)

	addr = "127.0.0.1:27017"
	if env.DbAddr != "" {
		addr = env.DbAddr
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
		Addr:     addr,
		User:     user,
		Password: password,
		DbName:   dbName,
	}
}
