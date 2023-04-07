package logger

import (
	"context"

	"github.com/go-logr/logr"
)

const (
	debugLogLevel = 2
	warnLogLevel  = 3
	traceLogLevel = 4
)

type Wrapper interface {
	Info(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Trace(msg string, keysAndValues ...interface{})
	Error(err error, msg string, keysAndValues ...interface{})
	GetLogger() logr.Logger
}

type Logger struct {
	logr.Logger
}

func NewLogger(log logr.Logger) *Logger {
	return &Logger{Logger: log}
}

func FromContext(ctx context.Context) *Logger {
	log := logr.FromContextOrDiscard(ctx)
	return &Logger{Logger: log}
}

var _ Wrapper = &Logger{}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.Logger.Info(msg, keysAndValues...)
}

func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.Logger.V(debugLogLevel).Info(msg, keysAndValues...)
}

func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.Logger.V(warnLogLevel).Info(msg, keysAndValues...)
}

func (l *Logger) Trace(msg string, keysAndValues ...interface{}) {
	l.Logger.V(traceLogLevel).Info(msg, keysAndValues...)
}

func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.Logger.Error(err, msg, keysAndValues...)
}

func (l *Logger) GetLogger() logr.Logger {
	return l.Logger
}
