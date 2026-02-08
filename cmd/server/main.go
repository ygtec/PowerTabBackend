package main

import (
	"log"
	"os"

	"powertab/config"
	"powertab/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	dsn := config.CreateDSN("root", "root", "localhost", "3306", "powertab")
	config.InitDB(dsn)

	// 创建 Gin 路由
	ginRouter := gin.Default()

	// 配置所有路由
	router.SetupRouter(ginRouter)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := ginRouter.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
