package router

import (
	"backend-go/api"
	"backend-go/api/admin"
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

		// 管理员相关路由
		v1.POST("/admin/login", admin.AdminLogin)
		v1.POST("/admin/logout", admin.AdminLogout)

		// 需要 JWT 验证的路由组（普通用户）
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

		// 需要管理员权限验证的路由组
		adminAuthorized := v1.Group("/admin")
		adminAuthorized.Use(middleware.AdminAuth(cfg))
		{
			// 用户管理路由
			adminAuthorized.GET("/users", admin.GetUsers)
			adminAuthorized.GET("/users/:id", admin.GetUser)
			adminAuthorized.PUT("/users/:id", admin.UpdateUser)
			adminAuthorized.PUT("/users/:id/disable", admin.DisableUser)
			adminAuthorized.PUT("/users/:id/enable", admin.EnableUser)

			// 账单管理路由
			adminAuthorized.GET("/bills", admin.GetAllBills)
			adminAuthorized.GET("/bills/:id", admin.GetBill)
			adminAuthorized.PUT("/bills/:id", admin.UpdateBill)
			adminAuthorized.DELETE("/bills/:id", admin.DeleteBill)

			// 统计管理路由
			adminAuthorized.GET("/statistics/system", admin.GetSystemStatistics)
			adminAuthorized.GET("/statistics/users", admin.GetUserStatistics)
			adminAuthorized.GET("/statistics/bills", admin.GetBillStatistics)
		}
	}
}