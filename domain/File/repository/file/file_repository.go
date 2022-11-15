package mysql

import (
	"context"
	"mime/multipart"

	"restful-file-zip/entities"
)

type fileRepository struct {
	Path string
}

func NewFileRepository(Path string) entities.FileRepository {
	return &fileRepository{Path}
}

func (r *fileRepository) Save(ctx context.Context, file *multipart.FileHeader) (err error) {
	return nil
}

func (r *fileRepository) Delete(ctx context.Context, name string) (err error) {
	return nil
}

func (r *fileRepository) Clear(ctx context.Context) (err error) {
	return nil
}

func (r *fileRepository) GetList(ctx context.Context) (list []entities.FileObject, err error) {
	return []entities.FileObject{}, nil
}
