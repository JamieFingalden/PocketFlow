package api

import (
	"backend-go/services"
	"backend-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMonthlySummary 获取月度概览统计
// @Summary 获取月度概览统计
// @Description 获取指定年月的收入、支出和结余统计
// @Tags statistics
// @Accept json
// @Produce json
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/v1/statistics/monthly-summary [get]
func GetMonthlySummary(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	// 获取查询参数
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid year parameter", nil))
		return
	}

	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid month parameter", nil))
		return
	}

	// 验证月份参数
	if month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Month must be between 1 and 12", nil))
		return
	}

	// 调用服务层获取月度概览统计
	summary, err := services.GetMonthlySummary(uint64(userID.(float64)), year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Monthly summary retrieved successfully", summary))
}

// GetCategoryStatistics 获取分类统计
// @Summary 获取分类统计
// @Description 获取指定年月的收入或支出分类统计
// @Tags statistics
// @Accept json
// @Produce json
// @Param type query string true "类型 (income 或 expense)"
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/v1/statistics/category [get]
func GetCategoryStatistics(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	// 获取查询参数
	billType := c.Query("type")
	if billType != "income" && billType != "expense" {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Type must be 'income' or 'expense'", nil))
		return
	}

	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid year parameter", nil))
		return
	}

	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid month parameter", nil))
		return
	}

	// 验证月份参数
	if month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Month must be between 1 and 12", nil))
		return
	}

	// 调用服务层获取分类统计
	stats, err := services.GetCategoryStatistics(uint64(userID.(float64)), year, month, billType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Category statistics retrieved successfully", stats))
}

// GetDailyStatistics 获取日常统计
// @Summary 获取日常统计
// @Description 获取指定年月的每日收支统计
// @Tags statistics
// @Accept json
// @Produce json
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/v1/statistics/daily [get]
func GetDailyStatistics(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	// 获取查询参数
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid year parameter", nil))
		return
	}

	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid month parameter", nil))
		return
	}

	// 验证月份参数
	if month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Month must be between 1 and 12", nil))
		return
	}

	// 调用服务层获取日常统计
	stats, err := services.GetDailyStatistics(uint64(userID.(float64)), year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Daily statistics retrieved successfully", stats))
}

// GetMonthlyChartStatistics 获取月度图表统计
// @Summary 获取月度图表统计
// @Description 获取指定年份的月度收支统计
// @Tags statistics
// @Accept json
// @Produce json
// @Param year query int true "年份"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/v1/statistics/monthly-chart [get]
func GetMonthlyChartStatistics(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	// 获取查询参数
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid year parameter", nil))
		return
	}

	// 调用服务层获取月度图表统计
	stats, err := services.GetMonthlyChartStatistics(uint64(userID.(float64)), year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Monthly chart statistics retrieved successfully", stats))
}
