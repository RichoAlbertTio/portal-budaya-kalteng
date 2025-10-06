// internal/models/content.go
package models

import "time"


type Content struct {
ID string `gorm:"type:uuid;primaryKey" json:"id"`
Title string `gorm:"size:180" json:"title"`
Slug string `gorm:"uniqueIndex;size:200" json:"slug"`
ImageURL *string `json:"image_url"`
Summary *string `json:"summary"`
Body string `gorm:"type:text" json:"body"`
Status string `gorm:"size:20;default:draft" json:"status"`
PublishedAt *time.Time `json:"published_at"`
CategoryID *string `gorm:"type:uuid" json:"category_id"`
Category *Category `json:"category"`
AuthorID *string `gorm:"type:uuid" json:"author_id"`
Author *User `json:"author"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`


Tribes []Tribe `gorm:"many2many:content_tribes;" json:"tribes"`
Regions []Region `gorm:"many2many:content_regions;" json:"regions"`
}