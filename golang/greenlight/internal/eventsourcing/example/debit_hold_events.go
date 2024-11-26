package main

import (
	"knockout/internal/eventsourcing"

	"github.com/ericlagergren/decimal"
)

// Define other events
type DebitHoldCreated struct {
	Amount *decimal.Big

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
