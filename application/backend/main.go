package main

import (
	"carlife-backend/gateway"
	"carlife-backend/router"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	gateway.InitGateways("gateway/config.json")
	// 注册路由
	r := router.SetupRouter()

	// 启动服务
	r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))

}
