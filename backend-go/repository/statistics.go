package repository

import (
	"backend-go/database"
	"backend-go/models"
	"time"
)

// GetMonthlySummary 获取月度概览统计
func GetMonthlySummary(userID uint64, year, month int) (*models.MonthlySummary, error) {
	var summary models.MonthlySummary
	
	// 构建日期范围
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)
	
	// 查询收入总额
	var totalIncome float64
	err := database.DB.Model(&models.Bill{}).
		Where("user_id = ? AND type = ? AND bill_time >= ? AND bill_time < ?", userID, "income", startDate, endDate).
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&totalIncome)
	
	if err != nil {
		return nil, err
	}
	
	// 查询支出总额
	var totalExpense float64
	err = database.DB.Model(&models.Bill{}).
		Where("user_id = ? AND type = ? AND bill_time >= ? AND bill_time < ?", userID, "expense", startDate, endDate).
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&totalExpense)
	
	if err != nil {
		return nil, err
	}
	
	summary.TotalIncome = totalIncome
	summary.TotalExpense = totalExpense
	summary.Balance = totalIncome - totalExpense
	
	return &summary, nil
}

// GetCategoryStatistics 获取分类统计
func GetCategoryStatistics(userID uint64, year, month int, billType string) ([]models.CategoryStatistics, error) {
	var stats []models.CategoryStatistics
	
	// 构建日期范围
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)
	
	// 查询分类统计
	result := database.DB.Model(&models.Bill{}).
		Select("category_id, COALESCE(SUM(amount), 0) as amount").
		Where("user_id = ? AND type = ? AND bill_time >= ? AND bill_time < ?", userID, billType, startDate, endDate).
		Group("category_id").
		Order("amount DESC").
		Scan(&stats)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	// 计算总额用于计算百分比
	var totalAmount float64
	for _, stat := range stats {
		totalAmount += stat.Amount
	}
	
	// 计算百分比
	for i := range stats {
		if totalAmount > 0 {
			stats[i].Percentage = stats[i].Amount / totalAmount * 100
		}
	}
	
	return stats, nil
}

// GetDailyStatistics 获取日常统计
func GetDailyStatistics(userID uint64, year, month int) ([]models.DailyStatistics, error) {
	var stats []models.DailyStatistics
	
	// 构建日期范围
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)
	
	// 查询每日统计
	result := database.DB.Model(&models.Bill{}).
		Select("DATE(bill_time) as date, type, COALESCE(SUM(amount), 0) as amount").
		Where("user_id = ? AND bill_time >= ? AND bill_time < ?", userID, startDate, endDate).
		Group("DATE(bill_time), type").
		Order("date").
		Scan(&stats)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return stats, nil
}

// GetMonthlyChartStatistics 获取月度图表统计
func GetMonthlyChartStatistics(userID uint64, year int) ([]models.DailyStatistics, error) {
	var stats []models.DailyStatistics
	
	// 构建日期范围
	startDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(1, 0, 0)
	
	// 查询月度统计
	result := database.DB.Model(&models.Bill{}).
		Select("MONTH(bill_time) as date, type, COALESCE(SUM(amount), 0) as amount").
		Where("user_id = ? AND bill_time >= ? AND bill_time < ?", userID, startDate, endDate).
		Group("MONTH(bill_time), type").
		Order("date").
		Scan(&stats)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return stats, nil
}