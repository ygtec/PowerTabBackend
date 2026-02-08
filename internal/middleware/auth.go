package middleware

import (
	"fmt"
	"strings"

	"powertab/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "Missing Authorization Header",
				"data": nil,
			})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  fmt.Sprintf("Invalid Token: %v", err),
				"data": nil,
			})
			c.Abort()
			return
		}

		// Store user ID in context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

// OptionalAuthMiddleware 可选的认证中间件（不强制要求登录）
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "" {
			token = strings.TrimPrefix(token, "Bearer ")
			if claims, err := utils.ParseToken(token); err == nil {
				c.Set("userID", claims.UserID)
			}
		}
		c.Next()
	}
}

// CORSMiddleware CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// ErrorHandlingMiddleware 错误处理中间件
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(500, gin.H{
					"code": 500,
					"msg":  fmt.Sprintf("Internal Server Error: %v", err),
					"data": nil,
				})
			}
		}()
		c.Next()
	}
}

// Claims JWT 声明
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
