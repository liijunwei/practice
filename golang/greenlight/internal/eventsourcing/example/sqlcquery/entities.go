// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlcquery

import (
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
