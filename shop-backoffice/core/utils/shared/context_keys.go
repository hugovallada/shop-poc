package shared

type CorrelationId string

type TraceId string

const (
	CORRELATION_ID CorrelationId = "correlationId"
	TRACE_ID       TraceId       = "traceId"
)
