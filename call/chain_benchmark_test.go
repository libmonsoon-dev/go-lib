package call_test

import (
	"github.com/libmonsoon-dev/go-lib/call"
	"testing"
)

var preventingInliningChain *call.Chain

func BenchmarkNewChainDo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = call.NewChain(call.ChainConfig{}).
			Do("empty", emptyFunc).
			Do("empty", emptyFunc).
			Do("empty", emptyFunc)
	}
}

func BenchmarkZeroValueDo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = (&call.Chain{}).
			Do("empty", emptyFunc).
			Do("empty", emptyFunc).
			Do("empty", emptyFunc)
	}
}

func BenchmarkAcquireChainDo(b *testing.B) {
	b.ReportAllocs()

	var c *call.Chain
	for i := 0; i < b.N; i++ {
		c = call.AcquireChain().
			Do("empty", emptyFunc).
			Do("empty", emptyFunc).
			Do("empty", emptyFunc)

		call.ReleaseChain(c)
	}
}

func BenchmarkNewChainMake(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = call.NewChain(call.ChainConfig{}).
			Make("empty", makeFoo).
			Make("empty", makeFoo).
			Make("empty", makeFoo)
	}
}

func BenchmarkZeroValueMake(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		preventingInliningChain = (&call.Chain{}).
			Make("empty", makeFoo).
			Make("empty", makeFoo).
			Make("empty", makeFoo)
	}
}

func BenchmarkAcquireChainMake(b *testing.B) {
	b.ReportAllocs()

	var c *call.Chain
	for i := 0; i < b.N; i++ {
		c = call.AcquireChain().
			Make("empty", makeFoo).
			Make("empty", makeFoo).
			Make("empty", makeFoo)

		call.ReleaseChain(c)
	}
}

func makeFoo() (any, error) {
	return NewFoo()
}
