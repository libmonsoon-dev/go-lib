package ctxbind

import (
	"context"
	"errors"
	"strings"
)

var ContextCanceled = context.Canceled

func IsContextCanceled(err error) bool {
	return isSameError(err, ContextCanceled)
}

func isSameError(err, target error) bool {
	return errors.Is(err, target) || containsErr(err, target)
}

func containsErr(err, target error) bool {
	if err == nil || target == nil {
		return false
	}

	return strings.HasSuffix(err.Error(), target.Error())
}
