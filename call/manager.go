package call

import "context"

type Manager = manager[any]

type manager[T any] interface {
	Do(DoArgs) Manager
	DoFunc(errorMessage string, fn DoFunc) Manager
	Make(MakeArgs[T]) Manager
	MakeFunc(errorMessage string, fn MakeFunc[T]) Manager

	ArgsGetter
	SetArgs(args Args) Manager
	GetError() error
	GetResult() (args Args, err error)
}

type ContextGetter interface {
	GetContext() context.Context
}
