package services

import (
	"{{.moduleName}}/pkg/infra/tracing"
	"{{.moduleName}}/pkg/services/factory"
)

// SetupServices initializes the services package.
// It is called by the main package.
// Extendable by adding more to the functions parameter list.
// and adding the return type to the return statement.

type Options struct {
	Trx tracing.Tracer
}

func SetupServices(
	opts *Options,
) error {
	factory.SetUpDependencies(opts.Trx)
	return nil
}
