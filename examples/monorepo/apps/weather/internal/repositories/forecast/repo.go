package forecastrepository

import (
	"{{.moduleName}}/apps/weather/internal/repositories"
)

type forecast_ struct{}

func New() repositories.Forecast {
	return &forecast_{}
}
