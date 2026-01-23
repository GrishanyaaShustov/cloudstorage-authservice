// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import "github.com/jackc/pgx/v5/pgxpool"

// UserRepository provides access to user persistence operations.
//
// It encapsulates database access logic for working with users
// and is responsible for executing queries related to user data.
type UserRepository struct {
	// pool is the PostgreSQL connection pool used for all repository operations.
	pool *pgxpool.Pool
}

// New creates a new UserRepository instance.
func New(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}
