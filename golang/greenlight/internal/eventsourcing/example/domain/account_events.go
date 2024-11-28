package domain

import (
	"greenlight/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
)

type AccountCreated struct {
	InitialBalance *decimal.Big `json:"initial_balance"`

	eventsourcing.BaseEvent
}

func (ace *AccountCreated) EventType() eventsourcing.EventType {
	return "account.created"
}

type BalanceChanged struct {
	AvailableDelta *decimal.Big `json:"available_delta"`
	PendingDelta   *decimal.Big `json:"pending_delta"`

	eventsourcing.BaseEvent
}

func (de *BalanceChanged) EventType() eventsourcing.EventType {
	return "account.balance_changed"
}
