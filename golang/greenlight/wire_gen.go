// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package greenlight

import (
	"greenlight/internal/config"
	"greenlight/internal/data"
	"greenlight/internal/dbconn"
	"greenlight/internal/sqlcdb"
)

// Injectors from wire.go:

func SetupModels(cfg config.Config) (data.Models, error) {
	db, err := dbconn.NewDB(cfg)
	if err != nil {
		return data.Models{}, err
	}
	queries := sqlcdb.New(db)
	models := data.NewModels(queries)
	return models, nil
}
