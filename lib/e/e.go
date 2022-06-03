package e

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("cant %s request:%w", msg, err)
}