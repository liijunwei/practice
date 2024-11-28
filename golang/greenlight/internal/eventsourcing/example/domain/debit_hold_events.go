package domain

import (
	"greenlight/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
)

// Define other events
type DebitHoldCreated struct {
	Amount *decimal.Big `json:"amount"`

	eventsourcing.BaseEvent
}

func (dre *DebitHoldCreated) EventType() eventsourcing.EventType {
	return "debit_hold.created"
}

type DebitHoldReturned struct {
	eventsourcing.BaseEvent
}

func (dse *DebitHoldReturned) EventType() eventsourcing.EventType {
	return "debit_hold.returned"
}

type DebitHoldReleased struct {
	eventsourcing.BaseEvent
}

func (dfe *DebitHoldReleased) EventType() eventsourcing.EventType {
	return "debit_hold.released"
}
