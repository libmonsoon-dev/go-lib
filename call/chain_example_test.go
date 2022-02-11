package call_test

import (
	"fmt"
	"github.com/libmonsoon-dev/go-lib/call"
	"io"
)

func ExampleChain() {
	callChain := call.NewChain(call.ChainConfig{})

	const (
		fooArgIndex = iota
		barArgIndex
	)

	// args getters:
	getFoo := call.MakeArgGetter[*foo](callChain, fooArgIndex)
	getBar := call.MakeArgGetter[*bar](callChain, barArgIndex)

	// functions evocations:
	callChain.
		Make("new foo", func() (any, error) { return NewFoo() }).
		Make("new bar", func() (any, error) { return NewBar(getFoo()) }).
		Make("new baz", func() (any, error) {
			foo := getFoo()
			bar := getBar()

			return NewBaz(foo, bar)
		}).
		Make("second foo", func() (any, error) { return NewFoo() }).
		Do("return first error", func() error { return io.EOF }).
		Make("unreachable", func() (any, error) { return NewBar(nil) })

	// Output: Arguments: [*call_test.foo, *call_test.bar, *call_test.baz, *call_test.foo]
	// Error: "return first error: EOF"
	fmt.Println("Arguments:", callChain.GetArgs())
	fmt.Printf("Error: %q\n", callChain.GetError())
}

type foo struct{}

func NewFoo() (*foo, error) {
	return &foo{}, nil
}

type bar struct {
	f *foo
}

func NewBar(f *foo) (*bar, error) {
	return &bar{f: f}, nil
}

type baz struct {
	f *foo
	b *bar
}

func NewBaz(f *foo, b *bar) (*baz, error) {
	return &baz{f, b}, nil
}
