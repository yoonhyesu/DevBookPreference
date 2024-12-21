package database

import (
	"DBP/config"
	"DBP/ent"
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMariaDBConnection() *ent.Client {
	config.LoadEnv()

	username := os.Getenv("MARIADB_USERNAME")
	password := os.Getenv("MARIADB_PASSWORD")
	host := os.Getenv("MARIADB_HOSTNAME")
	port := os.Getenv("MARIADB_PORT")
	dbname := os.Getenv("MARIADB_DBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", username, password, host, port, dbname)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// 스키마 자동 마이그레이션
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
