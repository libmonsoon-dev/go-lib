package call

type MakeArgs[T any] struct {
	Object Maker[T]
	Func   MakeFunc[T]

	ChainFunc
	GroupFunc

	ErrorMessage     string
	ErrorMessageFunc func() string
}

func (args MakeArgs[T]) GetFunc() MakeFunc[T] {
	if args.Func != nil {
		return args.Func
	}

	if args.Object != nil {
		return args.Object.Make
	}

	return nil
}

func (args MakeArgs[T]) GetObject() Maker[T] {
	if args.Object != nil {
		return args.Object
	}

	if args.Func != nil {
		return args.Func
	}

	return nil
}

func (args MakeArgs[T]) GetErrorMessage() string {
	if args.ErrorMessage != "" {
		return args.ErrorMessage
	}

	if args.ErrorMessageFunc != nil {
		return args.ErrorMessageFunc()
	}

	panic("Both args.ErrorMessage and args.ErrorMessageFunc are not set")
}

type Maker[T any] interface {
	Make() (T, error)
}

type MakeFunc[T any] func() (T, error)

func (m MakeFunc[T]) Make() (T, error) { return m() }

type MakerFuncAdaptor[T any] func() (T, error)

func (m MakerFuncAdaptor[T]) Make() (any, error) { return m() }

var _ Maker[int] = MakeFunc[int](nil)
var _ Maker[any] = MakeFunc[any](nil)

func NewMakerAdapter[T any](input Maker[T]) MakerAdapter[T] {
	return MakerAdapter[T]{
		input: input,
	}
}

type MakerAdapter[T any] struct {
	input Maker[T] //nolint:structcheck
}

func (m MakerAdapter[T]) Make() (any, error) {
	return m.input.Make()
}

var _ Maker[any] = NewMakerAdapter(Maker[int](nil))

func MakeFuncAdapter[T any](fn MakeFunc[T]) MakeFunc[any] {
	return func() (any, error) {
		return fn()
	}
}

var _ Maker[any] = MakeFuncAdapter(MakeFunc[int](nil))
