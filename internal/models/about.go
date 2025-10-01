// internal/models/about.go
package models


import "time"


type About struct {
ID string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
Title string `gorm:"size:150"`
Description string `gorm:"type:text"`
UpdatedBy *string `gorm:"type:uuid"`
CreatedAt time.Time
UpdatedAt time.Time
}