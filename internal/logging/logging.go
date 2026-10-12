// Package logging creates structured diagnostics for stderr.
package logging

import (
	"io"
	"log/slog"
)

func New(w io.Writer, level slog.Leveler) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: level}))
}
