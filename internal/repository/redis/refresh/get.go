// Package refreshrepo provides Redis-backed storage for refresh sessions.
package refreshrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Get returns the stored value for a refresh session (e.g. session ID / jti).
func (r *RefreshRepository) Get(ctx context.Context, userID, app, device string) (string, error) {
	key := makeKey(userID, app, device)

	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrNotFound
		}
		if isUnavailable(err) {
			return "", ErrUnavailable
		}
		return "", fmt.Errorf("%w: get refresh session", ErrInternal)
	}

	return val, nil
}
