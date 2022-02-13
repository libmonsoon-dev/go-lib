package call_test

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/libmonsoon-dev/go-lib/call"
)

func TestChainDo(t *testing.T) {
	type test struct {
		chain          *call.Chain
		doCallsArgs    []call.DoArgs
		makeCallsArgs  []call.MakeArgs[any]
		expectedErr    error
		expectedResult call.Args
	}

	tests := []test{
		{
			chain: call.NewChain(call.ChainConfig{}),
			doCallsArgs: []call.DoArgs{
				{ErrorMessage: "step 1", Func: emptyFunc},
				{ErrorMessage: "step 2", Func: emptyFunc},
				{ErrorMessage: "step 3", Func: emptyFunc},
			},
		},
		{
			chain: call.AcquireChain(),
			doCallsArgs: []call.DoArgs{
				{ErrorMessage: "step 1", Func: emptyFunc},
				{ErrorMessage: "step 2", Func: emptyFunc},
				{ErrorMessage: "step 3", Func: emptyFunc},
			},
		},
		{
			chain: &call.Chain{},
			doCallsArgs: []call.DoArgs{
				{ErrorMessage: "step 1", Func: emptyFunc},
				{ErrorMessage: "step 2", Func: emptyFunc},
				{ErrorMessage: "step 3", Func: emptyFunc},
			},
		},
		func() test {
			chain := call.NewChain(call.ChainConfig{})
			getFoo := call.MakeArgGetter[*foo](chain, 0)

			return test{
				chain: chain,
				makeCallsArgs: []call.MakeArgs[any]{
					{ErrorMessage: "step 1", Func: func() (any, error) { return NewFoo() }},
					{ErrorMessage: "step 2", Func: func() (any, error) { return NewBar(getFoo()) }},
					{ErrorMessage: "step 3", Func: returnEOF},
					{ErrorMessage: "unreachable", Func: func() (any, error) { panic("should not be called") }},
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
				makeCallsArgs: []call.MakeArgs[any]{
					{ErrorMessage: "step 1", Func: func() (any, error) { return NewFoo() }},
					{ErrorMessage: "step 2", Func: func() (any, error) { return NewBar(getFoo()) }},
					{ErrorMessage: "step 3", Func: returnEOF},
					{ErrorMessage: "unreachable", Func: func() (any, error) { panic("should not be called") }},
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
				makeCallsArgs: []call.MakeArgs[any]{
					{ErrorMessage: "step 1", Func: func() (any, error) { return NewFoo() }},
					{ErrorMessage: "step 2", Func: func() (any, error) { return NewBar(getFoo()) }},
					{ErrorMessage: "step 3", Func: returnEOF},
					{ErrorMessage: "unreachable", Func: func() (any, error) { panic("should not be called") }},
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
				test.chain.Do(args)
			}

			for _, args := range test.makeCallsArgs {
				test.chain.Make(args)
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
