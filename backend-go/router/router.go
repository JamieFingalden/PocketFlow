package router

import (
	"backend-go/api"
	"backend-go/config"
	"backend-go/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    "OK",
		})
	})

	// OCR 代理接口
	r.POST("/ocr/process", api.ProcessOCR)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 用户相关路由
		v1.POST("/register", api.Register)
		v1.POST("/login", api.Login)

		// 需要 JWT 验证的路由组
		authorized := v1.Group("/")
		authorized.Use(middleware.JWTAuth(cfg))
		{
			// 账单相关路由
			authorized.POST("/bills", api.CreateBill)
			authorized.GET("/bills", api.GetBills)
			authorized.GET("/bills/:id", api.GetBill)
			authorized.PUT("/bills/:id", api.UpdateBill)
			authorized.DELETE("/bills/:id", api.DeleteBill)
		}
	}
}