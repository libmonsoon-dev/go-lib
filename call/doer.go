package call

type DoArgs struct {
	Object Doer
	Func   DoFunc

	ChainFunc
	GroupFunc

	ErrorMessage     string
	ErrorMessageFunc func() string
}

func (args DoArgs) GetFunc() DoFunc {
	if args.Func != nil {
		return args.Func
	}

	if args.Object != nil {
		return args.Object.Do
	}

	return nil
}

func (args DoArgs) GetObject() Doer {
	if args.Object != nil {
		return args.Object
	}

	if args.Func != nil {
		return args.Func
	}

	return nil
}

func (args DoArgs) GetErrorMessage() string {
	if args.ErrorMessage != "" {
		return args.ErrorMessage
	}

	if args.ErrorMessageFunc != nil {
		return args.ErrorMessageFunc()
	}

	panic("Both args.ErrorMessage and args.ErrorMessageFunc are not set")
}

type Doer interface {
	Do() error
}

type DoFunc func() error

func (d DoFunc) Do() error {
	return d()
}

var _ Doer = DoFunc(nil)
