package app

import (
	"log/slog"
	"os"
)

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
