// cmd/server/main.go
package main


import (
"log"
"os"
"time"
"github.com/gin-gonic/gin"
"github.com/joho/godotenv"
"portal-budaya/internal/config"
"portal-budaya/internal/database"
"portal-budaya/internal/models"
"portal-budaya/internal/routes"
)


func main(){ _ = godotenv.Load()
cfg := config.Load()


db := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBSSLMode)
// auto-migrate (aman untuk development)
db.AutoMigrate(&models.User{}, &models.Category{}, &models.Tribe{}, &models.Region{}, &models.About{}, &models.Content{})


r := gin.Default()
routes.Register(r, db.DB, []byte(cfg.JWTSecret), cfg.BasicUser, cfg.BasicPass)


log.Println("listening on :"+cfg.AppPort)
if err := r.Run(":"+cfg.AppPort); err != nil { log.Fatal(err) }
}