package entities

import (
	"context"
	"mime/multipart"
)

type FileObject struct {
	Name string `json:"name"`
}

type FileUsecase interface {
	Save(ctx context.Context, file *multipart.FileHeader) (err error)
	Delete(ctx context.Context, name string) (err error)
	Clear(ctx context.Context) (err error)
	GetList(ctx context.Context) (list []FileObject, err error)
}

type FileRepository interface {
	Save(ctx context.Context, file *multipart.FileHeader) (err error)
	Delete(ctx context.Context, name string) (err error)
	Clear(ctx context.Context) (err error)
	GetList(ctx context.Context) (list []FileObject, err error)
}
