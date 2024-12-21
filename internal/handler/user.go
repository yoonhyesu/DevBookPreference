package handler

import (
	"DBP/ent"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 사용자 관련 핸들러 구조체
type UserHandler struct {
	client *ent.Client
}

type RegisterRequest struct {
	UserId      string `json:"USER_ID"`
	Password    string `json:"PASSWORD"`
	UserName    string `json:"USER_NAME"`
	Email       string `json:"EMAIL"`
	PhoneNumber string `json:"PHONE_NUMBER"`
}

// 새 UserHandler 인스턴스 생성
func NewUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{client: client}
}

// 비밀번호 해싱 - 평문의 패스워드에서 단방향 해시 생성
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 회원가입 핸들러
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// 요청 바디를 JSON으로 파싱
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "비밀번호 해싱 오류"})
		return
	}

	// 데이터베이스에 사용자 정보 저장 - 쿼리문
	newUser, err := h.client.User.Create().
		SetUserId(req.UserId).
		SetPassword(hashedPassword).
		SetUserName(req.UserName).
		SetEmail(req.Email).
		SetPhoneNumber(req.PhoneNumber).
		SetUserStatus(false).
		Save(context.Background())

	if err != nil {
		if ent.IsConstraintError(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "이미 존재하는 사용자입니다"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "회원가입 처리 중 오류가 발생했습니다"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "회원가입이 완료되었습니다.", "userId": newUser.UserId})
}
