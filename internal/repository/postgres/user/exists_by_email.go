package userrepo

import (
	"context"
	"fmt"
)

const existsByEmailQuery = `
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE email = $1
);
`

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(ctx, existsByEmailQuery, email).Scan(&exists)
	if err != nil {
		if isUnavailable(err) {
			return false, ErrUnavailable
		}
		return false, fmt.Errorf("%w: exists by email", ErrInternal)
	}

	return exists, nil
}
