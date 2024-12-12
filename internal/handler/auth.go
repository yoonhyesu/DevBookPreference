package handler

import (
	"DBP/ent"
	"DBP/ent/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 인증 관련 핸들러 구조체
type AuthHandler struct {
	client *ent.Client
}

// 새 AuthHandler 인스턴스 생성
func NewAuthHandler(client *ent.Client) *AuthHandler {
	return &AuthHandler{client: client}
}

// 아이디 중복 체크 핸들러
func (h *AuthHandler) CheckDuplicateID(c *gin.Context) {
	var req struct {
		UserId string `json:"USER_ID"` // 클라이언트에서 보낸 JSON userid필드 바인딩
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 요청입니다",
		})
		return
	}

	// 데이터베이스에서 해당 아이디 존재 여부
	exists, err := h.client.User.Query().Where(user.UserIDEQ(req.UserId)).Exist(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "서버 오류 발생",
		})
		return
	}

	// 결과 반환
	c.JSON(http.StatusOK, gin.H{
		"exists": exists, //true면 이미존재, false면 사용 가능
	})

}
