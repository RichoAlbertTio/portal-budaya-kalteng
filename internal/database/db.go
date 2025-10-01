// internal/database/db.go
package database


import (
"fmt"
"log"
"time"


"gorm.io/driver/postgres"
"gorm.io/gorm"
"gorm.io/gorm/logger"
)


type DB struct { *gorm.DB }


func Connect(host, port, user, pass, name, ssl string) *DB {
dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
host, port, user, pass, name, ssl,
)
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
Logger: logger.Default.LogMode(logger.Warn),
})
if err != nil { log.Fatal("DB connect error:", err) }


// connection pool
sqlDB, _ := db.DB()
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(25)
sqlDB.SetConnMaxLifetime(5 * time.Minute)


return &DB{db}
}