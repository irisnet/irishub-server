package main

import (
	"io"
	"os"
	
	conf "github.com/irisnet/iris-api-server/configs"
	_ "github.com/irisnet/iris-api-server/docs"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/rests"


"github.com/gin-contrib/cors"
"github.com/gin-gonic/gin"
"github.com/swaggo/gin-swagger"
"github.com/swaggo/gin-swagger/swaggerFiles"

)

// @title IRIS SERVER API
// @version 0.1.0
// @description IRIS API Server that supports various light clients

// @license.name API Spec Document
// @license.url https://github.com/kaifei-bianjie/share/blob/master/api_spec.md

// @host host
func main() {
	r := gin.New()

	// log
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	logger.Info.SetOutput(gin.DefaultWriter)
	
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
	}))

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	rests.RegisterRoutesCandidate(r)
	rests.RegisterRoutesDelegator(r)
	rests.RegisterRoutesShare(r)
	rests.RegisterStakeTxRoute(r)
	rests.RegisterCommonTxRoute(r)

	r.Run(conf.ServerConfig.Host) // listen and serve on 0.0.0.0:8080
	logger.Info.Println("server start")
}