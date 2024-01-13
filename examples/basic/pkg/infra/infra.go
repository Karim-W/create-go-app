package infra

import (
	"{{.moduleName}}/pkg/infra/logger"
	"{{.moduleName}}/pkg/infra/tracing"
)

type Options struct {
	Trx tracing.Tracer
}

type Results struct {
	Trx tracing.Tracer
}

// SetupInfra initializes the infra package.
// It is called by the main package.
func SetupInfra(
	opts *Options,
) (*Results, error) {
	res := &Results{}
	// init the logger
	logger.InitOrDie()
	// init the tracer
	res.Trx = tracing.InitOrDie(opts.Trx)
	// TODO: Add your other infra packages here
	return res, nil
}
