// Package sl provides small helpers for working with slog attributes.
package sl

import (
	"log/slog"
)

// Err converts an error into a slog attribute.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
