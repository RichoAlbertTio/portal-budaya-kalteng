package config

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppEnv      string
	AppPort     string
	BasicUser   string
	BasicPass   string
	JWTSecret   string
	JWTTTLHours int
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	DBSSLMode   string
	DatabaseURL string
}

func Load() Config {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" { appEnv = "development" }

	databaseURL := os.Getenv("DATABASE_URL")

	// Use PORT from env; fallback to 8080 if not set (so container still responds)
	appPort := os.Getenv("PORT")
	if appPort == "" { appPort = "8080" }

	if databaseURL != "" {
		dbConfig := parseDatabaseURL(databaseURL)
		required := []string{"BASIC_AUTH_USER", "BASIC_AUTH_PASS", "JWT_SECRET", "JWT_TTL_HOURS"}
		for _, key := range required { if os.Getenv(key) == "" { log.Fatalf("Required environment variable %s is not set", key) } }
		ttl, err := strconv.Atoi(os.Getenv("JWT_TTL_HOURS")); if err != nil { log.Fatalf("Invalid JWT_TTL_HOURS value: %s", os.Getenv("JWT_TTL_HOURS")) }
		return Config{AppEnv: appEnv, AppPort: appPort, BasicUser: os.Getenv("BASIC_AUTH_USER"), BasicPass: os.Getenv("BASIC_AUTH_PASS"), JWTSecret: os.Getenv("JWT_SECRET"), JWTTTLHours: ttl, DBHost: dbConfig.Host, DBPort: dbConfig.Port, DBUser: dbConfig.User, DBPass: dbConfig.Pass, DBName: dbConfig.Name, DBSSLMode: dbConfig.SSLMode, DatabaseURL: databaseURL}
	}

	// Fallback path (only if DATABASE_URL absent) uses PG* naming from env
	required := []string{"BASIC_AUTH_USER", "BASIC_AUTH_PASS", "JWT_SECRET", "JWT_TTL_HOURS", "PGHOST", "PGPORT", "PGUSER", "PGPASSWORD", "PGDATABASE", "DB_SSLMODE"}
	for _, key := range required { if os.Getenv(key) == "" { log.Fatalf("Required environment variable %s is not set", key) } }
	ttl, err := strconv.Atoi(os.Getenv("JWT_TTL_HOURS")); if err != nil { log.Fatalf("Invalid JWT_TTL_HOURS value: %s", os.Getenv("JWT_TTL_HOURS")) }
	return Config{AppEnv: appEnv, AppPort: appPort, BasicUser: os.Getenv("BASIC_AUTH_USER"), BasicPass: os.Getenv("BASIC_AUTH_PASS"), JWTSecret: os.Getenv("JWT_SECRET"), JWTTTLHours: ttl, DBHost: os.Getenv("PGHOST"), DBPort: os.Getenv("PGPORT"), DBUser: os.Getenv("PGUSER"), DBPass: os.Getenv("PGPASSWORD"), DBName: os.Getenv("PGDATABASE"), DBSSLMode: os.Getenv("DB_SSLMODE"), DatabaseURL: ""}
}

type dbConfig struct { Host, Port, User, Pass, Name, SSLMode string }

func parseDatabaseURL(databaseURL string) dbConfig {
	u, err := url.Parse(databaseURL)
	if err != nil { log.Fatalf("Invalid DATABASE_URL format: %v", err) }
	password, _ := u.User.Password()
	host := u.Hostname()
	port := u.Port(); if port == "" { port = "5432" }
	dbname := strings.TrimPrefix(u.Path, "/")
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		if q := u.Query().Get("sslmode"); q != "" { sslMode = q } else {
			if strings.Contains(host, "railway.internal") { sslMode = "disable" } else { sslMode = "require" }
		}
	}
	return dbConfig{Host: host, Port: port, User: u.User.Username(), Pass: password, Name: dbname, SSLMode: sslMode}
}