package logger

import (
	"strings"

	"go.uber.org/zap"
)

// logginInstance is the logger instance scoped to this package.
// using zap as the logger.
var logginInstance *zap.Logger

// GetLogger returns the logger instance.
func GetLogger() *zap.Logger {
	return logginInstance
}

// InitOrDie initializes the logger and panics if it fails.
func InitOrDie() *zap.Logger {
	var err error
	logginInstance, err = zap.NewProduction(
		zap.AddCaller(),
	)
	if err != nil {
		panic(err)
	}
	if logginInstance == nil {
		panic("Failed to initialize logger")
	}
	return logginInstance
}

// GetTracedLogger returns a logger with the trace info.
// It is used to log the trace info in the logs.
// Will return the default logger if the trace info is invalid.
// example of trace info: "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
func GetTracedLogger(
	W3CTraceContext string,
) *zap.Logger {
	traceInfo := strings.Split(W3CTraceContext, "-")
	if len(traceInfo) != 4 {
		return logginInstance
	}
	return logginInstance.With(
		zap.String("tid", traceInfo[1]),
		zap.String("pid", traceInfo[2]),
		zap.String("rid", traceInfo[3]),
	)
}
