package eventstore

import (
	"encoding/json"
	"time"

	"greenlight/internal/eventsourcing"

	"github.com/gofrs/uuid"
)

type EventModel struct {
	AggregateID uuid.UUID
	Version     int
	ParentID    uuid.UUID
	EventType   string
	Payload     json.RawMessage
	CreatedAt   time.Time
}

func (ev *EventModel) ToEvent(
	reg *eventsourcing.EventRegistry,
) (eventsourcing.Event, error) {
	event, err := reg.GetInstance(eventsourcing.EventType(ev.EventType))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ev.Payload, event)
	if err != nil {
		return nil, &EventUnmarshalError{err: err, eventModel: ev}
	}

	event.SetAggregateID(ev.AggregateID)
	event.SetCreatedAt(ev.CreatedAt)
	event.SetVersion(ev.Version)
	event.SetParentID(ev.ParentID)

	return event, nil
}

func buildEventModelFromEventPayload(event eventsourcing.Event) (*EventModel, error) {
	payloadData, err := json.Marshal(event)
	if err != nil {
		return nil, &EventMarshalError{
			err:   err,
			event: event,
		}
	}

	return &EventModel{
		AggregateID: event.GetAggregateID(),
		Version:     event.GetVersion(),
		ParentID:    event.GetParentID(),
		EventType:   string(event.EventType()),
		Payload:     payloadData,
		CreatedAt:   event.GetCreatedAt(),
	}, nil
}
