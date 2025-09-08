package admin

import (
	"net/http"
	"backend-go/services"
	"backend-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSystemStatistics 获取系统整体统计信息
func GetSystemStatistics(c *gin.Context) {
	// 调用服务层获取系统统计信息
	stats, err := services.GetSystemStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "System statistics retrieved successfully", stats))
}

// GetUserStatistics 获取用户统计信息
func GetUserStatistics(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	
	// 调用服务层获取用户统计信息
	stats, total, err := services.GetUserStatistics(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	
	// 构造响应数据
	responseData := gin.H{
		"statistics": stats,
		"total": total,
		"page":  page,
		"page_size": pageSize,
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "User statistics retrieved successfully", responseData))
}

// GetBillStatistics 获取账单统计信息
func GetBillStatistics(c *gin.Context) {
	// 获取查询参数
	year, _ := strconv.Atoi(c.DefaultQuery("year", "0"))
	month, _ := strconv.Atoi(c.DefaultQuery("month", "0"))
	
	// 调用服务层获取账单统计信息
	stats, err := services.GetBillStatistics(year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill statistics retrieved successfully", stats))
}