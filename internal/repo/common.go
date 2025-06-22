package repo

import (
	"github.com/jackc/pgx"
)

// Repository general package struct
type Repository struct {
	db *pgx.Conn
}

// New repository constructor
func New(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}
