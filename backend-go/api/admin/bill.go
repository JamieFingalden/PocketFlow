package admin

import (
	"net/http"
	"backend-go/models"
	"backend-go/services"
	"backend-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllBills 获取所有账单列表（分页）
func GetAllBills(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	
	// 可选的用户ID过滤
	userIDStr := c.Query("user_id")
	var userID uint64
	if userIDStr != "" {
		id, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid user ID", nil))
			return
		}
		userID = uint64(id)
	}
	
	// 调用服务层获取所有账单
	bills, total, err := services.GetAllBills(page, pageSize, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	
	// 构造响应数据
	responseData := gin.H{
		"bills": bills,
		"total": total,
		"page":  page,
		"page_size": pageSize,
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bills retrieved successfully", responseData))
}

// GetBill 获取单个账单详情
func GetBill(c *gin.Context) {
	// 获取账单ID
	billID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid bill ID", nil))
		return
	}
	
	// 调用服务层获取账单详情
	bill, err := services.GetBillByID(uint64(billID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill retrieved successfully", bill))
}

// DeleteBill 删除账单（管理员权限）
func DeleteBill(c *gin.Context) {
	// 获取账单ID
	billID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid bill ID", nil))
		return
	}
	
	// 调用服务层删除账单
	err = services.DeleteBillByAdmin(uint64(billID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill deleted successfully", nil))
}

// UpdateBill 更新账单（管理员权限）
func UpdateBill(c *gin.Context) {
	// 获取账单ID
	billID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid bill ID", nil))
		return
	}
	
	var bill models.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid request data", nil))
		return
	}
	
	// 调用服务层更新账单
	updatedBill, err := services.UpdateBillByAdmin(uint64(billID), &bill)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill updated successfully", updatedBill))
}