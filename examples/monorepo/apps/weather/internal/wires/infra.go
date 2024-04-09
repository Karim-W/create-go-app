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

// SetupInfra initializes the infra package.
// It is called by the main package.
func SetupInfra(
	opts InfraOptions,
) (InfraResults, error) {
	res := InfraResults{}

	// ================ Logger
	res.Logger = logger.InitOrDie()

	// ================ Tracing
	res.Trx = tracing.InitOrDie(nil)

	return res, nil
}
