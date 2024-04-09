package redistracer

import (
	"context"
	"time"
)

type Tracer interface {
	ExtractTraceInfo(
		ctx context.Context,
	) (ver, tid, pid, rid, flg string)
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
}
