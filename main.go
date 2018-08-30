package main

import (
	"bitbucket.org/blockchain/controllers"
	"bitbucket.org/blockchain/environment"
	"github.com/facebookgo/grace/gracehttp"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	environment.Instance()
	echo := environment.Instance().E
	group := echo.Group("/v1")
	controller := controllers.Controller{}
	group.GET("/blocks", controller.GetBlock)
	group.POST("/mine", controller.AddBlock)
	group.GET("/ws", controller.CreateWebSocketConnection)
	group.POST("/wss", controller.RegisterWebSocket)

	port := environment.Instance().Get("server.port")
	log.Infoln("[BlockChain] Server listening on: ", port)
	echo.Server.Addr = port.(string)
	echo.Logger.Fatal(gracehttp.Serve(echo.Server))
}
