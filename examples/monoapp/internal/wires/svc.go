package wires

import "{{.moduleName}}/services/factory"

// SetupServices initializes the services package.
// It is called by the main package.
// Extendable by adding more to the functions parameter list.
// and adding the return type to the return statement.

type ServiceOptions struct{}

func SetupServices(
	opts ServiceOptions,
) error {
	factory.Init(
		factory.Dependencies{},
	)
	return nil
}
