package factory

import "{{.moduleName}}/pkg/infra/tracing"

type depenencies struct {
	trx tracing.Tracer
}

var deps *depenencies

// not thread safe
func SetUpDependencies(trx tracing.Tracer) {
	if deps != nil {
		return
	}

	deps = &depenencies{trx}
}
