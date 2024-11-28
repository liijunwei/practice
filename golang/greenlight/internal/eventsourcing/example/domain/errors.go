package domain

import (
	"fmt"

	"greenlight/internal/eventsourcing"
)

type UnsupportedEventError struct {
	event eventsourcing.Event
}

func (e *UnsupportedEventError) Error() string {
	return fmt.Sprintf("unsupported event: %v", e.event.EventType())
}

type NotEnoughBalanceError struct{}

func (e *NotEnoughBalanceError) Error() string {
	return "not enough balance"
}

type TypeMismatchError struct {
	expect any
	got    any
}

func (tme *TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch, expect %t, got %t", tme.expect, tme.got)
}
