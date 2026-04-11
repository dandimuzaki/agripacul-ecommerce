package uploader

import "github.com/cloudinary/cloudinary-go/v2"

func NewClient() (*cloudinary.Cloudinary, error) {
	return cloudinary.New()
}
