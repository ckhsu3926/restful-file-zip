package usecase

import (
	"context"
	"mime/multipart"
	"time"

	"restful-file-zip/entities"
)

type fileUsecase struct {
	repo           entities.FileRepository
	contextTimeout time.Duration
}

func NewFileUsecase(r entities.FileRepository, timeout time.Duration) entities.FileUsecase {
	return &fileUsecase{
		repo:           r,
		contextTimeout: timeout,
	}
}

func (r *fileUsecase) Save(c context.Context, file *multipart.FileHeader) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	err = r.repo.Save(ctx, file)

	return
}

func (r *fileUsecase) Delete(c context.Context, name string) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	err = r.repo.Delete(ctx, name)

	return
}

func (r *fileUsecase) Clear(c context.Context) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	err = r.repo.Clear(ctx)

	return
}

func (r *fileUsecase) GetList(c context.Context) (list []entities.FileObject, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	list, err = r.repo.GetList(ctx)

	return
}
