package wires

import (
	"{{.moduleName}}/pkg/infrastructure/logger"
	"{{.moduleName}}/pkg/infrastructure/tracing"

	"go.uber.org/zap"
)

type InfraOptions struct{}

type InfraResults struct {
	Logger *zap.Logger
	Trx    tracing.Tracer
}

func (r *InfraResults) Close() error {
	r.Trx.Close()

	r.Logger.Sync()

	return nil
}

// SetupInfra initializes the infra package.
// It is called by the main package.
func SetupInfra(
	opts InfraOptions,
) (res InfraResults, err error) {
	// ================ Logger
	res.Logger = logger.InitOrDie()

	// ================ Tracing
	res.Trx = tracing.InitOrDie(nil)

	return
}
