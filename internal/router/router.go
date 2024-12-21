package router

import (
	"net/http"

	"DBP/ent"
	"DBP/internal/handler"
	"DBP/internal/middleware"
	"DBP/internal/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(client *ent.Client) *gin.Engine {
	router := gin.Default() // Default함수는 Engine인스턴스 리턴함. 반환된 인스턴스는 Logger와 Recovery 미들웨어 부착상태임

	// 디버깅을 위한 로그 미들웨어 추가
	router.Use(gin.Logger())

	// CORS 설정
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 정적 파일 제공 설정 (URL경로, 실제 파일시스템 경로)
	router.Static("/node_modules", "../view/node_modules")
	router.Static("/assets", "../view/assets")
	router.StaticFile("/favicon.ico", "../favicon.ico")

	// 템플릿 경로 설정 (여러경로를 호출하면 두번째 호출이 첫번째호출의 템플릿을 덮어씀)
	router.LoadHTMLGlob("../view/**/*.html")

	// 페이지 라우트
	// 메인 페이지 라우트
	router.GET("/", func(c *gin.Context) {
		// 쿠키에서 토큰 가져오기
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			// 토큰이 없으면 비로그인 상태로 렌더링
			c.HTML(http.StatusOK, "index.html", gin.H{
				"user": nil,
			})
			return
		}

		// 토큰 검증
		claims, err := utils.ValidateToken(accessToken)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"user": nil,
			})
			return
		}

		// 사용자 정보를 템플릿에 전달
		c.HTML(http.StatusOK, "index.html", gin.H{
			"user": gin.H{
				"UserName": claims.UserName,
				"IsAdmin":  claims.IsAdmin,
			},
		})
	})

	// 회원가입 페이지 라우트 (경로, 핸들러함수)
	router.GET("/join", func(c *gin.Context) {
		c.HTML(http.StatusOK, "join.html", nil)
	})

	// 로그인 페이지 라우트
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// auth 라우트 그룹 (로그인, 로그아웃, 토큰 갱신, 아이디 중복 체크)
	authHandler := handler.NewAuthHandler(client)
	auth := router.Group("/auth")
	{
		auth.POST("/check-duplicate", authHandler.CheckDuplicateID)
		auth.POST("/signin", authHandler.SignIn)
		auth.POST("/signout", middleware.AuthMiddleware(), authHandler.SignOut)
		auth.GET("/refresh-token", authHandler.RefreshToken)
	}

	// 유저 라우트 그룹 (회원가입, 마이페이지)
	userHandler := handler.NewUserHandler(client)
	user := router.Group("/user")
	user.Use(middleware.TransactionMiddleware(client))
	{
		user.POST("/register", userHandler.Register)
	}

	// 관리자 전용 라우트
	protected := router.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		admin := protected.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.GET("/board", func(c *gin.Context) {
				c.HTML(http.StatusOK, "board.html", nil)
			})
		}
	}

	return router
}
