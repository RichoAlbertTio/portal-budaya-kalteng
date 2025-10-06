package config

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppPort      string
	BasicUser    string
	BasicPass    string
	JWTSecret    string
	JWTTTLHours  int
	DBHost       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBName       string
	DBSSLMode    string
	DatabaseURL  string
}

func Load() Config {
	// Check if DATABASE_URL is provided (Railway style)
	databaseURL := os.Getenv("DATABASE_URL")
	
	// Get port from PORT (Railway) or PGPORT (Railway) or APP_PORT (fallback)
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = os.Getenv("PGPORT")
	}
	if appPort == "" {
		appPort = os.Getenv("APP_PORT")
	}
	if appPort == "" {
		appPort = "8090" // default fallback
	}
	
	if databaseURL != "" {
		// Parse DATABASE_URL
		dbConfig := parseDatabaseURL(databaseURL)
		
		// Still validate other required variables
		required := []string{"BASIC_AUTH_USER", "BASIC_AUTH_PASS", "JWT_SECRET", "JWT_TTL_HOURS"}
		for _, key := range required {
			if os.Getenv(key) == "" {
				log.Fatalf("Required environment variable %s is not set", key)
			}
		}
		
		ttl, err := strconv.Atoi(os.Getenv("JWT_TTL_HOURS"))
		if err != nil {
			log.Fatalf("Invalid JWT_TTL_HOURS value: %s", os.Getenv("JWT_TTL_HOURS"))
		}
		
		return Config{
			AppPort:      appPort,
			BasicUser:    os.Getenv("BASIC_AUTH_USER"),
			BasicPass:    os.Getenv("BASIC_AUTH_PASS"),
			JWTSecret:    os.Getenv("JWT_SECRET"),
			JWTTTLHours:  ttl,
			DBHost:       dbConfig.Host,
			DBPort:       dbConfig.Port,
			DBUser:       dbConfig.User,
			DBPass:       dbConfig.Pass,
			DBName:       dbConfig.Name,
			DBSSLMode:    dbConfig.SSLMode,
			DatabaseURL:  databaseURL,
		}
	}
	
	// Fallback to individual DB environment variables
	required := []string{
		"BASIC_AUTH_USER", "BASIC_AUTH_PASS", "JWT_SECRET",
		"JWT_TTL_HOURS", "DB_HOST", "DB_PORT", "DB_USER", "DB_PASS",
		"DB_NAME", "DB_SSLMODE",
	}
	
	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Required environment variable %s is not set", key)
		}
	}

	ttl, err := strconv.Atoi(os.Getenv("JWT_TTL_HOURS"))
	if err != nil {
		log.Fatalf("Invalid JWT_TTL_HOURS value: %s", os.Getenv("JWT_TTL_HOURS"))
	}

	return Config{
		AppPort:      appPort,
		BasicUser:    os.Getenv("BASIC_AUTH_USER"),
		BasicPass:    os.Getenv("BASIC_AUTH_PASS"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		JWTTTLHours:  ttl,
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPass:       os.Getenv("DB_PASS"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    os.Getenv("DB_SSLMODE"),
		DatabaseURL:  "",
	}
}

type dbConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	SSLMode string
}

func parseDatabaseURL(databaseURL string) dbConfig {
	// Parse postgresql://username:password@host:port/database
	u, err := url.Parse(databaseURL)
	if err != nil {
		log.Fatalf("Invalid DATABASE_URL format: %v", err)
	}
	
	password, _ := u.User.Password()
	host := u.Hostname()
	port := u.Port()
	if port == "" {
		port = "5432"
	}
	dbname := strings.TrimPrefix(u.Path, "/")
	
	// SSL mode - check from .env DB_SSLMODE or use Railway default
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		// Check URL query params
		if query := u.Query().Get("sslmode"); query != "" {
			sslMode = query
		} else {
			// Default for Railway - but check if it's internal (no SSL) or external (SSL)
			if strings.Contains(host, "railway.internal") {
				sslMode = "disable"
			} else {
				sslMode = "require"
			}
		}
	}
	
	return dbConfig{
		Host:    host,
		Port:    port,
		User:    u.User.Username(),
		Pass:    password,
		Name:    dbname,
		SSLMode: sslMode,
	}
}