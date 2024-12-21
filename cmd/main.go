package main

import (
	"DBP/config"
	"DBP/internal/router"
	"DBP/pkg/database"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal("환경변수 로드 실패:", err)
	}

	// MariaDB 연결
	client := database.NewMariaDBConnection()
	defer client.Close()

	// Redis 연결
	database.InitRedis()
	defer database.RedisClient.Close()

	// Redis 연결 확인
	ctx := context.Background()
	if err := database.RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis 연결 실패:", err)
	}
	log.Println("Redis 연결 성공!!")

	// 라우터 설정
	router := router.SetupRouter(client)

	// 서버 설정
	srv := &http.Server{
		Addr:    ":7031",
		Handler: router,
	}

	// 고루틴으로 서버 시작
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("서버 시작 실패: %v", err)
		}
	}()

	// 종료 시그널 처리
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("서버 종료 중...")

	// 일정 시간 동안 대기 후 종료
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("서버 강제 종료:", err)
	}

	log.Println("서버가 정상적으로 종료되었습니다!!")
}
