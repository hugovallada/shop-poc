package config

import (
	"context"
	"log/slog"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/utils/shared"
)

type CorrelationSlogContextHandler struct {
	slog.Handler
}

func (h CorrelationSlogContextHandler) Handle(ctx context.Context, record slog.Record) error {
	if correlationId, ok := ctx.Value(shared.CORRELATION_ID).(string); ok {
		record.Add("correlation", slog.StringValue(correlationId))
	}
	return h.Handler.Handle(ctx, record)
}
