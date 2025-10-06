// cmd/server/main.go
package main

import (
	"log"
	"os"

	"portal-budaya/internal/config"
	"portal-budaya/internal/database"
	"portal-budaya/internal/models"
	"portal-budaya/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Only load .env when not in production to avoid clobbering Railway env
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}
	cfg := config.Load()

	var db *database.DB
	if cfg.DatabaseURL != "" {
		log.Println("Using DATABASE_URL connection mode")
		db = database.ConnectWithURL(cfg.DatabaseURL)
	} else {
		log.Println("Using PG* variables connection mode")
		db = database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBSSLMode)
	}
	if db == nil {
		log.Fatal("Database connection not initialized")
	}

	// Try to enable extension (harmless if already exists)
	if err := db.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Println("warn: unable to ensure uuid-ossp extension:", err)
	}

	if cfg.AppEnv != "production" {
		if err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Tribe{}, &models.Region{}, &models.About{}, &models.Content{}); err != nil {
			log.Fatal("AutoMigrate failed:", err)
		}
	}

	r := gin.Default()
	routes.Register(r, db.DB, []byte(cfg.JWTSecret), cfg.BasicUser, cfg.BasicPass)

	log.Printf("listening on :%s (env=%s)", cfg.AppPort, cfg.AppEnv)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}