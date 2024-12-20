package domain

import (
	"fmt"
	"time"

	"greenlight/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
)

var _ eventsourcing.Aggregate = &DebitHold{}

type DebitHold struct {
	AccountID uuid.UUID
	Amount    *decimal.Big
	CreatedAt time.Time
	UpdatedAt time.Time

	Status string

	eventsourcing.BaseAggregate
}

func (dh *DebitHold) DomainName() string {
	return "debit_hold"
}

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

	dh.Version = event.GetVersion()

	return nil
}

func (dh *DebitHold) GetCurrentState() eventsourcing.State {
	return eventsourcing.State(dh.Status)
}

const (
	initState     eventsourcing.State = ""
	createdState  eventsourcing.State = "created"
	returnedState eventsourcing.State = "returned"
	releasedState eventsourcing.State = "released"
)

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

func NewDebitHold(accountID uuid.UUID, amount *decimal.Big) (*DebitHold, error) {
	request := &DebitHold{}

	event := &DebitHoldCreated{
		Amount: amount,
	}
	requestID := uuid.Must(uuid.NewV7())
	event.SetAggregateID(requestID)
	event.SetParentID(accountID)
	event.SetVersion(1)
	event.SetCreatedAt(time.Now())

	if err := request.Apply(event); err != nil {
		return nil, err
	}

	request.AppendChanges(event)

	return request, nil
}
