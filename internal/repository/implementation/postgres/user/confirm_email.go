// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import (
	"context"
	"fmt"

	contract "github.com/GrishanyaaShustov/cloudstorage-authservice/internal/repository/user"
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
			return contract.ErrUnavailable
		}
		return fmt.Errorf("%w: confirm email", contract.ErrInternal)
	}

	if cmd.RowsAffected() == 0 {
		return contract.ErrNotFound
	}

	return nil
}
