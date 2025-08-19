package models

// UserProfile 用户信息模型
type UserProfile struct {
	ID               uint64  `gorm:"primaryKey;column:id" json:"id"`
	UserID           uint64  `gorm:"uniqueIndex;not null;column:user_id" json:"user_id"`
	Avatar           string  `gorm:"column:avatar" json:"avatar"`
	MonthlyBudget    float64 `gorm:"column:monthly_budget;default:0" json:"monthly_budget"`
	TotalBalance     float64 `gorm:"column:total_balance;default:0" json:"total_balance"`
	PreferredCurrency string `gorm:"column:preferred_currency;default:'CNY'" json:"preferred_currency"`
}