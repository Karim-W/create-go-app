package wires

import "{{.moduleName}}/services/factory"

// SetupServices initializes the services package.
// It is called by the main package.
// Extendable by adding more to the functions parameter list.
// and adding the return type to the return statement.

type ServiceOptions struct{}

type ServiceResults struct{}

func (r *ServiceResults) Close() error {
	// Close all the services

	return nil
}

func SetupServices(
	opts ServiceOptions,
) (res ServiceResults, err error) {
	factory.Init(
		factory.Dependencies{},
	)

	return
}
