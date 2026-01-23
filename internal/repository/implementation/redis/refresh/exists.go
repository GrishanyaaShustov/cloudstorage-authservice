// Package refreshrepo provides Redis-backed storage for refresh sessions.
package refreshrepo

import (
	"context"
	"fmt"

	contract "github.com/GrishanyaaShustov/cloudstorage-authservice/internal/repository/refresh"
)

// Exists checks whether a refresh session exists for the provided composite key.
func (r *RefreshRepository) Exists(ctx context.Context, userID, app, device string) (bool, error) {
	key := makeKey(userID, app, device)

	n, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		if isUnavailable(err) {
			return false, contract.ErrUnavailable
		}
		return false, fmt.Errorf("%w: exists refresh session", contract.ErrInternal)
	}

	return n == 1, nil
}
