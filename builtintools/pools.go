package builtintools

import (
	"sync"
)

const maxAnySliceCap = 100

var anySlicesPool = sync.Pool{New: func() interface{} { return make([]any, 0, 10) }}

func AcquireAnySlice() []any {
	return anySlicesPool.Get().([]any)
}

func ReleaseAnySlice(val []any) {
	if cap(val) == 0 || cap(val) > maxAnySliceCap {
		return
	}

	for i := range val {
		val[i] = nil
	}
	val = val[:0]
	anySlicesPool.Put(val)
}
