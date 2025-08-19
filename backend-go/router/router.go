package router

import (
	"backend-go/api"
	"backend-go/config"
	"backend-go/middleware"
	_ "backend-go/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

			// 统计相关路由
			authorized.GET("/statistics/monthly-summary", api.GetMonthlySummary)
			authorized.GET("/statistics/category", api.GetCategoryStatistics)
			authorized.GET("/statistics/daily", api.GetDailyStatistics)
			authorized.GET("/statistics/monthly-chart", api.GetMonthlyChartStatistics)
		}
	}
}