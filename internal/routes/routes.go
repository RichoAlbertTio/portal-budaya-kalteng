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

	// Public API - Read Only
	api := r.Group("/api")
	{
		// Authentication
		api.POST("/auth/register", auth.Register)
		api.POST("/auth/login", auth.Login)

		// Contents - Public Read
		api.GET("/contents", content.List)
		api.GET("/contents/:slug", content.Get)

		// Categories - Public Read
		api.GET("/categories", cat.List)
		api.GET("/categories/:slug", cat.Get)

		// Tribes - Public Read
		api.GET("/tribes", tax.ListTribe)
		api.GET("/tribes/:slug", tax.GetTribe)

		// Regions - Public Read
		api.GET("/regions", tax.ListRegion)
		api.GET("/regions/:slug", tax.GetRegion)

		// About - Public Read
		api.GET("/abouts", about.Get)
	}

	// Admin API - Full CRUD (Basic Auth Required)
	admin := r.Group("/api/admin", middlware.BasicAuth(middlware.BasicConfig{Username: basicUser, Password: basicPass}))
	{
		// Categories CRUD
		admin.GET("/categories", cat.List)
		admin.POST("/categories", cat.Create)
		admin.PUT("/categories/:id", cat.Update)
		admin.GET("/categories/:slug", cat.Get)
		admin.DELETE("/categories/:id", cat.Delete)

		// Tribes CRUD
		admin.GET("/tribes", tax.ListTribe)
		admin.POST("/tribes", tax.CreateTribe)
		admin.PUT("/tribes/:id", tax.UpdateTribe)
		admin.GET("/tribes/:slug", tax.GetTribe)
		admin.DELETE("/tribes/:id", tax.DeleteTribe)

		// Regions CRUD
		admin.GET("/regions", tax.ListRegion)
		admin.POST("/regions", tax.CreateRegion)
		admin.PUT("/regions/:id", tax.UpdateRegion)
		admin.GET("/regions/:slug", tax.GetRegion)
		admin.DELETE("/regions/:id", tax.DeleteRegion)

		// Contents CRUD
		admin.GET("/contents", content.List)
		admin.POST("/contents", content.Create)
		admin.PUT("/contents/:id", content.Update)
		admin.GET("/contents/:slug", content.Get)
		admin.DELETE("/contents/:id", content.Delete)

		// About Upsert
		admin.GET("/abouts", about.Get)
		admin.POST("/abouts", about.Upsert)
	}

	// Protected API - JWT Required
	jwt := r.Group("/api/me", middlware.JWTAuth(middlware.JWTConfig{Secret: jwtSecret}))
	{
		jwt.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "ok"})
		})
	}
}