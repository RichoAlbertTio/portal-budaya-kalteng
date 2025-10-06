// cmd/migrate/main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func firstNonEmpty(keys ...string) string {
	for _, k := range keys {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}

func main() {
	// Load .env
	_ = godotenv.Load()

	// Prefer DATABASE_URL (Railway style)
	dsn := os.Getenv("DATABASE_URL")
	usedURL := false
	if dsn != "" {
		usedURL = true

	} else {
		// Build from individual vars (support both DB_* and PG* and the typo PHGHOST)
		host := firstNonEmpty( "PGHOST", "PHGHOST")
		port := firstNonEmpty("PGPORT")
		user := firstNonEmpty("PGUSER")
		pass := firstNonEmpty("PGPASSWORD", "POSTGRES_PASSWORD")
		name := firstNonEmpty("PGDATABASE", "POSTGRES_DB")
		ssl := firstNonEmpty("DB_SSLMODE")
		if ssl == "" {
			ssl = "disable"
		}

		if host == "" || port == "" || user == "" || name == "" {
			log.Fatal("Missing required database environment variables (host/port/user/name)")
		}
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl)
	}

	log.Printf("Running migration using %s", func() string {
		if usedURL {
			return "DATABASE_URL"
		}
		return "assembled DSN"
	}())

	// Log safe parts (mask password)
	if usedURL {
		if u, err := url.Parse(dsn); err == nil {
			pw, _ := u.User.Password()
			masked := strings.Repeat("*", len(pw))
			u.User = url.UserPassword(u.User.Username(), masked)
			log.Printf("Target: %s", u.Redacted())
		}
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to open connection:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	sqlContent, err := os.ReadFile("migrations/001_init.sql")
	if err != nil {
		log.Fatal("Failed to read migration file:", err)
	}

	log.Println("Applying migration 001_init.sql ...")
	if _, err = db.Exec(string(sqlContent)); err != nil {
		log.Fatal("Failed to execute migration:", err)
	}
	log.Println("✓ Migration completed successfully")
}
