// internal/models/region.go
package models

type Region struct {
	ID          string  `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string  `gorm:"size:100" json:"name"`
	Slug        string  `gorm:"uniqueIndex;size:120" json:"slug"`
	Description *string `json:"description,omitempty"`
}