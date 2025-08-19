package models

// MonthlySummary 月度概览统计
type MonthlySummary struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
}

// CategoryStatistics 分类统计
type CategoryStatistics struct {
	CategoryID uint64  `json:"category_id"`
	Amount     float64 `json:"amount"`
	Percentage float64 `json:"percentage"`
}

// DailyStatistics 日常统计
type DailyStatistics struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"` // "income" 或 "expense"
}

// StatisticsRequest 统计请求参数
type StatisticsRequest struct {
	Year  int `json:"year" binding:"required"`
	Month int `json:"month" binding:"required,min=1,max=12"`
}

// MonthlyChartRequest 月度图表请求参数
type MonthlyChartRequest struct {
	Year int `json:"year" binding:"required"`
}
