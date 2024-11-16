package data

import (
	"context"
	"database/sql"
	"greenlight/internal/assert"
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
	DB *sql.DB
}

func (m PermissinModel) GetAllForUser(userID int64) (Permissions, error) {
	query := `select perm.code
	from permissions
	inner join user_permissions user_perm on user_perm.permission_id = perm.id
	inner join users u on user_perm.user_id = u.id
	where u.id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var permissions Permissions

	for rows.Next() {
		var permission string
		if err := rows.Scan(&permission); err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
