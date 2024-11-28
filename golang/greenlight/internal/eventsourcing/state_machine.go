package eventsourcing

type State string

type Transition struct {
	FromState State
	ToState   State
	Event     Event
}

type StateMachine interface {
	GetCurrentState() State
	GetTransitions() []Transition
}

func TransistOnEvent(stateMachine StateMachine, event Event) (State, error) {
	currentState := stateMachine.GetCurrentState()

	transitions := stateMachine.GetTransitions()
	eventType := event.EventType()

	for _, transition := range transitions {
		if transition.FromState == currentState && transition.Event.EventType() == eventType {
			return transition.ToState, nil
		}
	}

	return "", &NoTransitionError{
		event: event,
		state: currentState,
	}
}
