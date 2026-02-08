package zlog

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Object map[string]any

type LogLevel uint8

const (
	LogLevelError LogLevel = iota
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
)

var (
	logLevel LogLevel = LogLevelDebug
)

func Setup() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &log.Logger
}

func SetLevel(level LogLevel) {
	logLevel = level
	switch level {
	case LogLevelError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case LogLevelWarn:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
