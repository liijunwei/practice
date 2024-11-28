package eventsourcing

import "fmt"

type NoTransitionError struct {
	event Event
	state State
}

func (nte *NoTransitionError) Error() string {
	return fmt.Sprintf("no valid transition: state %s, event %s",
		nte.state,
		nte.event.EventType())
}

type EventNotRegisteredError struct {
	event EventType
}

func (enre *EventNotRegisteredError) Error() string {
	return fmt.Sprintf("event not registered: %s", enre.event)
}
