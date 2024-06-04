package logs

import (
	"log/slog"
	"os"

	"github.com/hugovallada/correlationcontexthandler"
)

func SetDefaultSlogHandler() {
	env := os.Getenv("ENVIRONMENT")
	var addSource bool = true
	if env == "local" {
		addSource = false
	}
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: addSource,
	})
	keys := []any{correlationcontexthandler.CORRELATION_ID, correlationcontexthandler.TRACE_ID, correlationcontexthandler.FLOW_ID}
	ctxHandler := correlationcontexthandler.NewMultiKeyContextHandler(keys, handler)
	logger := slog.New(ctxHandler)
	slog.SetDefault(logger)
}
