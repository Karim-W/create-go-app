package adapters

import (
	"{{.moduleName}}/pkg/adapters/applicationinsightstrace"

	trace "github.com/BetaLixT/appInsightsTrace"
)

type Options struct {
	InstrumentationKey string
	ServiceName        string
}

type Results struct {
	Trx *trace.AppInsightsCore
}

// SetupAdapters initializes the adapters package.
// It is called by the main package.
// Extend this function to add your own adapters.
// Pass the adapters dependencies as parameters.
// Add your adapters to the function return type
func SetupAdapters(
	opts *Options,
) (*Results, error) {
	trx := applicationinsightstrace.InitOrDie(opts.InstrumentationKey, opts.ServiceName)
	return &Results{
		Trx: trx,
	}, nil
}
