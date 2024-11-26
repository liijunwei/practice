package db

import (
	"encoding/json"
	"reflect"
	"time"

	"greenlight/internal/eventsourcing"

	"github.com/google/uuid"
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

func NewEventModelFromEvent(event eventsourcing.Event) (*EventModel, error) {
	payload := eventPayload(event)

	payloadBytes, err := json.Marshal(payload)
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
		Payload:     payloadBytes,
		CreatedAt:   event.GetCreatedAt(),
	}, nil
}

func eventPayload(event eventsourcing.Event) map[string]any {
	payload := make(map[string]any)

	eventType := reflect.TypeOf(event).Elem()
	value := reflect.ValueOf(event).Elem()

	for i := 0; i < eventType.NumField(); i++ {
		field := eventType.Field(i)

		if field.Name == "BaseEvent" && field.Type.Kind() == reflect.Struct {
			continue
		}

		payload[field.Name] = value.FieldByName(field.Name).Interface()
	}

	return payload
}
