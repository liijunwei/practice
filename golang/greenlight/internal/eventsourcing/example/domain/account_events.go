package domain

import (
	"greenlight/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
)

type AccountCreated struct {
	InitialBalance *decimal.Big

	eventsourcing.BaseEvent
}

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
