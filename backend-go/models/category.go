package models

// Category 分类模型
type Category struct {
	ID     uint64 `gorm:"primaryKey;column:id" json:"id"`
	Name   string `gorm:"not null;column:name" json:"name"`
	Type   string `gorm:"not null;column:type" json:"type"`
	Icon   string `gorm:"column:icon" json:"icon"`
	UserID *uint64 `gorm:"column:user_id" json:"user_id"`
}