package usecase

import (
	"context"
	"mime/multipart"
)

type TxManager interface {
	WithinTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type ImageUploader interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (url string, publicID string, err error)
	Delete(ctx context.Context, publicID string) error
}