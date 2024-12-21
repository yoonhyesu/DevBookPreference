package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// 사용자 정보를 담은 클레임 생성
type TokenClaims struct {
	UserID               string `json:"USER_ID"`
	UserName             string `json:"USER_NAME"`
	IsAdmin              bool   `json:"IS_ADMIN"`
	TokenUUID            string `json:"TOKEN_UUID"`
	jwt.RegisteredClaims        // jwt의 기본클레임 타입가져옴
}

const (
	AccessTokenDuration  = time.Hour
	RefreshTokenDuration = time.Hour * 24 * 7
)

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GenerateTokenPair(userID, userName string, isAdmin bool) (string, string, error) {
	// Access Token 발행
	accessUUID := uuid.New().String()
	accessClaims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:    userID,
		UserName:  userName,
		IsAdmin:   isAdmin,
		TokenUUID: accessUUID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(getSecretKey())
	if err != nil {
		log.Printf("액세스 토큰 생성 실패: %v", err)
		return "", "", err
	}
	log.Printf("액세스 토큰 생성 성공: %v", accessTokenString)

	// Refresh Token 발행
	refreshUUID := uuid.New().String()
	refreshClaims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:    userID,
		TokenUUID: refreshUUID,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(getSecretKey())
	if err != nil {
		log.Printf("리프레시 토큰 생성 실패: %v", err)
		return "", "", err
	}
	log.Printf("리프레시 토큰 생성 성공: %v", refreshTokenString)

	return accessTokenString, refreshTokenString, nil
}

// 유효성검사 어렵다
func ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 서명 방식 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("잘못된 서명 방식: %v", token.Header["alg"])
		}
		return getSecretKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("유효하지 않은 토큰")
}

// 토큰 만료 시간 확인 함수 추가
func IsTokenExpiration(tokenString string) bool {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return false
	}
	// 만료 5분 전인지 확인
	expirationTime := claims.ExpiresAt.Time
	timeUntilExpiration := time.Until(expirationTime)
	return timeUntilExpiration <= 5*time.Minute
}
