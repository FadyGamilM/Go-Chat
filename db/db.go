package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// i will encapsulate the db_pool inside this struct and i defined this struct to allow methods into at as a reciever to deal with the database conn pool from outside packages
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
