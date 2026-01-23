// Package refreshrepo provides Redis-backed storage for refresh sessions.
package refreshrepo

import (
	"context"
	"errors"
	"fmt"

	contract "github.com/GrishanyaaShustov/cloudstorage-authservice/internal/repository/refresh"
	"github.com/redis/go-redis/v9"
)

// Get returns the stored value for a refresh session (e.g. session ID / jti).
func (r *RefreshRepository) Get(ctx context.Context, userID, app, device string) (string, error) {
	key := makeKey(userID, app, device)

	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", contract.ErrNotFound
		}
		if isUnavailable(err) {
			return "", contract.ErrUnavailable
		}
		return "", fmt.Errorf("%w: get refresh session", contract.ErrInternal)
	}

	return val, nil
}
