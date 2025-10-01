package config

import (
	"log"
	"os"
	"strconv"
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
}

func Load() Config {
	// Validate required environment variables
	required := []string{
		"APP_PORT", "BASIC_AUTH_USER", "BASIC_AUTH_PASS", "JWT_SECRET",
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
		AppPort:      os.Getenv("APP_PORT"),
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
	}
}