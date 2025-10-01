// internal/models/tribe.go
package models


type Tribe struct {
ID string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
Name string `gorm:"size:100"`
Slug string `gorm:"uniqueIndex;size:120"`
Description *string
}