package fs

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
)

// TODO: tests
func ToZip(fileSystem fs.FS, output io.Writer) error {
	conf := defaultZipConfig()
	return toArchive(fileSystem, ".", conf, output)
}

func ToTar(fileSystem fs.FS, output io.Writer) error {
	conf := defaultTarConfig()
	return toArchive(fileSystem, ".", conf, output)
}

func FromZipSlice(data []byte) (*zip.Reader, error) {
	r := bytes.NewReader(data)
	return zip.NewReader(r, r.Size())
}

func FromTarGzSlice(data []byte) (fileSystem *Tar, err error) {
	buf := bytes.NewBuffer(data)
	return NewTarGz(buf)
}

type reseter interface {
	Reset()
}

type readerReseter interface {
	io.Reader
	reseter
}

func NewTarGz(input readerReseter) (*Tar, error) {
	gzReader, err := gzip.NewReader(input)
	if err != nil {
		return nil, fmt.Errorf("new gzip reader: %w", err)
	}

	tarReader := tar.NewReader(gzReader)
	return &Tar{tarReader, input}, nil
}

func toArchive(fileSystem fs.FS, root string, config archiveConfig, output io.Writer) error {
	a := acquireArchiver()
	defer releaseArchiver(a)

	return a.ToArchive(fileSystem, root, config, output)
}
