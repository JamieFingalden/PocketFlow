package models

// QuickTemplate 快捷模板模型
type QuickTemplate struct {
	ID            uint64  `gorm:"primaryKey;column:id" json:"id"`
	UserID        uint64  `gorm:"not null;column:user_id" json:"user_id"`
	Name          string  `gorm:"column:name" json:"name"`
	Type          string  `gorm:"not null;column:type" json:"type"`
	CategoryID    uint64  `gorm:"not null;column:category_id" json:"category_id"`
	DefaultAmount float64 `gorm:"column:default_amount" json:"default_amount"`
	Icon          string  `gorm:"column:icon" json:"icon"`
}