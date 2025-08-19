package api

import (
	"net/http"
	"backend-go/models"
	"backend-go/services"
	"backend-go/utils"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid request data", nil))
		return
	}

	// 调用服务层注册用户
	registeredUser, err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "User registered successfully", registeredUser))
}

// Login 用户登录
func Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponse(http.StatusBadRequest, "Invalid request data", nil))
		return
	}

	// 调用服务层登录用户
	token, err := services.LoginUser(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Login successful", gin.H{"token": token}))
}

// ProcessOCR 处理OCR请求
func ProcessOCR(c *gin.Context) {
	// 这里可以实现转发图片到Python OCR服务的逻辑
	// 暂时返回一个占位响应
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "OCR processing placeholder", nil))
}