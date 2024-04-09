package forecastusecase

import (
	"{{.moduleName}}/apps/weather/internal/repositories"
	"{{.moduleName}}/apps/weather/internal/usecases"
)

type forecast_ struct {
	repository repositories.Forecast
}

func New(
	repository repositories.Forecast,
) usecases.Forecast {
	return &forecast_{repository}
}
