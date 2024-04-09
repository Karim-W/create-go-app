package forecastusecase

import (
	"{{.moduleName}}/pkg/domains/listable"
	"{{.moduleName}}/pkg/domains/weather"
	"{{.moduleName}}/services/factory"
)

func (f *forecast_) Get(ftx factory.Service) (res listable.QueryList[weather.Entity], err error) {
	forcast, err := f.repository.List(ftx)
	if err != nil {
		return
	}

	res.Count = len(forcast)
	res.Data = forcast

	return
}
