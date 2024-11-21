//go:build wireinject
// +build wireinject

package data

import (
	"database/sql"
	"greenlight/internal/config"
	"greenlight/internal/sqlcdb"

	"github.com/google/wire"
)

//go:generate wire
func SetupModels(cfg config.Config, db *sql.DB) (Models, error) {
	wire.Build(
		NewModels,
		sqlcdb.New,
		wire.Bind(new(sqlcdb.DBTX), new(*sql.DB)),
	)

	return Models{}, nil
}
