// Package slogdiscard provides a slog handler implementation
// that discards all log records.
package slogdiscard

import (
	"context"

	"golang.org/x/exp/slog"
)

// DiscardHandler implements slog.Handler and ignores all log records.
//
// It is intended to be used in environments where logging must be
// completely disabled, such as tests or silent runtime modes.
type DiscardHandler struct{}

// NewDiscardLogger returns a slog.Logger instance
// configured with a discard handler.
func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

// NewDiscardHandler creates a new DiscardHandler instance.
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// Handle ignores the provided log record.
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs returns the same handler since attributes are ignored.
func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup returns the same handler since groups are ignored.
func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}

// Enabled always returns false to prevent log records from being processed.
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
