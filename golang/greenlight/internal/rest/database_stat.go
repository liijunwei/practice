package rest

import (
	"database/sql"
	"greenlight/internal/common"
	"greenlight/internal/sse"
	"net/http"
	"time"
)

func DatabaseStatStream(db *sql.DB) sse.TypedHandler[sql.DBStats] {
	return func(req *http.Request) (*sse.Response[sql.DBStats], error) {
		ctx := req.Context()
		dataCh, errorCh := common.Tick(ctx, time.Second, func() (sql.DBStats, error) {
			return db.Stats(), nil
		})

		return &sse.Response[sql.DBStats]{
			Name:    "database_statistics",
			DataCh:  dataCh,
			ErrorCh: errorCh,
		}, nil
	}
}
