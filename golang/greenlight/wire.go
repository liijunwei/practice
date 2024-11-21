//go:build wireinject
// +build wireinject

package greenlight

import (
	"database/sql"
	"greenlight/internal/config"
	"greenlight/internal/data"
	"greenlight/internal/dbconn"
	"greenlight/internal/sqlcdb"

	"github.com/google/wire"
)

func SetupModels(cfg config.Config) (data.Models, error) {
	wire.Build(
		data.NewModels,
		sqlcdb.New,
		wire.Bind(new(sqlcdb.DBTX), new(*sql.DB)),
		dbconn.NewDB,
	)

	return data.Models{}, nil
}
