package router

import (
	"net/http"

	"DBP/ent"

	"github.com/gin-gonic/gin"
)

func SetupRouter(client *ent.Client) *gin.Engine {
	router := gin.Default()

	// CORS 설정 (추가적인 http  헤더를 사용하며 다른 출처의 선택한 자원에 접근할 수 있도록 브라우저에게 알려줌)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(204)
		}
		c.Next()
	})

	// 정적 파일 제공 설정
	router.Static("/assets", "../view/assets")
	router.Static("/node_modules", "../view/node_modules")

	router.LoadHTMLGlob("../view/shared/*")

	// 메인 페이지 라우트
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	// // 회원가입 페이지 라우트
	// router.GET("/join", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "main/join.html", nil)
	// })

	// // 회원가입 정보 처리 페이지 라우트
	// router.POST("/join", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "main/join.html", nil)
	// })

	// API 라우트
	//authHandler := handler.NewAuthHandler(client)

	//api := router.Group("/api")
	//{
	//	auth := api.Group("/auth")
	//	{
	//		auth.POST("/check-id", authHandler.CheckDuplicateID)
	//	}
	//}
	return router
}
