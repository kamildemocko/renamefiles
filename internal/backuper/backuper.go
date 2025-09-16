package backuper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type Backuper struct {
	Filename string
	file     *os.File
	writer   *zip.Writer
}

func NewBackuper(dir string, filename string) (Backuper, error) {
	fn := filepath.Join(dir, filename)
	file, err := os.Create(fn)
	if err != nil {
		return Backuper{}, nil
	}

	w := zip.NewWriter(file)

	return Backuper{filename, file, w}, err
}

func (b *Backuper) Close() {
	if b.writer != nil {
		b.writer.Close()
	}

	if b.file != nil {
		b.file.Close()
	}
}

func (b *Backuper) AddFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	baseName := filepath.Base(path)
	f, err := b.writer.Create(baseName)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil
}
