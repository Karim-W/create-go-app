package tracing

import (
	"context"
	"log"
	"time"
)

type testing struct{}

func (n *testing) Close() {
}

func (n *testing) ExtractTraceInfo(
	ctx context.Context,
) (ver string, tid string, pid string, rid string, flg string) {
	return "", "", "", "", ""
}

func (n *testing) TraceDependency(
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
	log.Printf("spanId: %s, dependencyType: %s, serviceName: %s, commandName: %s, success: %t, startTimestamp: %s, eventTimestamp: %s, fields: %v\n", spanId, dependencyType, serviceName, commandName, success, startTimestamp, eventTimestamp, fields)
}

func (n *testing) TraceDependencyWithIds(
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
	log.Printf("tid: %s, rid: %s, spanId: %s, dependencyType: %s, serviceName: %s, commandName: %s, success: %t, startTimestamp: %s, eventTimestamp: %s, fields: %v\n", tid, rid, spanId, dependencyType, serviceName, commandName, success, startTimestamp, eventTimestamp, fields)
}

func (n *testing) TraceEvent(
	ctx context.Context,
	name string,
	key string,
	statusCode int,
	startTimestamp time.Time,
	eventTimestamp time.Time,
	fields map[string]string,
) {
	log.Printf("name: %s, key: %s, statusCode: %d, startTimestamp: %s, eventTimestamp: %s, fields: %v\n", name, key, statusCode, startTimestamp, eventTimestamp, fields)
}

func (n *testing) TraceException(
	ctx context.Context,
	err interface{},
	skip int,
	fields map[string]string,
) {
	log.Printf("err: %v, skip: %d, fields: %v\n", err, skip, fields)
}

func (n *testing) TraceRequest(
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
	log.Printf("method: %s, path: %s, query: %s, statusCode: %d, bodySize: %d, ip: %s, userAgent: %s, startTimestamp: %s, eventTimestamp: %s, fields: %v\n", method, path, query, statusCode, bodySize, ip, userAgent, startTimestamp, eventTimestamp, fields)
}

func Testing() Tracer {
	return &testing{}
}
