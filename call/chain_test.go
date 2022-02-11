package call_test

import (
	"fmt"
	"github.com/libmonsoon-dev/go-lib/call"
	"io"
	"reflect"
	"testing"
)

func TestChainDo(t *testing.T) {
	type doCallsArgs struct {
		errorMessage string
		fn           func() error
	}

	type makeCallsArgs struct {
		errorMessage string
		fn           func() (any, error)
	}

	type test struct {
		chain          *call.Chain
		doCallsArgs    []doCallsArgs
		makeCallsArgs  []makeCallsArgs
		expectedErr    error
		expectedResult call.Args
	}

	tests := []test{
		{
			chain: call.NewChain(call.ChainConfig{}),
			doCallsArgs: []doCallsArgs{
				{"step 1", emptyFunc},
				{"step 2", emptyFunc},
				{"step 3", emptyFunc},
			},
		},
		{
			chain: call.AcquireChain(),
			doCallsArgs: []doCallsArgs{
				{"step 1", emptyFunc},
				{"step 2", emptyFunc},
				{"step 3", emptyFunc},
			},
		},
		{
			chain: &call.Chain{},
			doCallsArgs: []doCallsArgs{
				{"step 1", emptyFunc},
				{"step 2", emptyFunc},
				{"step 3", emptyFunc},
			},
		},
		func() test {
			chain := call.NewChain(call.ChainConfig{})
			getFoo := call.MakeArgGetter[*foo](chain, 0)

			return test{
				chain: chain,
				makeCallsArgs: []makeCallsArgs{
					{"step 1", func() (any, error) { return NewFoo() }},
					{"step 2", func() (any, error) { return NewBar(getFoo()) }},
					{"step 3", returnEOF},
					{"unreachable", func() (any, error) { panic("should not be called") }},
				},
				expectedResult: func() call.Args {
					foo, _ := NewFoo()
					bar, _ := NewBar(foo)
					return call.Args{foo, bar, nil}
				}(),
				expectedErr: fmt.Errorf("step 3: %w", io.EOF),
			}
		}(),
		func() test {
			chain := &call.Chain{}
			getFoo := call.MakeArgGetter[*foo](chain, 0)

			return test{
				chain: chain,
				makeCallsArgs: []makeCallsArgs{
					{"step 1", func() (any, error) { return NewFoo() }},
					{"step 2", func() (any, error) { return NewBar(getFoo()) }},
					{"step 3", returnEOF},
					{"unreachable", func() (any, error) { panic("should not be called") }},
				},
				expectedResult: func() call.Args {
					foo, _ := NewFoo()
					bar, _ := NewBar(foo)
					return call.Args{foo, bar, nil}
				}(),
				expectedErr: fmt.Errorf("step 3: %w", io.EOF),
			}
		}(),
		func() test {
			chain := call.AcquireChain()
			getFoo := call.MakeArgGetter[*foo](chain, 0)

			return test{
				chain: chain,
				makeCallsArgs: []makeCallsArgs{
					{"step 1", func() (any, error) { return NewFoo() }},
					{"step 2", func() (any, error) { return NewBar(getFoo()) }},
					{"step 3", returnEOF},
					{"unreachable", func() (any, error) { panic("should not be called") }},
				},
				expectedResult: func() call.Args {
					foo, _ := NewFoo()
					bar, _ := NewBar(foo)
					return call.Args{foo, bar, nil}
				}(),
				expectedErr: fmt.Errorf("step 3: %w", io.EOF),
			}
		}(),
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%d", i+1), func(t *testing.T) {
			for _, args := range test.doCallsArgs {
				test.chain.Do(args.errorMessage, args.fn)
			}

			for _, args := range test.makeCallsArgs {
				test.chain.Make(args.errorMessage, args.fn)
			}

			if !reflect.DeepEqual(test.expectedErr, test.chain.GetError()) {
				t.Errorf("Unexpected errors. Expected: %v, got: %v", test.expectedErr, test.chain.GetError())
			}

			if len(test.expectedResult) == 0 && len(test.chain.GetArgs()) == 0 {
				return
			}
			if !reflect.DeepEqual(test.expectedResult, test.chain.GetArgs()) {
				t.Errorf("Unexpected results. Expected: %#v, got: %#v", test.expectedResult, test.chain.GetArgs())
			}
		})
	}
}

func emptyFunc() error { return nil }

func returnEOF() (any, error) { return nil, io.EOF }
