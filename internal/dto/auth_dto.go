// internal/dto/auth_dto.go
package dto


type RegisterRequest struct {
Username string `json:"username" binding:"required,min=3"`
Email string `json:"email" binding:"required,email"`
DisplayName string `json:"display_name" binding:"required"`
Password string `json:"password" binding:"required,min=6"`
}


type LoginRequest struct {
UsernameOrEmail string `json:"username_or_email" binding:"required"`
Password string `json:"password" binding:"required"`
}