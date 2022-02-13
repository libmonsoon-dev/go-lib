package call_test

import (
	"testing"

	"github.com/libmonsoon-dev/go-lib/call"
)

var preventingInliningChain *call.Chain

func BenchmarkNewChainDo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = call.NewChain(call.ChainConfig{})
		preventingInliningChain.
			DoFunc("empty", emptyFunc).
			DoFunc("empty", emptyFunc).
			DoFunc("empty", emptyFunc)
	}
}

func BenchmarkZeroValueDo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = (&call.Chain{})
		preventingInliningChain.
			DoFunc("empty", emptyFunc).
			DoFunc("empty", emptyFunc).
			DoFunc("empty", emptyFunc)
	}
}

func BenchmarkAcquireChainDo(b *testing.B) {
	b.ReportAllocs()

	var c *call.Chain
	for i := 0; i < b.N; i++ {
		c = call.AcquireChain()
		c.
			DoFunc("empty", emptyFunc).
			DoFunc("empty", emptyFunc).
			DoFunc("empty", emptyFunc)

		call.ReleaseChain(c)
	}
}

func BenchmarkNewChainMake(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = call.NewChain(call.ChainConfig{})
		preventingInliningChain.
			MakeFunc("make foo", makeFoo).
			MakeFunc("make foo", makeFoo).
			MakeFunc("make foo", makeFoo)
	}
}

func BenchmarkZeroValueMake(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = (&call.Chain{})
		preventingInliningChain.
			MakeFunc("make foo", makeFoo).
			MakeFunc("make foo", makeFoo).
			MakeFunc("make foo", makeFoo)
	}
}

func BenchmarkAcquireChainMake(b *testing.B) {
	b.ReportAllocs()

	var c *call.Chain
	for i := 0; i < b.N; i++ {
		c = call.AcquireChain()
		c.
			MakeFunc("make foo", makeFoo).
			MakeFunc("make foo", makeFoo).
			MakeFunc("make foo", makeFoo)

		call.ReleaseChain(c)
	}
}

func makeFoo() (any, error) {
	return NewFoo()
}
