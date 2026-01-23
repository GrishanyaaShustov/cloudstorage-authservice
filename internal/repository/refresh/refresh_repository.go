package refreshrepo

import (
	"context"
	"errors"
	"time"
)

// RefreshRepository defines an interface for managing refresh sessions.
//
// Implementations are responsible for storing, validating, and removing
// refresh session records associated with a user, application, and device.
type RefreshRepository interface {
	// Save stores a refresh session with the provided TTL.
	Save(ctx context.Context, userID, app, device, value string, ttl time.Duration) error

	// Get returns the stored value for a refresh session.
	Get(ctx context.Context, userID, app, device string) (string, error)

	// Exists checks whether a refresh session exists.
	Exists(ctx context.Context, userID, app, device string) (bool, error)

	// Delete removes a refresh session.
	Delete(ctx context.Context, userID, app, device string) error
}

var (
	// ErrNotFound indicates that the requested session does not exist in storage.
	ErrNotFound = errors.New("repository: not found")

	// ErrUnavailable indicates that the repository backend is unavailable.
	ErrUnavailable = errors.New("repository: unavailable")

	// ErrInternal indicates an unexpected internal repository error.
	ErrInternal = errors.New("repository: internal")
)
