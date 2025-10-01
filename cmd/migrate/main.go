// cmd/migrate/main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	_ = godotenv.Load()

	// Get database config from environment
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	// Build connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSL,
	)

	log.Printf("Connecting to database: %s@%s:%s/%s", dbUser, dbHost, dbPort, dbName)

	// Connect to database
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Read migration file
	sqlContent, err := os.ReadFile("migrations/001_init.sql")
	if err != nil {
		log.Fatal("Failed to read migration file:", err)
	}

	// Execute migration
	log.Println("Running migration...")
	_, err = db.Exec(string(sqlContent))
	if err != nil {
		log.Fatal("Failed to execute migration:", err)
	}

	log.Println("✓ Migration completed successfully!")
}
