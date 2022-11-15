package entities

import (
	"context"
)

type ZipUsecase interface {
	Create(ctx context.Context) (path string, err error)
	Clear(ctx context.Context) (err error)
}

type ZipRepository interface {
	Create(ctx context.Context) (path string, err error)
	Clear(ctx context.Context) (err error)
}
