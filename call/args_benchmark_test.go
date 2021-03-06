package call_test

import (
	"testing"

	"github.com/libmonsoon-dev/go-lib/call"
)

func BenchmarkArgsGrowLoop(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var args call.Args
		for i := 0; i < 10; i++ {
			args.Grow(i)
		}
	}
}

func BenchmarkArgsGrow10(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var args call.Args
		args.Grow(10)
	}
}

func BenchmarkArgsGrowPool(b *testing.B) {
	b.ReportAllocs()

	var args *call.Args
	for i := 0; i < b.N; i++ {
		args = call.AcquireArgs()
		args.Grow(10)
		call.ReleaseArgs(args)
	}
}
