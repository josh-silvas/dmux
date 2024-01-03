package nlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// LogLevel is a global variable that can be used to set the log level
var LogLevel = &slog.LevelVar{}

// Logger : custom logger
type Logger struct {
	l     *slog.Logger
	ctx   context.Context
	level *slog.LevelVar
}

// New : create a new logger
func New() Logger {
	return Logger{
		l:     slog.New(cliHandler(&slog.HandlerOptions{Level: LogLevel})),
		ctx:   context.Background(),
		level: LogLevel,
	}
}

// NewWithGroup : create a new logger with a group
func NewWithGroup(name string) Logger {
	return Logger{
		l:     slog.New(cliHandler(&slog.HandlerOptions{Level: LogLevel})).WithGroup(name),
		ctx:   context.Background(),
		level: LogLevel,
	}
}

// Level : return the current log level
func (l Logger) Level() slog.Level {
	return l.level.Level()
}

// Error : log Error message
func (l Logger) Error(args ...any) {
	l.l.Log(l.ctx, slog.LevelError, fmt.Sprint(args...))
}

// Errorf : log Error message with fmt.Sprintf
func (l Logger) Errorf(format string, args ...any) {
	l.l.Log(l.ctx, slog.LevelError, fmt.Sprintf(format, args...))
}

// Fatal : log Fatal message
func (l Logger) Fatal(args ...any) {
	l.l.Log(l.ctx, LevelFatal, fmt.Sprint(args...))
	os.Exit(1)
}

// Fatalf : log Fatal message with fmt.Sprintf
func (l Logger) Fatalf(format string, args ...any) {
	l.l.Log(l.ctx, LevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}

// Warn : log Warn message
func (l Logger) Warn(args ...any) {
	l.l.Log(l.ctx, slog.LevelWarn, fmt.Sprint(args...))
}

// Warnf : log warn message with fmt.Sprintf
func (l Logger) Warnf(format string, args ...any) {
	l.l.Log(l.ctx, slog.LevelWarn, fmt.Sprintf(format, args...))
}

// Info : log info message
func (l Logger) Info(args ...any) {
	l.l.Log(l.ctx, slog.LevelInfo, fmt.Sprint(args...))
}

// Infof : log info message with fmt.Sprintf
func (l Logger) Infof(format string, args ...any) {
	l.l.Log(l.ctx, slog.LevelInfo, fmt.Sprintf(format, args...))
}

// Debug : log Debug message
func (l Logger) Debug(args ...any) {
	l.l.Log(l.ctx, slog.LevelDebug, fmt.Sprint(args...))
}

// Debugf : log Debug message with fmt.Sprintf
func (l Logger) Debugf(format string, args ...any) {
	l.l.Log(l.ctx, slog.LevelDebug, fmt.Sprintf(format, args...))
}
