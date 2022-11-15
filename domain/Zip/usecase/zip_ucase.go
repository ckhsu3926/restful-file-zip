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

func (r *zipUsecase) Get(c context.Context) (name, path string, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	err = r.repo.Prepare(ctx)
	if err != nil {
		return
	}

	err = r.repo.ValidateSrc(ctx)
	if err != nil {
		return
	}

	name = time.Now().Format("2006-01-02 150405") + ".zip"

	path, err = r.repo.CreateZip(ctx, name)
	return
}
