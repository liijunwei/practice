package main

import (
	"knockout/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
)

type AccountCreated struct {
	InitialBalance *decimal.Big

	// BaseEvent provides common event fields and functionality like Get/SetVersion
	eventsourcing.BaseEvent
}

// EventType returns the name of event
func (ace *AccountCreated) EventType() eventsourcing.EventType {
	return "account.created"
}

type BalanceChanged struct {
	AvailableDelta *decimal.Big
	PendingDelta   *decimal.Big

	eventsourcing.BaseEvent
}

func (de *BalanceChanged) EventType() eventsourcing.EventType {
	return "account.balance_changed"
}
