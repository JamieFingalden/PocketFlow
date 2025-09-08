package admin

import (
	"net/http"
	"backend-go/models"
	"backend-go/services"
	"backend-go/utils"

	"github.com/gin-gonic/gin"
)

// AdminLogin 管理员登录
func AdminLogin(c *gin.Context) {
	var loginReq models.AdminLoginRequest
	
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid request data", nil))
		return
	}
	
	// 调用服务层进行管理员登录
	token, err := services.AdminLogin(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Admin login successful", gin.H{"token": token}))
}

// AdminLogout 管理员登出
func AdminLogout(c *gin.Context) {
	// 在JWT模式下，登出通常在客户端删除token即可
	// 服务端可以维护一个黑名单，但这里简化处理
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Admin logout successful", nil))
}