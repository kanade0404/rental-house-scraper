package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func Init() {
	zerolog.TimeFieldFormat = time.RFC3339
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Infof(format string, msgs ...interface{}) {
	log.Info().Msgf(format, msgs...)
}

func Fatal(msg string) {
	log.Fatal().Msg(msg)
}

func Fatalf(format string, msgs ...interface{}) {
	log.Fatal().Msgf(format, msgs...)
}

func Error(msg string) {
	log.Error().Msg(msg)
}

func Errorf(format string, msgs ...interface{}) {
	log.Error().Msgf(format, msgs...)
}
