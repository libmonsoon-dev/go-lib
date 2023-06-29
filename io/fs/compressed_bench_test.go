package fs_test

import (
	"embed"
	"io"
	"testing"

	"github.com/libmonsoon-dev/go-lib/io/fs"
)

//go:embed testdata
var testFs embed.FS

func BenchmarkToZipArchive(b *testing.B) {
	b.ReportAllocs()
	var err error

	for i := 0; i < b.N; i++ {
		err = fs.ToZip(testFs, io.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkToTarArchive(b *testing.B) {
	b.ReportAllocs()
	var err error

	for i := 0; i < b.N; i++ {
		err = fs.ToTar(testFs, io.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}
