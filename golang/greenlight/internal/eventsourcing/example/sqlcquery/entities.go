// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlcquery

import (
	"database/sql"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
)

type Account struct {
	ID        uuid.UUID
	Balance   decimal.Big
	Available decimal.Big
	Pending   decimal.Big
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int32
}

type AccountEvent struct {
	AggregateID uuid.UUID
	Version     int32
	ParentID    uuid.UUID
	EventType   sql.NullString
	Payload     []byte
	CreatedAt   time.Time
}

type DebitHoldEvent struct {
	AggregateID uuid.UUID
	Version     int32
	ParentID    uuid.UUID
	EventType   sql.NullString
	Payload     []byte
	CreatedAt   time.Time
}