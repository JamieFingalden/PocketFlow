package main

import (
	"log"
	"backend-go/config"
	"backend-go/database"
	"backend-go/router"

	"github.com/gin-gonic/gin"
)

// @title PocketFlow API
// @version 1.0
// @description 记账应用后端 API
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @tag.name bills
// @tag.description 账单相关接口

// @tag.name statistics
// @tag.description 统计相关接口

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库连接
	database.InitDB(cfg)

	// 设置 Gin 模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化 Gin 引擎
	r := gin.New()

	// 注册路由
	router.RegisterRoutes(r, cfg)

	// 启动服务器
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}