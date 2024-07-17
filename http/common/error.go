package common

import "fmt"

type CodeError struct {
	Code int
	Err  error
}

func (ce *CodeError) Error() string {
	return fmt.Sprintf("code: %d, error: %s", ce.Code, ce.Err)
}
