package entities

import (
	"context"
)

type ZipUsecase interface {
	Get(ctx context.Context) (name, path string, err error)
}

type ZipRepository interface {
	Prepare(ctx context.Context) (err error)
	CreateZip(ctx context.Context, name string) (path string, err error)
}
