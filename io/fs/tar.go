package fs

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"io/fs"
)

type Tar struct {
	*tar.Reader
	reseter
}

func (t *Tar) Open(path string) (fs.File, error) {
	t.reseter.Reset()

	for {
		header, err := t.Next()
		if errors.Is(err, io.EOF) {
			return nil, &fs.PathError{Op: "OpenFile", Path: path, Err: fs.ErrNotExist}
		}

		if err != nil {
			return nil, fmt.Errorf("tar reader next: %w", err)
		}

		if header.Name == path {
			return newFile(header, t), nil
		}
	}

}

func newFile(header *tar.Header, r io.Reader) *tarFile {
	r = io.LimitReader(r, header.Size)
	return &tarFile{header, r}
}

type tarFile struct {
	header  *tar.Header
	archive io.Reader
}

func (f *tarFile) Stat() (fs.FileInfo, error) {
	return f.header.FileInfo(), nil
}

func (f *tarFile) Read(data []byte) (int, error) {
	return f.archive.Read(data)
}

func (f *tarFile) Close() error {
	return nil
}
