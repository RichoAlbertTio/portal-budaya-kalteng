package models

import "time"


type User struct {
ID string `gorm:"type:uuid;primaryKey" json:"id"`
Username string `gorm:"uniqueIndex;size:50" json:"username"`
Email string `gorm:"uniqueIndex;size:120" json:"email"`
DisplayName string `gorm:"size:120" json:"display_name"`
PasswordHash string `json:"-"`
Role string `gorm:"size:20;default:member" json:"role"`
Bio *string `json:"bio"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`
Contents []Content `gorm:"foreignKey:AuthorID" json:"-"`
}