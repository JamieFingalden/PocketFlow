package services

import (
	"errors"
	"backend-go/models"
	"backend-go/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser 用户注册服务
func RegisterUser(user *models.User) (*models.User, error) {
	// 检查用户名是否已存在
	existingUser, err := repository.GetUserByUsername(user.Username)
	if err == nil && existingUser.ID > 0 {
		return nil, errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	existingUser, err = repository.GetUserByEmail(user.Email)
	if err == nil && existingUser.ID > 0 {
		return nil, errors.New("email already exists")
	}

	// 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// 创建用户
	return repository.CreateUser(user)
}

// LoginUser 用户登录服务
func LoginUser(username, password string) (string, error) {
	// 根据用户名获取用户
	user, err := repository.GetUserByUsername(username)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid username or password")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// 生成JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// 这里应该从配置中获取密钥，暂时使用硬编码
	tokenString, err := token.SignedString([]byte("pocketflow-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}