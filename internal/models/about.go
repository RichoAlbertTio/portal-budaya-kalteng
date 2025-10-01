// internal/models/about.go
package models

import "time"

type About struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string    `gorm:"size:150" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	UpdatedBy   *string   `gorm:"type:uuid" json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

