package sql_queries

import (
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db *sqlx.DB
}
