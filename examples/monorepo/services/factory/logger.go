package factory

import "go.uber.org/zap"

// Logger() Returns the logger with the traceinfo
func (s *sf) Logger() *zap.Logger {
	return s.logger
}
