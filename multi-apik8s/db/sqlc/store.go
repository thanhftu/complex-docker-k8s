package db

import "database/sql"

// Store defines all functions to exe db queries
type Store interface {
	Querier
}

// SQLStore defines all functions to exe db queries
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
