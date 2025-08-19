package middleware

import (
	"net/http"
	"backend-go/config"
	"backend-go/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// JWTAuth JWT 鉴权中间件
func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "Authorization header is required", nil))
			c.Abort()
			return
		}

		// 检查 Bearer 前缀
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "Bearer token is required", nil))
			c.Abort()
			return
		}

		// 解析和验证 JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, utils.BuildResponse(http.StatusUnauthorized, "Invalid token", nil))
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
		}

		c.Next()
	}
}