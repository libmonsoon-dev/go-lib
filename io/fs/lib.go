package fs

import libfs "io/fs"

type FS = libfs.FS

type WalkDirFunc = libfs.WalkDirFunc

type DirEntry = libfs.DirEntry

func WalkDir(fsys FS, root string, fn WalkDirFunc) error {
	return libfs.WalkDir(fsys, root, fn)
}

func ReadFile(fsys FS, name string) ([]byte, error) {
	return libfs.ReadFile(fsys, name)
}
