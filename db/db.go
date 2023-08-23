package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type AppDB struct {
	// private access to the `db` package only
	db_pool *sql.DB
}

func Connect() (*AppDB, error) {
	db, err := sql.Open("postgres", "postgresql://go_chat:go_chat@localhost:1234/go_chat_db?sslmode=false")
	if err != nil {
		return nil, err
	}
	return &AppDB{db_pool: db}, nil
}

func (db *AppDB) Close() {
	db.db_pool.Close()
}

func (db *AppDB) GetDbPool() *sql.DB {
	return db.db_pool
}

// this interface type is used so any repo implements it can receive a transaction "tx" or the database pool itself "db"
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)

	PrepareContext(context.Context, string) (*sql.Result, error)

	QueryContext(context.Context, string, ...interface{}) (*sql.Result, error)

	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
