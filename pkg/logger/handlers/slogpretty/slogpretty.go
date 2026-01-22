// Package slogpretty provides a human-friendly slog handler implementation
// that prints colored log records and pretty-formatted attributes.
package slogpretty

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"

	"github.com/fatih/color"
)

// PrettyHandlerOptions describes configuration options for PrettyHandler.
//
// It allows customizing the underlying slog handler behavior by providing
// slog.HandlerOptions used to build a JSON handler that PrettyHandler wraps.
type PrettyHandlerOptions struct {
	// SlogOpts defines options used to configure the underlying slog handler.
	// These options control level filtering, attribute replacement, and other behavior.
	SlogOpts *slog.HandlerOptions
}

// PrettyHandler implements slog.Handler and renders log records in a readable form.
//
// It wraps a JSON slog handler to preserve slog semantics while additionally
// printing colored level and message output along with pretty-formatted attributes.
type PrettyHandler struct {
	// opts stores handler configuration used to initialize the underlying slog handler.
	opts PrettyHandlerOptions

	// Handler is the underlying slog handler used for grouping and attribute semantics.
	slog.Handler

	// l is the standard logger used to write formatted output.
	l *stdLog.Logger

	// attrs contains accumulated attributes added via WithAttrs.
	attrs []slog.Attr
}

// NewPrettyHandler constructs a PrettyHandler writing logs to the provided writer.
func (opts PrettyHandlerOptions) NewPrettyHandler(out io.Writer) *PrettyHandler {
	h := &PrettyHandler{
		opts:    opts,
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

// Handle formats and writes a log record using colored level and message output.
func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

// WithAttrs returns a new handler instance with the provided attributes attached.
func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	merged := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	merged = append(merged, h.attrs...)
	merged = append(merged, attrs...)

	return &PrettyHandler{
		opts:    h.opts,
		Handler: h.Handler,
		l:       h.l,
		attrs:   merged,
	}
}

// WithGroup returns a new handler instance with the provided group applied.
func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	return &PrettyHandler{
		opts:    h.opts,
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
		attrs:   h.attrs,
	}
}
