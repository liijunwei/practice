package eventsourcing

import (
	"time"

	"github.com/gofrs/uuid"
)

type EventType string

type Event interface {
	EventType() EventType
	GetAggregateID() uuid.UUID
	SetAggregateID(uuid.UUID)
	GetParentID() uuid.UUID
	SetParentID(uuid.UUID)
	GetVersion() int
	SetVersion(int)
	GetCreatedAt() time.Time
	SetCreatedAt(time.Time)
}

type BaseEvent struct {
	ID          uuid.UUID
	AggregateID uuid.UUID
	ParentID    uuid.UUID
	Version     int
	CreatedAt   time.Time
}

var _ Event = (*BaseEvent)(nil)

func (*BaseEvent) EventType() EventType {
	panic("implement me in derived event")
}

func (be *BaseEvent) GetAggregateID() uuid.UUID {
	return be.AggregateID
}

func (be *BaseEvent) SetAggregateID(id uuid.UUID) {
	be.AggregateID = id
}

func (be *BaseEvent) GetParentID() uuid.UUID {
	return be.ParentID
}

func (be *BaseEvent) SetParentID(parentID uuid.UUID) {
	be.ParentID = parentID
}

func (be *BaseEvent) GetVersion() int {
	return be.Version
}

func (be *BaseEvent) SetVersion(version int) {
	be.Version = version
}

func (be *BaseEvent) GetCreatedAt() time.Time {
	return be.CreatedAt
}

func (be *BaseEvent) SetCreatedAt(createdAt time.Time) {
	be.CreatedAt = createdAt
}
