// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import (
	"context"
	"fmt"

	contract "github.com/GrishanyaaShustov/cloudstorage-authservice/internal/repository/user"
)

const existsByEmailQuery = `
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE email = $1
);
`

// ExistsByEmail checks whether a user with the specified email exists.
func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(ctx, existsByEmailQuery, email).Scan(&exists)
	if err != nil {
		if isUnavailable(err) {
			return false, contract.ErrUnavailable
		}
		return false, fmt.Errorf("%w: exists by email", contract.ErrInternal)
	}

	return exists, nil
}
