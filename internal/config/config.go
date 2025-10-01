package config

import (
	"os"
	"strconv"
)


type Config struct {
AppPort string
BasicUser string
BasicPass string
JWTSecret string
JWTTTLHours int
DBHost string
DBPort string
DBUser string
DBPass string
DBName string
DBSSLMode string
}


func Load() Config {
ttl, _ := strconv.Atoi(get("JWT_TTL_HOURS", "24"))
return Config{
AppPort: get("APP_PORT", "8080"),
BasicUser: get("BASIC_AUTH_USER", "admin"),
BasicPass: get("BASIC_AUTH_PASS", "admin"),
JWTSecret: get("JWT_SECRET", "secret"),
JWTTTLHours: ttl,
DBHost: get("DB_HOST", "localhost"),
DBPort: get("DB_PORT", "5432"),
DBUser: get("DB_USER", "postgres"),
DBPass: get("DB_PASS", "postgres"),
DBName: get("DB_NAME", "portal_budaya"),
DBSSLMode: get("DB_SSLMODE", "disable"),
}
}


func get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}