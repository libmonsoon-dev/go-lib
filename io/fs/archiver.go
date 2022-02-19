package fs

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/libmonsoon-dev/go-lib/bytes"
)

func newArchiver() *archiver {
	return &archiver{}
}

type archiver struct {
	conf archiveConfig

	zipCompressor *zip.Writer
	tarCompressor *tar.Writer

	buf     []byte
	closers []io.Closer
}

const clearBufferOnReset = false

func (a *archiver) Reset() {
	a.conf = archiveConfig{}

	for i := range a.closers {
		a.closers[i] = nil
	}
	a.closers = a.closers[:0]
	if clearBufferOnReset {
		for i := range a.buf {
			a.buf[i] = 0
		}
	}

	a.zipCompressor = nil
	a.tarCompressor = nil
}

func (a *archiver) ToArchive(fileSystem fs.FS, root string, config archiveConfig, output io.Writer) (err error) {
	a.initBuf()
	a.setConf(config)
	a.initCompressors(output)

	err = fs.WalkDir(fileSystem, root, func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("outer error: %w", err)
		}

		if dirEntry == nil {
			fmt.Println("skipping", fileName)
			return nil
		}

		fileInfo, err := dirEntry.Info()
		if err != nil {
			return fmt.Errorf("get file %s info: %w", fileName, err)
		}

		err = a.addFile(fileSystem, fileName, fileInfo)
		if err != nil {
			return fmt.Errorf("generate header: %w", err)
		}

		return nil
	})

	return
}

func (a *archiver) Close() (err error) {
	for i := range a.closers {
		err = a.closers[i].Close()
		if err != nil {
			return fmt.Errorf("fs.archiver Close[%d] (%T): %w", i, a.closers[i], err)
		}
	}

	return nil
}

func (a *archiver) initBuf() {
	if a.buf == nil {
		if a.conf.BufferSize == 0 {
			a.conf.BufferSize = bytes.MB
		}
		a.buf = make([]byte, a.conf.BufferSize)
	}
}

func (a *archiver) initCompressors(output io.Writer) {
	switch a.conf.ArchiveType {
	case Zip:
		a.zipCompressor = zip.NewWriter(output)

		a.closers = append(a.closers, a.zipCompressor)
	case TarGz:
		gzipWriter := gzip.NewWriter(output)
		a.tarCompressor = tar.NewWriter(gzipWriter)

		a.closers = append(a.closers, gzipWriter, a.tarCompressor)
	default:
		panic("invalid archive type")
	}
}

func (a *archiver) setConf(config archiveConfig) {
	a.conf = config
}

func (a *archiver) addFile(fileSystem fs.FS, fileName string, fileInfo fs.FileInfo) (err error) {
	fileName = filepath.ToSlash(fileName)
	var output io.Writer

	switch a.conf.ArchiveType {
	case Zip:
		var zipHeader *zip.FileHeader
		zipHeader, err = zip.FileInfoHeader(fileInfo)
		if err != nil {
			return fmt.Errorf("generate zip header: %w", err)
		}
		zipHeader.Name = fileName
		zipHeader.Method = a.conf.ZipMethod

		output, err = a.zipCompressor.CreateHeader(zipHeader)
		if err != nil {
			return fmt.Errorf("write zip header: %w", err)
		}
	case TarGz:
		var tarHeader *tar.Header
		tarHeader, err = tar.FileInfoHeader(fileInfo, fileName)
		if err != nil {
			return fmt.Errorf("generate tar header: %w", err)
		}
		tarHeader.Name = fileName

		if err := a.tarCompressor.WriteHeader(tarHeader); err != nil {
			return fmt.Errorf("write tar header: %w", err)
		}
		output = a.tarCompressor
	}

	// if not a dir, write file content
	if fileInfo.IsDir() {
		return nil
	}

	source, err := fileSystem.Open(fileName)
	if err != nil {
		return fmt.Errorf("open %q: %w", fileName, err)
	}
	defer source.Close()

	_, err = io.CopyBuffer(output, source, a.buf)
	if err != nil {
		return fmt.Errorf("copy from file %q: %w", fileName, err)
	}

	return nil
}
