package logs

import (
	"log/slog"
	"os"

	"github.com/hugovallada/correlationcontexthandler"
)

func SetDefaultSlogHandler() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})
	keys := []any{correlationcontexthandler.CORRELATION_ID, correlationcontexthandler.TRACE_ID}
	ctxHandler := correlationcontexthandler.NewMultiKeyContextHandler(keys, handler)
	logger := slog.New(ctxHandler)
	slog.SetDefault(logger)
}
