// internal/models/tribe.go
package models

type Tribe struct {
	ID          string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string  `gorm:"size:100" json:"name"`
	Slug        string  `gorm:"uniqueIndex;size:120" json:"slug"`
	Description *string `json:"description,omitempty"`
}