package entities

import (
	"context"
	"mime/multipart"
	"time"
)

type FileObject struct {
	Name    string    `json:"name" example:"123.png" extensions:"x-order=1"`
	Size    int64     `json:"size" example:"456" extensions:"x-order=2"`
	ModTime time.Time `json:"modeTime" example:"2022-11-15T15:37:55.1779141+08:00" extensions:"x-order=3"`
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
