package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// 현재 실행 파일의 경로를 기준으로 상위 디렉토리로 이동
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("현재 디렉토리 확인 실패: %v", err)
		return err
	}

	// cmd 폴더에서 실행되는 경우 상위 디렉토리로 이동
	if filepath.Base(currentDir) == "cmd" {
		currentDir = filepath.Dir(currentDir)
	}

	envPath := filepath.Join(currentDir, ".env")
	log.Printf("환경변수 파일 경로: %s", envPath) // 디버깅을 위한 로그 추가

	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("환경변수 로드 실패 (%s): %v", envPath, err)
		return err
	}

	log.Println("환경변수 로드 성공!!!")
	return nil
}
