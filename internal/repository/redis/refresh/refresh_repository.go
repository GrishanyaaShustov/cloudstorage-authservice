// Package refreshrepo provides Redis-backed storage for refresh sessions.
package refreshrepo

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

const keyPrefix = "refresh"

// RefreshRepository provides access to refresh session persistence operations.
//
// It stores refresh sessions in Redis using a composite key built from
// user ID, application identifier, and device identifier.
type RefreshRepository struct {
	// rdb is the Redis client used for all repository operations.
	rdb *redis.Client
}

// New creates a new RefreshRepository instance.
func New(rdb *redis.Client) *RefreshRepository {
	return &RefreshRepository{rdb: rdb}
}

// makeKey builds a Redis key for a refresh session.
func makeKey(userID, app, device string) string {
	return fmt.Sprintf("%s:%s:%s:%s", keyPrefix, userID, app, device)
}
