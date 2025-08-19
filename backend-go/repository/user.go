package repository

import (
	"backend-go/database"
	"backend-go/models"
)

// CreateUser 创建用户
func CreateUser(user *models.User) (*models.User, error) {
	result := database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail 根据邮箱获取用户
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID 根据ID获取用户
func GetUserByID(id uint64) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}