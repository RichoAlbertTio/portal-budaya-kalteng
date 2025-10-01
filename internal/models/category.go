// internal/models/category.go
package models


import "time"


type Category struct {
ID string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
Name string `gorm:"uniqueIndex;size:80"`
Slug string `gorm:"uniqueIndex;size:100"`
Description *string
CreatedAt time.Time
UpdatedAt time.Time
Contents []Content `json:"-"`
}