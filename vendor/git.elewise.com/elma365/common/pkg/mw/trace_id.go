package mw

import (
	"context"

	"git.elewise.com/elma365/common/pkg/config"

	"go.uber.org/zap"
)

// TraceIDFromContext extract trace and span ids if span started or empty string otherwise
func TraceIDFromContext(ctx context.Context) (traceID, spanID string) {
	return config.TraceIDFromContext(ctx)
}

// LoggerWithTraceID добавляет поля TraceID и SpanID к полям логгера
func LoggerWithTraceID(ctx context.Context, logger *zap.Logger) *zap.Logger {
	return config.LoggerWithTraceID(ctx, logger)
}
