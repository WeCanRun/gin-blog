package errcode

import "fmt"

type InternalError struct {
	err    error
	status int
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%d: %s", e.status, e.err.Error())
}
