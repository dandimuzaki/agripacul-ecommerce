package uploader

import (
	"context"
	"debian-ecommerce/pkg/utils"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Uploader struct {
	cld *cloudinary.Cloudinary
}

func NewUploader(cfg utils.CloudinaryConfig) (*Uploader, error) {
	cld, err := cloudinary.NewFromParams(
		cfg.CloudName,
		cfg.APIKey,
		cfg.APISecret,
	)
	if err != nil {
		return nil, err
	}

	return &Uploader{cld: cld}, nil
}

func (u *Uploader) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, string, error) {
	resp, err := u.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: folder,
	})
	if err != nil {
		return "", "", err
	}

	return resp.SecureURL, resp.PublicID, nil
}

func (u *Uploader) Delete(
	ctx context.Context,
	publicID string,
) error {

	if publicID == "" {
		// No image to delete, treat as success
		return nil
	}

	res, err := u.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return err
	}

	// Cloudinary may return "not found" but still 200 OK
	if res.Result != "ok" && res.Result != "not found" {
		return fmt.Errorf("failed to delete image: %s", res.Result)
	}

	return nil
}

