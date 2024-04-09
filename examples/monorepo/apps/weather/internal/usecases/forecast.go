package usecases

import (
	"{{.moduleName}}/pkg/domains/listable"
	"{{.moduleName}}/pkg/domains/weather"
	"{{.moduleName}}/services/factory"
)

type Forecast interface {
	// Get() returns a list of weather Forecasts
	Get(
		ftx factory.Service,
	) (listable.QueryList[weather.Entity], error)
}
