package main

import (
	"exchange_backend/config"
	"exchange_backend/router"
	"log"
)

func main() {
	// 假设 InitConfig 返回 error
	config.InitConfig()

	r := router.SetupRouter()
	port := config.AppConfig.App.Port

	// 检查 r.Run 的错误
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed to run on port %s: %v", port, err)
	}
}
