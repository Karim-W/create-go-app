package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{.moduleName}}/cmd/rest"
	"{{.moduleName}}/internal/config"
	"{{.moduleName}}/internal/constants"
	"{{.moduleName}}/pkg/adapters"
	"{{.moduleName}}/pkg/infra"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// read .env file for local development
	godotenv.Load()
	// read config
	// ========= SetupConfig =========
	config.InitOrDie()
	// set up adapters
	iKey, err := config.GetAppInsightsInstrumentationKey()
	assert(err)
	// ========= SetupAdapters =========
	adpt, err := adapters.SetupAdapters(&adapters.Options{
		InstrumentationKey: iKey,
		ServiceName:        constants.SERVICE_NAME,
	})
	assert(err)
	// ========= SetupInfra =========
	infras, err := infra.SetupInfra(&infra.Options{
		Trx: adpt.Trx,
	})
	assert(err)
	// ========= Setup Repositories =========
	// ========= Setup Services =========
	// ========= Setup Usecases =========
	// ========= Start the app ========
	port, err := config.GetServerPort()
	assert(err)
	go rest.SetupRoutes(&rest.Options{
		Port:   port,
		Logger: zap.L(),
		Trx:    infras.Trx,
	})
	// *** Sigterm handler ***
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutting down")
	adpt.Trx.Close()
	zap.L().Sync()
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err) // just in case, i have trust issues
	}
}
