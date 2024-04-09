package repositories

import (
	"{{.moduleName}}/pkg/domains/weather"
	"{{.moduleName}}/services/factory"
)

type Forecast interface {
	List(
		ftx factory.Service,
	) ([]weather.Entity, error)
}
