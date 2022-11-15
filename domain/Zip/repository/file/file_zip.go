package mysql

import (
	"context"

	"restful-file-zip/entities"
)

type fileZipRepository struct {
	SrcPath string
	ZipPath string
}

func NewFileZipRepository(SrcPath, ZipPath string) entities.ZipRepository {
	return &fileZipRepository{SrcPath, ZipPath}
}

func (r *fileZipRepository) Create(ctx context.Context) (path string, err error) {
	return "", nil
}

func (r *fileZipRepository) Clear(ctx context.Context) (err error) {
	return nil
}
