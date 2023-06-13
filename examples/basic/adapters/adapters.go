package adapters

import (
	trace "github.com/BetaLixT/appInsightsTrace"
	"{{.moduleName}}/adapters/applicationinsightstrace"
)

// SetupAdapters initializes the adapters package.
// It is called by the main package.
// Extend this function to add your own adapters.
// Pass the adapters dependencies as parameters.
// Add your adapters to the function return type
func SetupAdapters(
	instrumentationKey string,
	serviceName string,
) (*trace.AppInsightsCore, error) {
	trx := applicationinsightstrace.InitOrDie(instrumentationKey, serviceName)
	return trx, nil
}
