package repository

import (
	"backend-go/models"
	"backend-go/database"
)

// GetUsers 获取用户列表
func GetUsers(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	
	offset := (page - 1) * pageSize
	
	// 获取总数
	if err := database.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 获取用户列表
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	
	return users, total, nil
}

// UpdateUser 更新用户
func UpdateUser(user *models.User) (*models.User, error) {
	if err := database.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllBills 获取所有账单
func GetAllBills(page, pageSize int, userID uint64) ([]models.Bill, int64, error) {
	var bills []models.Bill
	var total int64
	
	offset := (page - 1) * pageSize
	
	db := database.DB.Model(&models.Bill{})
	
	// 如果指定了用户ID，则过滤
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}
	
	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 获取账单列表
	if err := db.Offset(offset).Limit(pageSize).Find(&bills).Error; err != nil {
		return nil, 0, err
	}
	
	return bills, total, nil
}

// GetSystemStatistics 获取系统统计信息
func GetSystemStatistics() (*models.SystemStatistics, error) {
	var stats models.SystemStatistics
	
	// 获取用户总数
	if err := database.DB.Model(&models.User{}).Count(&stats.UserCount).Error; err != nil {
		return nil, err
	}
	
	// 获取账单总数
	if err := database.DB.Model(&models.Bill{}).Count(&stats.BillCount).Error; err != nil {
		return nil, err
	}
	
	// 获取总收入
	if err := database.DB.Model(&models.Bill{}).Where("type = ?", "income").Select("COALESCE(SUM(amount), 0)").Scan(&stats.TotalIncome).Error; err != nil {
		return nil, err
	}
	
	// 获取总支出
	if err := database.DB.Model(&models.Bill{}).Where("type = ?", "expense").Select("COALESCE(SUM(amount), 0)").Scan(&stats.TotalExpense).Error; err != nil {
		return nil, err
	}
	
	return &stats, nil
}

// GetUserStatistics 获取用户统计信息
func GetUserStatistics(page, pageSize int) ([]models.UserStatistics, int64, error) {
	var userStats []models.UserStatistics
	var total int64
	
	offset := (page - 1) * pageSize
	
	// 获取总数
	if err := database.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 获取用户统计信息
	query := `
		SELECT 
			u.id,
			u.username,
			u.email,
			u.enabled,
			COUNT(b.id) as bill_count,
			COALESCE(SUM(CASE WHEN b.type = 'income' THEN b.amount ELSE 0 END), 0) as total_income,
			COALESCE(SUM(CASE WHEN b.type = 'expense' THEN b.amount ELSE 0 END), 0) as total_expense
		FROM users u
		LEFT JOIN bills b ON u.id = b.user_id
		GROUP BY u.id, u.username, u.email, u.enabled
		ORDER BY u.id
		LIMIT ? OFFSET ?
	`
	
	if err := database.DB.Raw(query, pageSize, offset).Scan(&userStats).Error; err != nil {
		return nil, 0, err
	}
	
	return userStats, total, nil
}

// GetBillStatistics 获取账单统计信息
func GetBillStatistics(year, month int) (*models.BillStatistics, error) {
	var stats models.BillStatistics
	
	db := database.DB.Model(&models.Bill{})
	
	// 如果指定了年份，则过滤
	if year > 0 {
		db = db.Where("YEAR(bill_time) = ?", year)
	}
	
	// 如果指定了月份，则过滤
	if month > 0 {
		db = db.Where("MONTH(bill_time) = ?", month)
	}
	
	// 获取总收入
	if err := db.Where("type = ?", "income").Select("COALESCE(SUM(amount), 0)").Scan(&stats.TotalIncome).Error; err != nil {
		return nil, err
	}
	
	// 获取总支出
	if err := db.Where("type = ?", "expense").Select("COALESCE(SUM(amount), 0)").Scan(&stats.TotalExpense).Error; err != nil {
		return nil, err
	}
	
	// 获取账单总数
	if err := db.Count(&stats.BillCount).Error; err != nil {
		return nil, err
	}
	
	return &stats, nil
}