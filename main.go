package main

import (
	"io"
	"os"

	conf "github.com/irisnet/iris-api-server/configs"
	_ "github.com/irisnet/iris-api-server/docs"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rests"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title IRIS SERVER API
// @version 0.1.0
// @description IRIS API Server that supports various light clients

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host host
func main() {
	r := gin.New()

	// log
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	logger.Info.SetOutput(gin.DefaultWriter) // You may need this

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	rests.RegisterRoutesCandidate(r)
	rests.RegisterRoutesDelegator(r)
	rests.RegisterRoutesShare(r)

	r.Run(conf.ServerConfig.Host) // listen and serve on 0.0.0.0:8080
	logger.Info.Println("server start")
}