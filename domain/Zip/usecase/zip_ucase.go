package usecase

import (
	"context"
	"time"

	"restful-file-zip/entities"
)

type zipUsecase struct {
	repo           entities.ZipRepository
	contextTimeout time.Duration
}

func NewZipUsecase(r entities.ZipRepository, timeout time.Duration) entities.ZipUsecase {
	return &zipUsecase{
		repo:           r,
		contextTimeout: timeout,
	}
}

func (r *zipUsecase) Create(c context.Context) (path string, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	path, err = r.repo.Create(ctx)

	return
}

func (r *zipUsecase) Clear(c context.Context) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	err = r.repo.Clear(ctx)

	return
}
