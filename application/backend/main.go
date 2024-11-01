package main

import (
	"carlife-backend/config"
	"carlife-backend/gateway"
	"carlife-backend/router"
)

func main() {
	// Init config
	config.InitConfig()
	gateway.InitGateways("gateway/config.json")

	// Init router
	r := router.SetupRouter()

	// Run server
	listenAddr := config.AppConfig.App.Host + ":" + config.AppConfig.App.Port
	if config.AppConfig.App.Host == "" || config.AppConfig.App.Port == "" {
		listenAddr = "localhost:8080"
	}
	r.Run(listenAddr)
}
