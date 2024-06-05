package logs

import (
	"log/slog"
	"os"

	"github.com/hugovallada/correlationcontexthandler"
	"github.com/spf13/viper"
)

func SetDefaultSlogHandler() {
	addSource := viper.GetBool("log.add_source")
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: addSource,
	})
	keys := []any{correlationcontexthandler.CORRELATION_ID, correlationcontexthandler.TRACE_ID, correlationcontexthandler.FLOW_ID}
	ctxHandler := correlationcontexthandler.NewMultiKeyContextHandler(keys, handler)
	logger := slog.New(ctxHandler)
	slog.SetDefault(logger)
}
