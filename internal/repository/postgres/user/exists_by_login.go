// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import (
	"context"
	"fmt"
)

const existsByLoginQuery = `
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE login = $1
);
`

// ExistsByLogin checks whether a user with the specified login exists.
func (r *UserRepository) ExistsByLogin(ctx context.Context, login string) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(ctx, existsByLoginQuery, login).Scan(&exists)
	if err != nil {
		if isUnavailable(err) {
			return false, ErrUnavailable
		}
		return false, fmt.Errorf("%w: exists by login", ErrInternal)
	}

	return exists, nil
}
