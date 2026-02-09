package adaptersoutlog

import (
	"context"
	"log/slog"
	"os"

	portsoutlog "github.com/dhonanhibatullah/basuke/backend/internal/ports/out/log"
)

type Log struct {
	logger *slog.Logger
}

func New(level slog.Level) portsoutlog.Log {
	return &Log{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})),
	}
}

func (l *Log) Error(ctx context.Context, tag string, action string, message string, metadata map[string]any) {
	l.logger.ErrorContext(
		ctx,
		message,
		slog.String("tag", tag),
		slog.String("action", action),
		slog.Any("meta", metadata),
	)
}

func (l *Log) Warn(ctx context.Context, tag string, action string, message string, metadata map[string]any) {
	l.logger.WarnContext(
		ctx,
		message,
		slog.String("tag", tag),
		slog.String("action", action),
		slog.Any("meta", metadata),
	)
}

func (l *Log) Info(ctx context.Context, tag string, action string, message string, metadata map[string]any) {
	l.logger.InfoContext(
		ctx,
		message,
		slog.String("tag", tag),
		slog.String("action", action),
		slog.Any("meta", metadata),
	)
}

func (l *Log) Debug(ctx context.Context, tag string, action string, message string, metadata map[string]any) {
	l.logger.DebugContext(
		ctx,
		message,
		slog.String("tag", tag),
		slog.String("action", action),
		slog.Any("meta", metadata),
	)
}
