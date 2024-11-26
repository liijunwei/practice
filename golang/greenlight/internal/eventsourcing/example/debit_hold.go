package main

import (
	"fmt"
	"time"

	"knockout/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
)

// Define your Aggregate struct that match your domain.
type DebitHold struct {
	AccountID uuid.UUID
	Amount    *decimal.Big
	CreatedAt time.Time
	UpdatedAt time.Time

	Status string

	// BaseAggregate provides common functionality.
	eventsourcing.BaseAggregate
}

// ensure Account implements Aggregate interface.
var _ eventsourcing.Aggregate = &DebitHold{}

// EventTable return the event table name.
func (dh *DebitHold) EventTable() string {
	return "debit_hold_event"
}

// Apply defines how aggregate handles events.
func (dh *DebitHold) Apply(event eventsourcing.Event) error {
	// state machine
	newState, err := eventsourcing.TransistOnEvent(dh, event)
	if err != nil {
		return fmt.Errorf("no valid transition: %w", err)
	}

	dh.Status = string(newState)

	// apply changes
	switch event := event.(type) {
	case *DebitHoldCreated:
		dh.AccountID = event.GetParentID()
		dh.Amount = event.Amount
		dh.CreatedAt = event.CreatedAt
		dh.UpdatedAt = event.CreatedAt
		dh.SetAggregateID(event.GetAggregateID())
	case *DebitHoldReturned:
		dh.UpdatedAt = event.CreatedAt
	case *DebitHoldReleased:
		dh.UpdatedAt = event.CreatedAt
	default:
		return &UnsupportedEventError{event: event}
	}

	// update aggregate version
	dh.Version = event.GetVersion()

	return nil
}

// Define StateMachine related method for your aggregate.

// GetCurrentState returns current status
func (dh *DebitHold) GetCurrentState() eventsourcing.State {
	return eventsourcing.State(dh.Status)
}

// Define state constant
const (
	initState     eventsourcing.State = ""
	createdState  eventsourcing.State = "created"
	returnedState eventsourcing.State = "returned"
	releasedState eventsourcing.State = "released"
)

// GetStates returns all possible state transitions
func (dh *DebitHold) GetTransitions() []eventsourcing.Transition {
	return []eventsourcing.Transition{
		{
			FromState: initState,
			Event:     &DebitHoldCreated{},
			ToState:   createdState,
		},
		{
			FromState: createdState,
			Event:     &DebitHoldReturned{},
			ToState:   returnedState,
		},
		{
			FromState: createdState,
			Event:     &DebitHoldReleased{},
			ToState:   releasedState,
		},
	}
}

// Define user facing APIs

// NewDebitHold create a uncommitted debit hold.
func NewDebitHold(accountID uuid.UUID, amount *decimal.Big) (*DebitHold, error) {
	// create a init account
	request := &DebitHold{}

	// create a event
	event := &DebitHoldCreated{
		Amount: amount,
	}
	requestID := uuid.Must(uuid.NewV4())
	// fill base event data
	event.SetAggregateID(requestID)
	event.SetParentID(accountID)
	event.SetVersion(1)
	event.SetCreatedAt(time.Now())

	// apply the event
	if err := request.Apply(event); err != nil {
		return nil, err
	}
	// record uncommitted events
	request.AppendChanges(event)

	return request, nil
}
