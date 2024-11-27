package eventsourcing

import "github.com/gofrs/uuid"

type Aggregate interface {
	Apply(event Event) error
	GetChanges() []Event
	AppendChanges(event Event)
	SetAggregateID(uuid.UUID)
	GetAggregateID() uuid.UUID
	GetVersion() int
	SetVersion(int)
	StateMachine
}

// BaseAggregate implements common functionality for Aggregate interface.
//
// It doesn't implement Apply, EventTable.
type BaseAggregate struct {
	ID                uuid.UUID
	Version           int
	uncommittedEvents []Event
}

var _ Aggregate = (*BaseAggregate)(nil)

func (ba *BaseAggregate) GetAggregateID() uuid.UUID {
	return ba.ID
}

func (*BaseAggregate) Apply(_ Event) error {
	panic("implement me in derived aggregate")
}

func (*BaseAggregate) EventTable() string {
	panic("implement me in derived aggregate")
}

func (ba *BaseAggregate) SetAggregateID(id uuid.UUID) {
	ba.ID = id
}

func (ba *BaseAggregate) GetChanges() []Event {
	return ba.uncommittedEvents
}

func (ba *BaseAggregate) AppendChanges(event Event) {
	ba.uncommittedEvents = append(ba.uncommittedEvents, event)
}

func (ba *BaseAggregate) GetVersion() int {
	return ba.Version
}

func (ba *BaseAggregate) SetVersion(version int) {
	ba.Version = version
}

func (ba *BaseAggregate) GetCurrentState() State {
	return ""
}

func (ba *BaseAggregate) GetTransitions() []Transition {
	return []Transition{}
}

func (ba *BaseAggregate) SkipTransition(Event) bool {
	return false
}
