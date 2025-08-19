package repository

import (
	"backend-go/database"
	"backend-go/models"
)

// CreateBill 创建账单
func CreateBill(bill *models.Bill) (*models.Bill, error) {
	result := database.DB.Create(bill)
	if result.Error != nil {
		return nil, result.Error
	}
	return bill, nil
}

// GetBillsByUserID 根据用户ID获取账单列表
func GetBillsByUserID(userID uint64) ([]models.Bill, error) {
	var bills []models.Bill
	result := database.DB.Where("user_id = ?", userID).Find(&bills)
	if result.Error != nil {
		return nil, result.Error
	}
	return bills, nil
}

// GetBillByID 根据ID获取账单
func GetBillByID(id uint64) (*models.Bill, error) {
	var bill models.Bill
	result := database.DB.First(&bill, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bill, nil
}

// UpdateBill 更新账单
func UpdateBill(bill *models.Bill) (*models.Bill, error) {
	result := database.DB.Save(bill)
	if result.Error != nil {
		return nil, result.Error
	}
	return bill, nil
}

// DeleteBill 删除账单
func DeleteBill(id uint64) error {
	result := database.DB.Delete(&models.Bill{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}