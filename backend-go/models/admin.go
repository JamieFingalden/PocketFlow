package models

// Admin 管理员模型
type Admin struct {
	ID       uint64 `gorm:"primaryKey;column:id" json:"id"`
	UserID   uint64 `gorm:"not null;column:user_id" json:"user_id"`
	Username string `gorm:"uniqueIndex;not null;column:username" json:"username"`
	Password string `gorm:"not null;column:password" json:"password"`
	Role     string `gorm:"column:role;default:'ADMIN'" json:"role"`
	Email    string `gorm:"column:email" json:"email"`
	Enabled  bool   `gorm:"column:enabled;default:true" json:"enabled"`
}

// AdminLoginRequest 管理员登录请求
type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminUserInfo 管理员用户信息
type AdminUserInfo struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Enabled  bool   `json:"enabled"`
}