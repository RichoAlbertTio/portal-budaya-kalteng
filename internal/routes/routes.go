// internal/routes/routes.go
package routes

import (
	"time"

	"portal-budaya/internal/handlers"
	"portal-budaya/internal/middlware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func Register(r *gin.Engine, db *gorm.DB, jwtSecret []byte, basicUser, basicPass string) {
	// Handlers
	auth := &handlers.AuthHandler{DB: db, JWTSecret: jwtSecret, JWTTTL: 24 * time.Hour}
	content := &handlers.ContentHandler{DB: db}
	cat := &handlers.CategoryHandler{DB: db}
	tax := &handlers.TaxHandler{DB: db}
	about := &handlers.AboutHandler{DB: db}

	// Public
	api := r.Group("/api")
	{
		api.POST("/auth/register", auth.Register)
		api.POST("/auth/login", auth.Login) // JWT

		api.GET("/contents", content.List)
		api.GET("/contents/:id", content.Get) // by id atau slug
		api.GET("/categories", cat.List)
		api.GET("/tribes", tax.ListTribe)
		api.GET("/regions", tax.ListRegion)
		api.GET("/about", about.Get)
	}

	// Admin via Basic Auth (sesuai brief)
	admin := r.Group("/api/admin", middlware.BasicAuth(middlware.BasicConfig{Username: basicUser, Password: basicPass}))
	{
		admin.POST("/categories", cat.Create)
		admin.POST("/tribes", tax.CreateTribe)
		admin.POST("/regions", tax.CreateRegion)
		admin.POST("/about", about.Upsert)
		admin.POST("/contents", content.Create)
	}

	// Contoh: area yang butuh JWT (nilai plus)
	jwt := r.Group("/api/me", middlware.JWTAuth(middlware.JWTConfig{Secret: jwtSecret}))
	{
		jwt.GET("/profile", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ok"}) })
	}
}