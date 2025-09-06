package authpsql

import "github.com/atareversei/quardian/services/api/internal/repository/psql"

type DB struct {
	db *psql.DB
}

func New(db *psql.DB) *DB {
	return &DB{
		db: db,
	}
}
