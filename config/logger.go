package config

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Log zerolog.Logger

func InitiateLog() {
	if LogMode == "DEBUG" {
		Log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Logger()
	} else {
		Log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	}
}
