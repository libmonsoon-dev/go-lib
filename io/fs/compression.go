package fs

import (
	"archive/zip"
	"io"
	"io/fs"
)

type ArchiveType int

const (
	Zip ArchiveType = iota
	TarGz
)

type archiveConfig struct {
	ArchiveType
	ZipMethod  uint16
	BufferSize int
}

func ToZip(fileSystem fs.FS, output io.Writer) error {
	conf := defaultZipConfig()
	return toArchive(fileSystem, ".", conf, output)
}

func ToTar(fileSystem fs.FS, output io.Writer) error {
	conf := defaultTarConfig()
	return toArchive(fileSystem, ".", conf, output)
}

func defaultZipConfig() archiveConfig {
	return archiveConfig{ArchiveType: Zip, ZipMethod: zip.Deflate}
}

func defaultTarConfig() archiveConfig {
	return archiveConfig{ArchiveType: TarGz}
}

func toArchive(fileSystem fs.FS, root string, config archiveConfig, output io.Writer) error {
	a := acquireArchiver()
	defer releaseArchiver(a)

	return a.ToArchive(fileSystem, root, config, output)
}
