package errutils

import (
	"bytes"

	"golang.org/x/xerrors"
)

const bundleSeparator = "\n"

var _ error = (*Bundle)(nil)
var _ xerrors.Wrapper = (*Bundle)(nil)

func NewBundle() *Bundle {
	return &Bundle{}
}

type Bundle struct {
	errs []error
}

func (b *Bundle) Unwrap() error {
	if b == nil || len(b.errs) < 1 {
		return nil
	}

	return b.errs[0]
}

func (b *Bundle) List() []error {
	if b == nil {
		return nil
	}
	return b.errs
}

func (b *Bundle) Add(errs ...error) {
	for _, err := range errs {
		if err != nil {
			b.errs = append(b.errs, err)
		}
	}
}

func (b *Bundle) Error() string {
	buf := bytes.NewBuffer(nil)
	lastIndex := len(b.errs) - 1
	for i, err := range b.errs {
		ignoreWriteErr(buf.WriteString(err.Error()))
		if i != lastIndex {
			ignoreWriteErr(buf.WriteString(bundleSeparator))
		}
	}
	return buf.String()
}

func (b *Bundle) ErrorOrNil() error {
	if b == nil {
		return nil
	}
	if !b.IsEmpty() {
		return b
	}
	return nil
}

func (b *Bundle) IsEmpty() bool {
	return len(b.errs) == 0
}

func ignoreWriteErr(n int, err error) {
}
