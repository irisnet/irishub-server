package main

import (
	"io"
	"os"

	conf "github.com/irisnet/iris-api-server/configs"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rests"
	_ "github.com/irisnet/iris-api-server/rests"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	//log
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	logger.Info.SetOutput(gin.DefaultWriter) // You may need this


	rests.RegisterRoutesCandidate(r)
	rests.RegisterRoutesDelegator(r)

	r.Run(conf.ServerConfig.Host) // listen and serve on 0.0.0.0:8080
	logger.Info.Println("server start")
}