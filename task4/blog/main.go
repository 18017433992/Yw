package main

import (
	"example.com/blog/config"
	"example.com/blog/router"
)

func main() {
	config.InitConfig() //初始化config和数据库

	r := router.SetupRouter() //初始化路由
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	// 监听并在 0.0.0.0:3000 上启动服务
	r.Run(port)
}
