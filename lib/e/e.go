package e

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s : %w", msg, err)
}

// WrapIfErr is needed to treat err = nil and other errors differently

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}

	return Wrap(msg, err)
}
