package models

// SystemStatistics 系统统计信息
type SystemStatistics struct {
	UserCount     int64   `json:"user_count"`
	BillCount     int64   `json:"bill_count"`
	TotalIncome   float64 `json:"total_income"`
	TotalExpense  float64 `json:"total_expense"`
}

// UserStatistics 用户统计信息
type UserStatistics struct {
	ID           uint64  `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Enabled      bool    `json:"enabled"`
	BillCount    int64   `json:"bill_count"`
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
}

// BillStatistics 账单统计信息
type BillStatistics struct {
	BillCount    int64   `json:"bill_count"`
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
}