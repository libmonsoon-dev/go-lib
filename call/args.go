package call

import (
	"fmt"
	"strings"
)

type Args []any

func (a Args) String() string {
	if a == nil {
		return "nil"
	}

	if len(a) == 0 {
		return "[]"
	}
	format := strings.Repeat("%T, ", len(a))
	format = "[" + format[:len(format)-2] + "]"
	return fmt.Sprintf(format, a...)
}

func (a *Args) Reset() {
	for i := range *a {
		(*a)[i] = nil
	}

	*a = (*a)[:0]
}

func (a *Args) GetArgs() Args { return *a }

func (a *Args) Grow(n int) {
	if n < 0 {
		panic("call.Args.Grow: negative count")
	}

	if cap(*a)-len(*a) < n {
		a.grow(n)
	}
}

func (a *Args) Copy() Args {
	if a == nil {
		return nil
	}

	return append(Args(nil), *a...)
}

func (a *Args) grow(n int) {
	next := make(Args, len(*a), 2*cap(*a)+n)
	copy(next, *a)
	*a = next
}
