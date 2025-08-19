package services

import (
	"backend-go/models"
	"backend-go/repository"
)

// GetMonthlySummary 获取月度概览统计服务
func GetMonthlySummary(userID uint64, year, month int) (*models.MonthlySummary, error) {
	return repository.GetMonthlySummary(userID, year, month)
}

// GetCategoryStatistics 获取分类统计服务
func GetCategoryStatistics(userID uint64, year, month int, billType string) ([]models.CategoryStatistics, error) {
	return repository.GetCategoryStatistics(userID, year, month, billType)
}

// GetDailyStatistics 获取日常统计服务
func GetDailyStatistics(userID uint64, year, month int) ([]models.DailyStatistics, error) {
	return repository.GetDailyStatistics(userID, year, month)
}

// GetMonthlyChartStatistics 获取月度图表统计服务
func GetMonthlyChartStatistics(userID uint64, year int) ([]models.DailyStatistics, error) {
	return repository.GetMonthlyChartStatistics(userID, year)
}