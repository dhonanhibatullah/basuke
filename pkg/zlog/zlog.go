package zlog

import (
	"context"

	"github.com/rs/zerolog/log"
)

func Error(ctx context.Context, tag string, action string, message string, data any) {
	zPrint(ctx, LogLevelError, tag, action, message, data)
}

func Warn(ctx context.Context, tag string, action string, message string, data any) {
	zPrint(ctx, LogLevelWarn, tag, action, message, data)
}

func Info(ctx context.Context, tag string, action string, message string, data any) {
	zPrint(ctx, LogLevelInfo, tag, action, message, data)
}

func Debug(ctx context.Context, tag string, action string, message string, data any) {
	zPrint(ctx, LogLevelDebug, tag, action, message, data)
}

func zPrint(ctx context.Context, level LogLevel, tag string, action string, message string, data any) {
	switch level {
	case LogLevelError:
		log.Ctx(ctx).
			Error().
			Str("tag", tag).
			Str("action", action).
			Interface("data", data).
			Msg(message)

	case LogLevelWarn:
		log.Ctx(ctx).
			Warn().
			Str("tag", tag).
			Str("action", action).
			Interface("data", data).
			Msg(message)

	case LogLevelInfo:
		log.Ctx(ctx).
			Info().
			Str("tag", tag).
			Str("action", action).
			Interface("data", data).
			Msg(message)

	case LogLevelDebug:
		log.Ctx(ctx).
			Debug().
			Str("tag", tag).
			Str("action", action).
			Interface("data", data).
			Msg(message)
	}
}
