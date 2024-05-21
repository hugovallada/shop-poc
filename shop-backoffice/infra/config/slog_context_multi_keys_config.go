package config

import (
	"context"
	"fmt"
	"log/slog"
)

type CorrelationSlogContextMultiKeysHandler struct {
	slog.Handler
	keysToLog []any
}

func NewCorrelationMultiKey(keys []any, handler slog.Handler) CorrelationSlogContextMultiKeysHandler {
	return CorrelationSlogContextMultiKeysHandler{
		handler,
		keys,
	}
}

func (h CorrelationSlogContextMultiKeysHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, keyToLog := range h.keysToLog {
		if value, ok := ctx.Value(keyToLog).(string); ok {
			record.Add(fmt.Sprintf("%v", keyToLog), slog.StringValue(value))
		}
	}
	return h.Handler.Handle(ctx, record)
}
