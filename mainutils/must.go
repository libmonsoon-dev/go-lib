package mainutils

func Must[T any](val T, err error) T {
	check(err)
	return val
}

func Must2[A, B any](a A, b B, err error) (A, B) {
	check(err)
	return a, b
}

func Must3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	check(err)
	return a, b, c
}

func Check(err error) {
	check(err)
}

func check(err error) {
	if err != nil {
		DieFunc(err)
	}
}

// DieFunc called with error if any. It must terminate process in some way.
// You need to change to panic-based implementation if you want to correct defers behavior
var DieFunc = logAndExit
