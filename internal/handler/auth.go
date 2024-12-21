package handler

import (
	"DBP/ent"
	"DBP/ent/user"
	"DBP/internal/utils"
	"DBP/pkg/database"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 인증 관련 핸들러 구조체
type AuthHandler struct {
	client *ent.Client
}

// 새 AuthHandler 인스턴스 생성
func NewAuthHandler(client *ent.Client) *AuthHandler {
	return &AuthHandler{client: client}
}

type CheckDuplicateIDRequest struct {
	UserId string `json:"USER_ID"`
}

// ✔️회원가입 시 아이디 중복 체크
func (h *AuthHandler) CheckDuplicateID(c *gin.Context) {
	log.Println("아이디 중복 체크 요청 받음")

	var req CheckDuplicateIDRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 요청입니다",
		})
		return
	}

	// 데이터베이스에서 해당 아이디 존재 여부
	exists, err := h.client.User.Query().Where(user.UserId(req.UserId)).Exist(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "서버 오류 발생",
		})
		return
	}
	log.Printf("Checking ID: %s", req.UserId)
	// 결과 반환
	c.JSON(http.StatusOK, gin.H{
		"exists": exists, //true면 이미존재, false면 사용 가능
	})

}

// 로그인 요청 구조체
type SignInRequest struct {
	UserId   string `json:"USER_ID"`
	Password string `json:"PASSWORD"`
}

// ✔️로그인
func (h *AuthHandler) SignIn(c *gin.Context) { // AuthHandler의 리시버
	var req SignInRequest

	// json이 똑바로 들어왔는지
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}

	// 들어온 데이터가 db 사용자정보와 일치하는지 비교
	user, err := h.client.User.Query().Where(user.UserId(req.UserId)).First(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "아이디 또는 비밀번호가 일치하지 않습니다"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "아이디 또는 비밀번호가 일치하지 않습니다"})
		return
	}

	// 관리자 권한 확인 (IS_ADMIN이 1인 경우만 관리자)
	isAdmin := user.IsAdmin == true
	log.Printf("사용자 권한 확인 - UserID: %s, IsAdmin: %v", user.UserId, isAdmin)

	// 토큰 생성 (관리자 권한 정보 포함)
	accessToken, refreshToken, err := utils.GenerateTokenPair(user.UserId, user.UserName, isAdmin)
	if err != nil {
		log.Printf("토큰 생성 실패: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "토큰 발행 오류"})
		return
	}

	// 토큰 claims 확인
	claims, _ := utils.ValidateToken(accessToken)
	log.Printf("로그인 시도 - UserID: %s, TokenUUID: %s", user.UserId, claims.TokenUUID)

	// Redis에 토큰 메타데이터 저장 (TokenUUID를 key로 사용)
	ctx := context.Background()
	if err := database.SaveTokenMetaData(ctx, user.UserId, claims.TokenUUID, utils.AccessTokenDuration); err != nil {
		log.Printf("Redis 토큰 저장 실패: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "토큰 메타데이터 저장 오류"})
		return
	}

	// Redis에 저장된 데이터 확인
	storedUserID, err := database.RedisClient.Get(ctx, claims.TokenUUID).Result()
	if err != nil {
		log.Printf("Redis 저장 확인 실패: %v", err)
	} else {
		log.Printf("Redis에 저장된 데이터 확인 - TokenUUID: %s, UserID: %s", claims.TokenUUID, storedUserID)
	}

	// 쿠키에 토큰 설정
	c.SetCookie("access_token", accessToken, int(utils.AccessTokenDuration.Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, int(utils.RefreshTokenDuration.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "로그인 성공",
		"user": gin.H{
			"UserName": user.UserName,
			"IsAdmin":  isAdmin,
		},
	})
}

// ✔️로그아웃
func (h *AuthHandler) SignOut(c *gin.Context) {
	accessToken, _ := c.Cookie("access_token")

	if claims, err := utils.ValidateToken(accessToken); err == nil {
		log.Printf("로그아웃 시도 - TokenUUID: %s", claims.TokenUUID)

		// Redis에서 토큰 삭제
		ctx := context.Background()
		err := database.DeleteTokenMetaData(ctx, claims.TokenUUID)
		if err != nil {
			log.Printf("Redis에서 토큰 삭제 실패: %v", err)
		} else {
			log.Printf("Redis에서 토큰 삭제 성공 - TokenUUID: %s", claims.TokenUUID)
		}
	}

	// 쿠키 삭제 (항상 실행)
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "성공적으로 로그아웃 되었습니다"})
}

// 토큰 자동 갱신
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "리프레시 토큰이 없습니다"})
		return
	}
	// 리프레시 토큰 검증 후 claims얻기
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 리프레�� 토큰입니다"})
		return
	}

	//레디스에서 토큰 유효성 확인
	ctx := context.Background()
	userID, err := database.RedisClient.Get(ctx, claims.TokenUUID).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "만료된 세션입니다"})
		return
	}

	// DB에서 사용자 정보 조회
	userData, err := h.client.User.Query().Where(user.UserId(userID)).First(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "사용자 정보 조회 실패"})
		return
	}

	// 새 토큰 쌍 생성
	newAccessToken, newRefreshToken, err := utils.GenerateTokenPair(userID, userData.UserName, userData.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "새 토큰 생성 오류"})
		return
	}

	// Redis 업데이트
	if err := database.SaveTokenMetaData(ctx, userID, newAccessToken, utils.AccessTokenDuration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "새 토큰 저장 시 오류"})
		return
	}

	if err := database.SaveTokenMetaData(ctx, userID, newRefreshToken, utils.RefreshTokenDuration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "새 토큰 저장 시 오류"})
		return
	}

	// 이전 토큰 삭제
	database.DeleteTokenMetaData(ctx, claims.TokenUUID)

	// 새 토큰 쿠키에 설정
	c.SetCookie("access_token", newAccessToken, int(utils.AccessTokenDuration.Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", newRefreshToken, int(utils.RefreshTokenDuration.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "토큰 자동 갱신 완료"})
}
