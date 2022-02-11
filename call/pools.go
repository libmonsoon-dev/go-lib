package call

import (
	"github.com/libmonsoon-dev/go-lib/builtintools"
	"sync"
)

var chainsPool = sync.Pool{New: func() interface{} {
	return NewChain(ChainConfig{
		Args: *builtintools.AcquireAnySlice(),
	})
}}

var groupPool = sync.Pool{New: func() interface{} {
	return NewGroup(GroupConfig{})
}}

func AcquireChain() *Chain {
	return chainsPool.Get().(*Chain)
}

func ReleaseChain(val *Chain) {
	val.Reset()
	chainsPool.Put(val)
}

func AcquireGroup() *Group {
	return groupPool.Get().(*Group)
}

func ReleaseGroup(val *Group) {
	val.Reset()
	groupPool.Put(val)
}

func AcquireArgs() *Args {
	return (*Args)(builtintools.AcquireAnySlice())
}
func ReleaseArgs(val *Args) {
	builtintools.ReleaseAnySlice((*[]any)(val))
}
