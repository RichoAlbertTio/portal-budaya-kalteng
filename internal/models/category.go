// internal/models/category.go
package models

import "time"

type Category struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;size:80" json:"name"`
	Slug        string    `gorm:"uniqueIndex;size:100" json:"slug"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Contents    []Content `json:"-"`
}