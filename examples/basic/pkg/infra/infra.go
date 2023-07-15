package infra

import (
	"{{.moduleName}}/pkg/infra/logger"
	"{{.moduleName}}/pkg/infra/tracing"

	"go.uber.org/zap"
)

type Options struct {
	Trx tracing.Tracer
}

type Results struct {
	Logger *zap.Logger
	Trx    tracing.Tracer
}

// SetupInfra initializes the infra package.
// It is called by the main package.
func SetupInfra(
	opts *Options,
) (*Results, error) {
	// init the logger
	l := logger.InitOrDie()
	// init the tracer
	trx := tracing.InitOrDie(opts.Trx)
	// TODO: Add your other infra packages here
	return &Results{
		Logger: l,
		Trx:    trx,
	}, nil
}
