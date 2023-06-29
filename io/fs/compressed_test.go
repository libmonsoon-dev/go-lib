package fs_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/libmonsoon-dev/go-lib/io/fs"
)

func TestTar(t *testing.T) {
	t.Skip("TODO")
	var buf bytes.Buffer

	err := fs.ToTar(testFs, &buf)
	if err != nil {
		t.Fatal("to tar:", err)
	}

	restored, err := fs.FromTarGzSlice(buf.Bytes())
	if err != nil {
		t.Fatal("from tar:", err)
	}

	if !fsEqual(testFs, restored) {
		t.Fatal("fs not equal")
	}
}

func TestZip(t *testing.T) {
	t.Skip("TODO")
	var buf bytes.Buffer

	err := fs.ToZip(testFs, &buf)
	if err != nil {
		t.Fatal("to zip:", err)
	}

	restored, err := fs.FromZipSlice(buf.Bytes())
	if err != nil {
		t.Fatal("from zip:", err)
	}

	if !fsEqual(testFs, restored) {
		t.Fatal("fs not equal")
	}
}

func fsEqual(a, b fs.FS) bool {
	aContent := make(map[string][]byte)
	bContent := make(map[string][]byte)

	err := readAll(a, aContent)
	if err != nil {
		panic(err)
	}

	err = readAll(b, bContent)
	if err != nil {
		panic(err)
	}

	return reflect.DeepEqual(aContent, bContent)
}

func readAll(fsys fs.FS, output map[string][]byte) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d == nil || d.IsDir() {
			return err
		}

		output[path], err = fs.ReadFile(fsys, path)

		return err
	})

}
