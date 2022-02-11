package call

type ArgsGetter interface {
	GetArgs() Args
}

// MakeArgGetter creates a function that returns an argument at index n from chain c that asserted to type T
// You can create (but not use) this getter before the corresponding argument is in the corresponding chain.
//It is the caller's responsibility to ensure that this argument is provided and can be type asserted to T.
func MakeArgGetter[T any](source ArgsGetter, index int, indexes ...int) func() T {
	return func() T {
		item := source.GetArgs()[index]
		for _, index = range indexes {
			item = item.(Args)[index]
		}
		return item.(T)
	}
}

// TODO: tests
func MakeArgsGetter[T any](source ArgsGetter, from, to int, indexes ...int) func() []T {
	return func() []T {
		var item any = source.GetArgs()
		for _, index := range indexes {
			item = item.(Args)[index]
		}

		result := make([]T, 0, to-from)
		for _, item = range item.(Args)[from:to] {
			result = append(result, item.(T))
		}
		return result
	}
}
