package db

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

type AppDB struct {
	// private access to the `db` package only
	db_pool *sql.DB
}

func Connect() (*AppDB, error) {
	db, err := sql.Open("pgx", "postgresql://go_chat:go_chat@localhost:1234/go_chat_db?sslmode=false")
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
	// i named the params here because i need the users to know what they should pass to this func
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	PrepareContext(ctx context.Context, query string) (*sql.Result, error)

	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Result, error)

	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
