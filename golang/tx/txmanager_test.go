package tx

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
	"github.com/peterldowns/pgtestdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go clean -testcache && go test ./tx -v

func TestManager(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := t.Context()
		conn := prepareDB(t)
		txmanager := New(conn)

		updateFunc := func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.ExecContext(ctx, "insert into t1 (name) values ('t1')")
			require.NoError(t, err)

			_, err = tx.ExecContext(ctx, "insert into t1 (name) values ('t2')")
			require.NoError(t, err)

			return nil
		}

		err := txmanager.WithinTransaction(ctx, updateFunc)
		require.NoError(t, err)

		verifyCount(t, conn, 2)
	})

	t.Run("rollback on error", func(t *testing.T) {
		t.Parallel()

		ctx := t.Context()
		conn := prepareDB(t)
		txmanager := New(conn)

		updateFunc := func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.ExecContext(ctx, "insert into t1 (name) values ('t1')")
			require.NoError(t, err)

			return errors.New("simulated error")
		}

		err := txmanager.WithinTransaction(ctx, updateFunc)
		require.Error(t, err)
		spew.Dump(err.Error())
		verifyCount(t, conn, 0)
	})
}

func Test_pgtestdb(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	conn := prepareDB(t)

	_, err := conn.ExecContext(ctx, "insert into t1 (name) values ('t1')")
	require.NoError(t, err)
	_, err = conn.ExecContext(ctx, "insert into t1 (name) values ('t2')")
	require.NoError(t, err)

	verifyCount(t, conn, 2)
}

func prepareDB(t *testing.T) *sql.DB {
	t.Helper()

	ctx := t.Context()
	conn := newDB(t)

	_, err := conn.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS t1 (id SERIAL PRIMARY KEY, name TEXT)")
	require.NoError(t, err)

	return conn
}

// https://github.com/peterldowns/pgtestdb
func newDB(t *testing.T) *sql.DB {
	t.Helper()

	// 	dsn := "postgres://admin:123@192.168.31.143:5432/demo?sslmode=disable&application_name=demo-user"
	conf := pgtestdb.Config{
		DriverName: "postgres",
		User:       "admin",
		Password:   "123",
		Host:       "192.168.31.143",
		Port:       "5432",
		Options:    "sslmode=disable",
		Database:   "demo",
	}

	var migrator pgtestdb.Migrator = pgtestdb.NoopMigrator{}
	db := pgtestdb.New(t, conf, migrator)

	return db
}

func verifyCount(t *testing.T, conn *sql.DB, expectedCount int) {
	t.Helper()

	ctx := t.Context()

	c, err := conn.QueryContext(ctx, "select count(*) from t1")
	require.NoError(t, err)

	for c.Next() {
		var count int
		require.NoError(t, c.Scan(&count))
		assert.Equal(t, expectedCount, count)
	}
}
