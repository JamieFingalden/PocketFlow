package models

import (
	"time"
)

// Bill 账单模型
type Bill struct {
	ID         uint64    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint64    `gorm:"not null;column:user_id" json:"user_id"`
	Type       string    `gorm:"not null;column:type" json:"type"` // "income" 或 "expense"
	CategoryID uint64    `gorm:"not null;column:category_id" json:"category_id"`
	Amount     float64   `gorm:"not null;column:amount" json:"amount"`
	Remark     string    `gorm:"column:remark" json:"remark"`
	BillTime   time.Time `gorm:"not null;column:bill_time" json:"bill_time"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName 指定表名
func (Bill) TableName() string {
	return "bills"
}