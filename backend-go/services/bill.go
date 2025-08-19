package services

import (
	"errors"
	"backend-go/models"
	"backend-go/repository"
)

// CreateBill 创建账单服务
func CreateBill(bill *models.Bill) (*models.Bill, error) {
	return repository.CreateBill(bill)
}

// GetBills 获取账单列表服务
func GetBills(userID uint64) ([]models.Bill, error) {
	return repository.GetBillsByUserID(userID)
}

// GetBill 获取单个账单服务
func GetBill(billID, userID uint64) (*models.Bill, error) {
	bill, err := repository.GetBillByID(billID)
	if err != nil {
		return nil, err
	}

	// 验证账单是否属于该用户
	if bill.UserID != userID {
		return nil, errors.New("bill not found")
	}

	return bill, nil
}

// UpdateBill 更新账单服务
func UpdateBill(billID, userID uint64, bill *models.Bill) (*models.Bill, error) {
	// 首先获取原始账单
	existingBill, err := repository.GetBillByID(billID)
	if err != nil {
		return nil, err
	}

	// 验证账单是否属于该用户
	if existingBill.UserID != userID {
		return nil, errors.New("bill not found")
	}

	// 更新账单
	existingBill.Type = bill.Type
	existingBill.CategoryID = bill.CategoryID
	existingBill.Amount = bill.Amount
	existingBill.Remark = bill.Remark
	existingBill.BillTime = bill.BillTime

	return repository.UpdateBill(existingBill)
}

// DeleteBill 删除账单服务
func DeleteBill(billID, userID uint64) error {
	// 首先获取原始账单
	bill, err := repository.GetBillByID(billID)
	if err != nil {
		return err
	}

	// 验证账单是否属于该用户
	if bill.UserID != userID {
		return errors.New("bill not found")
	}

	// 删除账单
	return repository.DeleteBill(billID)
}