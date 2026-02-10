package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	instance *sqlx.DB
	once     sync.Once
)

func initConnection() {
	user := getEnv("DB_USER")
	pass := getEnv("DB_PASS")
	host := getEnv("DB_HOST")
	name := getEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		user, pass, host, name,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Hour)

	instance = db
}

func GetConnection() *sqlx.DB {
	once.Do(initConnection)
	return instance
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}
