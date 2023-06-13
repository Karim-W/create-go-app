package infra

import (
	"{{.moduleName}}/infra/logger"
	"{{.moduleName}}/infra/tracing"
	"go.uber.org/zap"
)

// SetupInfra initializes the infra package.
// It is called by the main package.
func SetupInfra(
	tracer tracing.Tracer,
) (*zap.Logger, tracing.Tracer, error) {
	// init the logger
	l := logger.InitOrDie()
	// init the tracer
	trx := tracing.InitOrDie(tracer)
	// TODO: Add your other infra packages here
	return l, trx, nil
}
