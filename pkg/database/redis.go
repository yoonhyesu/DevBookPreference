package database

import (
	"DBP/config"
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	config.LoadEnv()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})
}

func SaveTokenMetaData(ctx context.Context, userID, tokenUUID string, exp time.Duration) error {
	// Redis 연결 상태 확인
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Printf("Redis 연결 확인 실패: %v", err)
		return err
	}

	// 토큰 저장
	err := RedisClient.Set(ctx, tokenUUID, userID, exp).Err()
	if err != nil {
		log.Printf("Redis 토큰 저장 실패 - TokenUUID: %s, UserID: %s, Error: %v", tokenUUID, userID, err)
		return err
	}

	// 저장 확인
	storedValue, err := RedisClient.Get(ctx, tokenUUID).Result()
	if err != nil {
		log.Printf("Redis 저장 확인 실패 - TokenUUID: %s, Error: %v", tokenUUID, err)
		return err
	}
	log.Printf("Redis 토큰 저장 성공 - TokenUUID: %s, StoredValue: %s", tokenUUID, storedValue)

	return nil
}

func DeleteTokenMetaData(ctx context.Context, tokenUUID string) error {
	return RedisClient.Del(ctx, tokenUUID).Err()
}
