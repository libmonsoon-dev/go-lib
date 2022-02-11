package errutils

import "fmt"

func Wrap(errorMessage string, err error) error {
	if err != nil {
		err = fmt.Errorf(errorMessage+": %w", err)
	}

	return err
}
