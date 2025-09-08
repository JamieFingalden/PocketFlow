package admin

import (
	"net/http"
	"backend-go/models"
	"backend-go/services"
	"backend-go/utils"

	"github.com/gin-gonic/gin"
)

// RegisterAdmin 注册管理员用户
func RegisterAdmin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid request data", nil))
		return
	}
	
	// 设置用户角色为管理员
	user.Role = "ADMIN"
	
	// 调用服务层注册管理员用户
	registeredUser, err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Admin user registered successfully", registeredUser))
}