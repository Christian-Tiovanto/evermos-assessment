// Package log is to handle logging operation
package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is interface to logger
type Logger interface {
	Info(string, ...zapcore.Field)
	Warning(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
	With(...zapcore.Field) Logger
	WithContext(context.Context) Logger
}

// Logs is the struct that implements Logger
type Logs struct {
	fields []zapcore.Field
	Logger *zap.Logger
}

// NewLogger initialize logger
func NewLogger() *Logs {
	logger := zap.Must(zapConfig(true).Build())
	return &Logs{
		Logger: logger,
		fields: make([]zapcore.Field, 0),
	}
}

// Info level log
func (l *Logs) Info(message string, fields ...zapcore.Field) {
	l.Logger.Info(message,
		append([]zapcore.Field{zap.String("severity", "INFO")}, fields...)...,
	)
}

// Warning level log
func (l *Logs) Warning(message string, fields ...zapcore.Field) {
	l.Logger.Warn(message,
		append([]zapcore.Field{zap.String("severity", "WARNING")}, fields...)...,
	)
}

// Error level log
func (l *Logs) Error(message string, fields ...zapcore.Field) {
	l.Logger.Error(message,
		append([]zapcore.Field{zap.String("severity", "ERROR")}, fields...)...,
	)
}

// With append logger field and return new Logger
func (l *Logs) With(fields ...zapcore.Field) Logger {
	return &Logs{
		fields: append(l.fields, fields...),
		Logger: l.Logger,
	}
}

func zapConfig(isSamplingEnabled bool) zap.Config {
	config := zap.NewProductionConfig()
	if !isSamplingEnabled {
		config.Sampling = nil
	}

	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       zapcore.OmitKey,
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      zapcore.OmitKey,
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
	}
	return config
}
