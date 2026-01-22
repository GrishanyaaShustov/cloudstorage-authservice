// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import (
	"context"
	"fmt"
)

const createUserQuery = `
	INSERT INTO users (email, login, password_hash)
	VALUES ($1, $2, $3);
`

// Create inserts a new user record into the storage.
func (r *UserRepository) Create(ctx context.Context, email, login, passwordHash string) error {
	_, err := r.pool.Exec(ctx, createUserQuery, email, login, passwordHash)
	if err != nil {
		if isUniqueViolation(err) {
			return mapUniqueViolation(err)
		}
		if isUnavailable(err) {
			return ErrUnavailable
		}
		return fmt.Errorf("%w: create user", ErrInternal)
	}

	return nil
}
