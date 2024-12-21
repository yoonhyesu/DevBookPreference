package middleware

import (
	"DBP/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 트랜잭션 미들웨어(클로저패턴)
func TransactionMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := client.Tx(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "트랜잭션 시작 실패"})
			c.Abort() // 보류 중인 핸들러 호출 방지
			return
		}
		c.Set("tx", tx)

		defer func() {
			if v := recover(); v != nil {
				tx.Rollback()
				panic(v)
			}
		}() // 패닉이 발생했을 때만  롤백 수행

		c.Next() // 보류중인 핸들러 실행

		// 요청 처리 중 에러 있었는지 확인
		if len(c.Errors) > 0 {
			tx.Rollback()
			return
		}

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "트랜잭션 커밋 실패"})
			return
		}
	}
}

// 트랜잭션을 가져오는 헬퍼 함수
func GetTx(c *gin.Context) *ent.Tx {
	tx, _ := c.Get("tx")
	return tx.(*ent.Tx)
}
