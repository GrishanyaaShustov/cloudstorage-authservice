// Package refreshrepo provides Redis-backed storage for refresh sessions.
package refreshrepo

import (
	"context"
	"errors"
)

var (
	// ErrNotFound indicates that the requested session does not exist in storage.
	ErrNotFound = errors.New("repository: not found")

	// ErrUnavailable indicates that the repository backend is unavailable.
	ErrUnavailable = errors.New("repository: unavailable")

	// ErrInternal indicates an unexpected internal repository error.
	ErrInternal = errors.New("repository: internal")
)

func isUnavailable(err error) bool {
	return errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled)
}
