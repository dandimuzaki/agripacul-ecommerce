package request

import (
	"mime/multipart"
)

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateProfileRequest struct {
	FullName     string                `form:"full_name"`
	PhoneNumber        string                `form:"phone_number"`
	DateOfBirth string `form:"date_of_birth"`
	ProfileImage *multipart.FileHeader `form:"profile_image"`
}
