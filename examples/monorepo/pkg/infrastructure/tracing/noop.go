package tracing

import (
	"context"
	"time"
)

type noop struct{}

func (n *noop) ExtractTraceInfo(
	ctx context.Context,
) (ver string, tid string, pid string, rid string, flg string) {
	return "", "", "", "", ""
}

func (n *noop) TraceDependency(
	ctx context.Context,
	spanId string,
	dependencyType string,
	serviceName string,
	commandName string,
	success bool,
	startTimestamp time.Time,
	eventTimestamp time.Time,
	fields map[string]string,
) {
}

func (n *noop) TraceDependencyWithIds(
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
) {
}

func (n *noop) TraceEvent(
	ctx context.Context,
	name string,
	key string,
	statusCode int,
	startTimestamp time.Time,
	eventTimestamp time.Time,
	fields map[string]string,
) {
}

func (n *noop) TraceException(
	ctx context.Context,
	err interface{},
	skip int,
	fields map[string]string,
) {
}

func (n *noop) TraceRequest(
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
) {
}

func Noop() Tracer {
	return &noop{}
}
