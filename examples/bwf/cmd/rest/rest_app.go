package rest

import (
	"{{.moduleName}}/cmd/apps"
	"{{.moduleName}}/internal/config"
	"{{.moduleName}}/pkg/adapters"
	"{{.moduleName}}/pkg/infra"

	"go.uber.org/zap"
)

type restApp struct {
	// TODO: add fields to close resources here
}

func RestApp() apps.Application {
	return &restApp{}
}

func (a *restApp) Setup() error {
	config.InitRestConfigOrDie()

	// set up adapters

	// ========= SetupAdapters =========
	_, err := adapters.SetupAdapters(&adapters.Options{})
	if err != nil {
		return err
	}

	// ========= SetupInfra =========
	infras, err := infra.SetupInfra(&infra.Options{})
	if err != nil {
		return err
	}

	// ========= Setup Repositories =========
	// ========= Setup Services =========
	// ========= Setup Usecases =========
	// ========= Start the app ========

	port, err := config.GetServerPort()
	if err != nil {
		return err
	}

	SetupRoutes(&Options{
		Port:   port,
		Logger: zap.L(),
		Trx:    infras.Trx,
	})

	return nil
}

func (a *restApp) Close() {
	zap.L().Info("Closing rest app")
	zap.L().Sync()
	// TODO: close resources here
}
