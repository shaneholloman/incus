package logger

import (
	"github.com/sirupsen/logrus"
)

// Ctx is the logging context.
type Ctx logrus.Fields

// Log contains the logger used by all the logging functions.
var Log Logger

// Logger is the main logging interface.
type Logger interface {
	Panic(msg string, args ...Ctx)
	Fatal(msg string, args ...Ctx)
	Error(msg string, args ...Ctx)
	Warn(msg string, args ...Ctx)
	Info(msg string, args ...Ctx)
	Debug(msg string, args ...Ctx)
	Trace(msg string, args ...Ctx)
	AddContext(Ctx) Logger
}

// targetLogger represents the subset of logrus.Logger and logrus.Entry that we care about.
type targetLogger interface {
	Panic(args ...any)
	Fatal(args ...any)
	Error(args ...any)
	Warn(args ...any)
	Info(args ...any)
	Debug(args ...any)
	Trace(args ...any)
	WithFields(fields logrus.Fields) *logrus.Entry
}
