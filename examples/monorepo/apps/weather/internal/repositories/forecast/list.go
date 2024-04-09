package forecastrepository

import (
	"{{.moduleName}}/pkg/domains/weather"
	"{{.moduleName}}/services/factory"
)

func (f *forecast_) List(ftx factory.Service) ([]weather.Entity, error) {
	return []weather.Entity{
		{
			Date:         "2024-04-10",
			TemperatureC: 25,
			TemperatureF: 76,
			Summary:      "Mild",
		},
		{
			Date:         "2024-04-11",
			TemperatureC: 5,
			TemperatureF: 40,
			Summary:      "Sweltering",
		},
		{
			Date:         "2024-04-12",
			TemperatureC: 48,
			TemperatureF: 118,
			Summary:      "Chilly",
		},
		{
			Date:         "2024-04-13",
			TemperatureC: 5,
			TemperatureF: 40,
			Summary:      "Warm",
		},
		{
			Date:         "2024-04-14",
			TemperatureC: -1,
			TemperatureF: 31,
			Summary:      "Bracing",
		},
	}, nil
}
