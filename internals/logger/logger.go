package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

type Logger struct {
	Prefix string
}

func NewLogger(prefix string) *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: prefix,
	})
	return logger
}
