package mysql

import (
	"archive/zip"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"restful-file-zip/entities"
)

type fileZipRepository struct {
	SrcPath     string
	ArchivePath string
}

func NewFileZipRepository(SrcPath, ArchivePath string) entities.ZipRepository {
	return &fileZipRepository{SrcPath, ArchivePath}
}

func (r *fileZipRepository) Prepare(ctx context.Context) (err error) {
	err = os.RemoveAll(r.ArchivePath)
	if err != nil {
		return
	}

	err = os.Mkdir(r.ArchivePath, 0750)
	return err
}

func (r *fileZipRepository) ValidateSrc(ctx context.Context) (err error) {
	files, err := ioutil.ReadDir(r.SrcPath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New(`source buffer is empty`)
	}

	return nil
}

func (r *fileZipRepository) CreateZip(ctx context.Context, name string) (result string, err error) {
	result = r.ArchivePath + "/" + name

	f, err := os.Create(result)
	if err != nil {
		return "", err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	err = filepath.Walk(r.SrcPath, func(path string, info os.FileInfo, err error) error {
		return fileWalker(r.SrcPath, writer, path, info, err)
	})

	return
}

func fileWalker(src string, writer *zip.Writer, path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// set compression
	header.Method = zip.Deflate

	// Set relative path of a file as the header name
	header.Name, err = filepath.Rel(filepath.Dir(src), path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		header.Name += "/"
	}

	// Create writer for the file header and save content of the file
	headerWriter, err := writer.CreateHeader(header)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(headerWriter, f)
	return err
}
