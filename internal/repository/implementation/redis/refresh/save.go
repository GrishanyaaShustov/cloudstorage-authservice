// Package refreshrepo provides Redis-backed storage for refresh sessions.
package refreshrepo

import (
	"context"
	"fmt"
	"time"

	contract "github.com/GrishanyaaShustov/cloudstorage-authservice/internal/repository/refresh"
)

// Save stores a refresh session with the provided TTL.
func (r *RefreshRepository) Save(ctx context.Context, userID, app, device, value string, ttl time.Duration) error {
	key := makeKey(userID, app, device)

	if err := r.rdb.Set(ctx, key, value, ttl).Err(); err != nil {
		if isUnavailable(err) {
			return contract.ErrUnavailable
		}
		return fmt.Errorf("%w: save refresh session", contract.ErrInternal)
	}

	return nil
}
