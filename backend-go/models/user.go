package models

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint64    `gorm:"primaryKey;column:id" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null;column:username" json:"username"`
	Password  string    `gorm:"not null;column:password" json:"password"`
	Role      string    `gorm:"column:role;default:'USER'" json:"role"`
	Email     string    `gorm:"column:email" json:"email"`
	Enabled   bool      `gorm:"column:enabled;default:true" json:"enabled"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}