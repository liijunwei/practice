package domain

import (
	"time"

	"greenlight/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
)

type Account struct {
	Balance   decimal.Big // Balance = Available + Pending
	Available decimal.Big
	Pending   decimal.Big

	CreatedAt time.Time
	UpdatedAt time.Time

	eventsourcing.BaseAggregate
}

func (acc *Account) Apply(event eventsourcing.Event) error {
	switch event := event.(type) {
	case *AccountCreated:
		acc.Balance = decimal.Big{}
		acc.Available = decimal.Big{}
		acc.Pending = decimal.Big{}

		acc.Balance = *acc.Balance.Copy(event.InitialBalance)
		acc.Available = *acc.Available.Copy(event.InitialBalance)

		acc.CreatedAt = event.CreatedAt
		acc.UpdatedAt = event.CreatedAt
		acc.SetAggregateID(event.AggregateID)
	case *BalanceChanged:
		acc.Available.Add(&acc.Available, event.AvailableDelta)
		acc.Pending.Add(&acc.Pending, event.PendingDelta)

		acc.Balance.Add(&acc.Balance, event.AvailableDelta)
		acc.Balance.Add(&acc.Balance, event.PendingDelta)
	default:
		return &UnsupportedEventError{event: event}
	}

	acc.Version = event.GetVersion()

	return nil
}

const (
	accountInitState    eventsourcing.State = ""
	accountCreatedState eventsourcing.State = "created"
)

func (acc *Account) GetTransitions() []eventsourcing.Transition {
	return []eventsourcing.Transition{
		{
			FromState: accountInitState,
			Event:     &AccountCreated{},
			ToState:   accountCreatedState,
		},
		{
			FromState: accountCreatedState,
			Event:     &BalanceChanged{},
			ToState:   accountCreatedState,
		},
	}
}

// Define user facing APIs

// NewAccount creates a account.
func NewAccount(initBalance *decimal.Big) (*Account, error) {
	// create a init account
	acc := &Account{}

	// create a event
	event := &AccountCreated{
		InitialBalance: initBalance,
	}

	accountID := uuid.Must(uuid.NewV7())
	// fill base event data
	event.SetAggregateID(accountID)
	event.SetVersion(1)
	event.SetCreatedAt(time.Now())

	// apply the event
	if err := acc.Apply(event); err != nil {
		return nil, err
	}
	// record uncommitted events
	acc.AppendChanges(event)

	return acc, nil
}

// MoveAvailableToPending moves balance from available state to pending state.
// It doesn't change total balance.
func (acc *Account) MoveAvailableToPending(amount *decimal.Big) error {
	pendingDelta := &decimal.Big{}
	availableDelta := &decimal.Big{}

	event := &BalanceChanged{
		AvailableDelta: availableDelta.Sub(availableDelta, amount),
		PendingDelta:   pendingDelta.Add(pendingDelta, amount),
	}

	event.SetAggregateID(acc.GetAggregateID())
	event.SetVersion(acc.Version + 1)
	event.SetCreatedAt(time.Now())

	if err := acc.Apply(event); err != nil {
		return err
	}

	acc.AppendChanges(event)

	return nil
}
