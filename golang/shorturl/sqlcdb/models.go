// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlcdb

import (
	"time"
)

type Shorturl struct {
	ID        int64
	Original  string
	Shorturl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
