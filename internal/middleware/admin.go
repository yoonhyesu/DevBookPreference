package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 관리자 권한 검사
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("is_admin")
		if !exists || !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "관리자 권한이 필요합니다"})
			c.Abort()
			return
		}
		c.Next()
	}
}
