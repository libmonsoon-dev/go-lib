package app

import (
	"log/slog"
	"os"

	"github.com/libmonsoon-dev/go-lib/errors"
)

func Must[T any](val T, err error) T {
	Check(err)
	return val
}

func Checkf(err error, format string, args ...any) {
	Check(errors.Format(err, format, args...))
}

func Check(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
