// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import "github.com/jackc/pgx/v5/pgxpool"

// Repository provides access to user persistence operations.
//
// It encapsulates database access logic for working with users
// and is responsible for executing queries related to user data.
type Repository struct {
	// pool is the PostgreSQL connection pool used for all repository operations.
	pool *pgxpool.Pool
}

// New creates a new Repository instance.
func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}
