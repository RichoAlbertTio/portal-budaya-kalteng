// internal/database/db.go
package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct { *gorm.DB }

// Connect builds DSN from explicit params (legacy usage)
func Connect(host, port, user, pass, name, ssl string) *DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl)
	return open(dsn)
}

// ConnectWithURL opens using a full DATABASE_URL
func ConnectWithURL(databaseURL string) *DB { return open(databaseURL) }

// ConnectAuto matches env style you provided:
// 1. Use DATABASE_URL if present.
// 2. Else assemble from PG*/DB_* variables (accept PHGHOST typo).
func ConnectAuto() *DB {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return open(url)
	}
	// Gather pieces from multiple possible names
	host := firstNonEmpty("DB_HOST", "PGHOST", "PHGHOST")
	port := firstNonEmpty("DB_PORT", "PGPORT")
	user := firstNonEmpty("DB_USER", "PGUSER", "POSTGRES_USER")
	pass := firstNonEmpty("DB_PASS", "PGPASSWORD", "POSTGRES_PASSWORD")
	name := firstNonEmpty("DB_NAME", "PGDATABASE", "POSTGRES_DB")
	ssl := firstNonEmpty("DB_SSLMODE", "PGSSLMODE")
	if ssl == "" { ssl = "disable" }

	if host == "" || port == "" || user == "" || name == "" {
		log.Println("ConnectAuto: missing required DB env vars; skipping auto connect")
		return nil
	}
	return open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl))
}

func firstNonEmpty(keys ...string) string {
	for _, k := range keys { if v := os.Getenv(k); v != "" { return v } }
	return ""
}

func open(dsn string) *DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{ Logger: logger.Default.LogMode(logger.Warn) })
	if err != nil { log.Fatal("DB connect error:", err) }
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	return &DB{db}
}