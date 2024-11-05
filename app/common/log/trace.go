package log

import (
	"context"

	"go.uber.org/zap/zapcore"
)

type loggingCtx struct{}

// AppendContext append logging context into context
func AppendContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	return context.WithValue(ctx, loggingCtx{}, append(LoggingContext(ctx), fields...))
}

// LoggingContext retrieve custom logging data inside context
func LoggingContext(ctx context.Context) []zapcore.Field {
	fields, ok := ctx.Value(loggingCtx{}).([]zapcore.Field)
	if ok {
		return fields
	}

	return []zapcore.Field{}
}

// FromContext initialize new logger using context to get trace_id
func FromContext(ctx context.Context) Logger {
	return NewLogger().WithContext(ctx)
}

// WithContext decorate current logger with trace_id from context
func (l *Logs) WithContext(ctx context.Context) Logger {
	logger := &Logs{
		fields: l.fields,
		Logger: l.Logger,
	}
	return logger.With(LoggingContext(ctx)...)
}
