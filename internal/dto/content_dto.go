// internal/dto/content_dto.go
package dto


type ContentCreate struct {
Title string `json:"title" binding:"required"`
ImageURL *string `json:"image_url"`
Summary *string `json:"summary"`
Body string `json:"body" binding:"required"`
CategoryID *string `json:"category_id"`
TribeIDs []string `json:"tribe_ids"`
RegionIDs []string `json:"region_ids"`
Status string `json:"status"`
}