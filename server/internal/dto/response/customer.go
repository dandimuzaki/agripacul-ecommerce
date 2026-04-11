package response

import (
	"debian-ecommerce/internal/data/entity"
	"time"
)

type CustomerProfile struct {
	ID              uint      `json:"id"`
	FullName            string    `json:"full_name"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone_number"`
	ProfileImageURL string    `json:"profile_image_url"`
	DateOfBirth     *time.Time `json:"date_of_birth"`
}

func ToCustomerProfile(cust *entity.Customer) *CustomerProfile {
	return &CustomerProfile{
		ID: cust.ID,
		FullName: cust.FullName,
		Email: cust.User.Email,
		PhoneNumber: cust.PhoneNumber,
		ProfileImageURL: cust.ProfileImageURL,
		DateOfBirth: cust.DateOfBirth,
	}
}