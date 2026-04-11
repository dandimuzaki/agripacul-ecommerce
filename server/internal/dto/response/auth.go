package response

import "debian-ecommerce/internal/data/entity"

type UserResponse struct {
	ID    uint       `json:"id"`
	Email string     `json:"email"`
	Role  entity.UserRole `json:"role"`
}

type AuthResponse struct {
	User UserResponse `json:"user"`
	Token string `json:"token"`
}

type CheckEmailResponse struct {
	IsAvailable bool `json:"is_available"`
}