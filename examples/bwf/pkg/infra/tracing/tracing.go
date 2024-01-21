package tracing

import (
	"context"
	"time"
)

var tracerInstance Tracer

func GetTracer() Tracer {
	return tracerInstance
}

func InitOrDie(
	tracer Tracer,
) Tracer {
	if tracer == nil {
		tracer = &NoopTracer{}
	}

	tracerInstance = tracer

	return tracerInstance
}

type Tracer interface {
	ExtractTraceInfo(
		ctx context.Context,
	) (ver, tid, pid, rid, flg string)
	TraceRequest(
		ctx context.Context,
		method string,
		path string,
		query string,
		statusCode int,
		bodySize int,
		ip string,
		userAgent string,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields map[string]string,
	)
	TraceEvent(
		ctx context.Context,
		name string,
		key string,
		statusCode int,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields map[string]string,
	)
	TraceDependency(
		ctx context.Context,
		spanId string,
		dependencyType string,
		serviceName string,
		commandName string,
		success bool,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields map[string]string,
	)
	TraceException(
		ctx context.Context,
		err interface{},
		skip int,
		fields map[string]string,
	)
	TraceDependencyWithIds(
		tid string,
		rid string,
		spanId string,
		dependencyType string,
		serviceName string,
		commandName string,
		success bool,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields map[string]string,
	)
}
