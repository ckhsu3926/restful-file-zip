package mysql

import (
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"restful-file-zip/entities"
)

type fileRepository struct {
	Path string
}

func NewFileRepository(Path string) entities.FileRepository {
	return &fileRepository{Path}
}

func (r *fileRepository) Save(ctx context.Context, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(r.Path + "/" + file.Filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (r *fileRepository) Delete(ctx context.Context, name string) error {
	filepath := r.Path + "/" + name
	err := os.Remove(filepath)

	return err
}

func (r *fileRepository) Clear(ctx context.Context) (err error) {
	err = os.RemoveAll(r.Path)
	if err != nil {
		return
	}

	err = os.Mkdir(r.Path, 0750)
	return err
}

func (r *fileRepository) GetList(ctx context.Context) (list []entities.FileObject, err error) {
	list = []entities.FileObject{}

	files, err := ioutil.ReadDir(r.Path)
	if err != nil {
		return
	}

	for _, file := range files {
		list = append(list, entities.FileObject{
			Name:    file.Name(),
			Size:    file.Size(),
			ModTime: file.ModTime(),
		})
	}

	return
}
