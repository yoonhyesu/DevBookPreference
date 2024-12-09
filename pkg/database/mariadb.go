package database

import (
	"SpaceDev/ent"
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMariaDBConnection() *ent.Client {
	client, err := ent.Open("mysql", "root:0413@tcp(localhost:3306)/DBP?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// 스키마 자동 마이그레이션
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
