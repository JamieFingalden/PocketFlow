package services

import (
	"errors"
	"backend-go/models"
	"backend-go/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// AdminLogin 管理员登录服务
func AdminLogin(username, password string) (string, error) {
	// 根据用户名获取管理员用户
	user, err := repository.GetUserByUsername(username)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid username or password")
	}
	
	// 检查用户是否为管理员角色
	if user.Role != "ADMIN" {
		return "", errors.New("user is not an administrator")
	}
	
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}
	
	// 生成JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	
	// 这里应该从配置中获取密钥，暂时使用硬编码
	tokenString, err := token.SignedString([]byte("pocketflow-secret-key"))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}

// GetUsers 获取用户列表服务
func GetUsers(page, pageSize int) ([]models.User, int64, error) {
	return repository.GetUsers(page, pageSize)
}

// GetUserByID 根据ID获取用户服务
func GetUserByID(userID uint64) (*models.User, error) {
	return repository.GetUserByID(userID)
}

// UpdateUser 更新用户信息服务
func UpdateUser(userID uint64, user *models.User) (*models.User, error) {
	// 首先获取原始用户
	existingUser, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	
	// 更新用户信息
	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Enabled = user.Enabled
	
	// 如果密码不为空，则更新密码
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		existingUser.Password = string(hashedPassword)
	}
	
	return repository.UpdateUser(existingUser)
}

// DisableUser 禁用用户服务
func DisableUser(userID uint64) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}
	
	user.Enabled = false
	_, err = repository.UpdateUser(user)
	return err
}

// EnableUser 启用用户服务
func EnableUser(userID uint64) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}
	
	user.Enabled = true
	_, err = repository.UpdateUser(user)
	return err
}

// GetAllBills 获取所有账单服务
func GetAllBills(page, pageSize int, userID uint64) ([]models.Bill, int64, error) {
	return repository.GetAllBills(page, pageSize, userID)
}

// GetBillByID 根据ID获取账单服务
func GetBillByID(billID uint64) (*models.Bill, error) {
	return repository.GetBillByID(billID)
}

// DeleteBillByAdmin 管理员删除账单服务
func DeleteBillByAdmin(billID uint64) error {
	return repository.DeleteBill(billID)
}

// UpdateBillByAdmin 管理员更新账单服务
func UpdateBillByAdmin(billID uint64, bill *models.Bill) (*models.Bill, error) {
	// 首先获取原始账单
	existingBill, err := repository.GetBillByID(billID)
	if err != nil {
		return nil, err
	}
	
	// 更新账单信息
	existingBill.Type = bill.Type
	existingBill.CategoryID = bill.CategoryID
	existingBill.Amount = bill.Amount
	existingBill.Remark = bill.Remark
	existingBill.BillTime = bill.BillTime
	
	return repository.UpdateBill(existingBill)
}

// GetSystemStatistics 获取系统统计信息服务
func GetSystemStatistics() (*models.SystemStatistics, error) {
	return repository.GetSystemStatistics()
}

// GetUserStatistics 获取用户统计信息服务
func GetUserStatistics(page, pageSize int) ([]models.UserStatistics, int64, error) {
	return repository.GetUserStatistics(page, pageSize)
}

// GetBillStatistics 获取账单统计信息服务
func GetBillStatistics(year, month int) (*models.BillStatistics, error) {
	return repository.GetBillStatistics(year, month)
}