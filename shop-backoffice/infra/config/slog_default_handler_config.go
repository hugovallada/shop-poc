package config

import (
	"log/slog"
	"os"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/utils/shared"
)

func SetDefaultSlogHandler() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})
	keys := []any{shared.CORRELATION_ID, shared.TRACE_ID}
	ctxHandler := NewCorrelationMultiKey(keys, handler)
	logger := slog.New(ctxHandler)
	slog.SetDefault(logger)
}
