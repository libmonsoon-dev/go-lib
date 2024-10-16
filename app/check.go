package app

import (
	"log/slog"
	"os"

	"github.com/libmonsoon-dev/go-lib/errors"
)

func Mustf[T any](val T, err error, format string, args ...any) T {
	Check(errors.Format(err, format, args...))
	return val
}

func Must[T any](val T, err error) T {
	Check(err)
	return val
}

func Check(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
