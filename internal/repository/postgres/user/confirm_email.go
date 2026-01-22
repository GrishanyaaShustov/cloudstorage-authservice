// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import (
	"context"
	"fmt"
)

const confirmEmailQuery = `
	UPDATE users
	SET email_verified_at = now()
	WHERE id = $1
	  AND email_verified_at IS NULL;
`

// ConfirmEmail marks the user's email as verified.
func (r *UserRepository) ConfirmEmail(ctx context.Context, userID string) error {
	cmd, err := r.pool.Exec(ctx, confirmEmailQuery, userID)
	if err != nil {
		if isUnavailable(err) {
			return ErrUnavailable
		}
		return fmt.Errorf("%w: confirm email", ErrInternal)
	}

	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
