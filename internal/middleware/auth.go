package middleware

import (
	"DBP/internal/utils"
	"DBP/pkg/database"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			log.Printf("쿠키에서 토큰을 찾을 수 없습니다: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "인증이 필요합니다"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(accessToken)
		if err != nil {
			log.Printf("토큰 검증 실패: %v", err)
			if strings.Contains(err.Error(), "token is expired") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "토큰이 만료되었습니다", "code": "TOKEN_EXPIRED"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 토큰입니다"})
			}
			c.Abort()
			return
		}

		// Redis 검증 로직 개선
		ctx := c.Request.Context()
		userID, err := database.RedisClient.Get(ctx, claims.TokenUUID).Result()
		if err != nil {
			log.Printf("Redis에서 토큰 검증 실패: %v %s", err, claims.UserID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 세션입니다"})
			c.Abort()
			return
		}

		if userID != claims.UserID {
			log.Printf("토큰의 사용자 ID 불일치: %s != %s", userID, claims.UserID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 세션입니다"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.UserName)
		c.Set("is_admin", claims.IsAdmin)

		c.Next()
	}
}
