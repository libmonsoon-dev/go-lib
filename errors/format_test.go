package errors_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/libmonsoon-dev/go-lib/errors"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	type args struct {
		err    error
		format string
		args   []any
	}
	tests := []struct {
		name     string
		args     args
		expected error
	}{
		{
			name: "Return nil if nil arg",
			args: args{
				format: "do somthing with str %s: %w",
				args:   []any{"some string"},
			},
			expected: nil,
		},
		{
			name: "Return wrapped error",
			args: args{
				err:    io.EOF,
				format: "do somthing with str %s: %w",
				args:   []any{"some string"},
			},
			expected: fmt.Errorf("do somthing with str %s: %w", "some string", io.EOF),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.Format(tt.args.err, tt.args.format, tt.args.args...)
			require.Equal(t, tt.expected, err)
		})
	}
}
