package errors

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

// Format wraps err with fmt.Errorf if err is not nil
// If err is nil returns nil
func Format(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	if strings.Contains(format, "%w") {
		args = append(args, err)
	} else {
		slog.Debug("format " + strconv.Quote(format) + " not contains %w")
	}

	return fmt.Errorf(format, args...)
}

// Appendf returns an error that wraps the given errors.
// Any nil error values are discarded.
// Appendf returns nil if every value in errs is nil.
// The error formats as the concatenation of the strings obtained
// by calling the Error method of each element of errs, with a newline between each string.
// Second error will be wrapped with Format
//
// A non-nil error returned by Appendf implements the Unwrap() []error method.
func Appendf(first, second error, format string, args ...any) error {
	return Join(first, Format(second, format, args...))
}
