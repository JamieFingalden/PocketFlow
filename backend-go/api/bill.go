package api

import (
	"net/http"
	"backend-go/models"
	"backend-go/services"
	"backend-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBill 创建账单
func CreateBill(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	var bill models.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid request data", nil))
		return
	}

	// 设置用户ID
	bill.UserID = uint64(userID.(float64))

	// 调用服务层创建账单
	createdBill, err := services.CreateBill(&bill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill created successfully", createdBill))
}

// GetBills 获取账单列表
func GetBills(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	// 调用服务层获取账单列表
	bills, err := services.GetBills(uint64(userID.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bills retrieved successfully", bills))
}

// GetBill 获取单个账单
func GetBill(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	// 获取账单ID
	billID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid bill ID", nil))
		return
	}

	// 调用服务层获取账单
	bill, err := services.GetBill(uint64(billID), uint64(userID.(float64)))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill retrieved successfully", bill))
}

// UpdateBill 更新账单
func UpdateBill(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

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
	updatedBill, err := services.UpdateBill(uint64(billID), uint64(userID.(float64)), &bill)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill updated successfully", updatedBill))
}

// DeleteBill 删除账单
func DeleteBill(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "User not found", nil))
		return
	}

	// 获取账单ID
	billID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid bill ID", nil))
		return
	}

	// 调用服务层删除账单
	err = services.DeleteBill(uint64(billID), uint64(userID.(float64)))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Bill deleted successfully", nil))
}