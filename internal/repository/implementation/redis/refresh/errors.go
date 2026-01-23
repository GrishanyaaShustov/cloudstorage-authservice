// Package refreshrepo provides Redis-backed storage for refresh sessions.
package refreshrepo

import (
	"context"
	"errors"
)

func isUnavailable(err error) bool {
	return errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled)
}
