package fs

import "sync"

var archiverPool = sync.Pool{New: func() interface{} {
	return newArchiver()
}}

func acquireArchiver() *archiver {
	return archiverPool.Get().(*archiver)
}

func releaseArchiver(a *archiver) {
	a.Reset()
	archiverPool.Put(a)
}
