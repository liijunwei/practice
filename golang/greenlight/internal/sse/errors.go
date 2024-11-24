package sse

import "fmt"

type InvalidEventError struct {
	msg string
}

func (e *InvalidEventError) Error() string {
	return fmt.Sprintf("invalid event: %s", e.msg)
}
