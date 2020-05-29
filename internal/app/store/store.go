package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(conf *Config) *Store {
	return &Store{
		config: conf,
	}
}

func (store *Store) Open() error {
	db, err := sql.Open("postgres", store.config.dbURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db

	return nil
}

func (store *Store) Close() {
	store.db.Close()
}
