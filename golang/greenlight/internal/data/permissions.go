package data

import (
	"context"
	"greenlight/internal/assert"
	"greenlight/internal/sqlcdb"
	"time"
)

type Permissions []string

func (p Permissions) Include(code string) bool {
	assert.Assert(code != "")
	assert.Assert(len(p) > 0)

	for _, permissionCode := range p {
		if code == permissionCode {
			return true
		}
	}

	return false
}

type PermissinModel struct {
	queries *sqlcdb.Queries
}

func (m PermissinModel) GetAllForUser(ctx context.Context, userID int64) (Permissions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	permissions, err := m.queries.GetAllForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (m PermissinModel) AddForUser(ctx context.Context, userID int64, codes ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.queries.AddForUser(ctx, sqlcdb.AddForUserParams{
		UserID: userID,
		Codes:  codes,
	}); err != nil {
		return err
	}

	return nil
}
