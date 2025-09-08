package admin

import (
	"net/http"
	"backend-go/models"
	"backend-go/services"
	"backend-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers 获取所有用户列表
func GetUsers(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	
	// 调用服务层获取用户列表
	users, total, err := services.GetUsers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	
	// 构造响应数据
	responseData := gin.H{
		"users": users,
		"total": total,
		"page":  page,
		"page_size": pageSize,
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Users retrieved successfully", responseData))
}

// GetUser 获取单个用户信息
func GetUser(c *gin.Context) {
	// 获取用户ID
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid user ID", nil))
		return
	}
	
	// 调用服务层获取用户信息
	user, err := services.GetUserByID(uint64(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "User retrieved successfully", user))
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	// 获取用户ID
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid user ID", nil))
		return
	}
	
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid request data", nil))
		return
	}
	
	// 调用服务层更新用户信息
	updatedUser, err := services.UpdateUser(uint64(userID), &user)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "User updated successfully", updatedUser))
}

// DisableUser 禁用用户
func DisableUser(c *gin.Context) {
	// 获取用户ID
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid user ID", nil))
		return
	}
	
	// 调用服务层禁用用户
	err = services.DisableUser(uint64(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "User disabled successfully", nil))
}

// EnableUser 启用用户
func EnableUser(c *gin.Context) {
	// 获取用户ID
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid user ID", nil))
		return
	}
	
	// 调用服务层启用用户
	err = services.EnableUser(uint64(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildResponse(http.StatusNotFound, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "User enabled successfully", nil))
}