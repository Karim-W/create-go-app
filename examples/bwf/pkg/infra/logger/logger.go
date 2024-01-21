package logger

import (
	"go.uber.org/zap"
)

// InitOrDie initializes the logger and panics if it fails.
func InitOrDie() *zap.Logger {
	logginInstance, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logginInstance)
	return logginInstance
}
