package main

import (
	"{{.moduleName}}/cmd/rest"
	"{{.moduleName}}/internals/config"
	"{{.moduleName}}/pkg/adapters"
	"{{.moduleName}}/pkg/infra"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	// read .env file for local development
	godotenv.Load()
	// read config
	// ========= SetupConfig =========
	config.InitOrDie()
	// set up adapters
	iKey, err := config.GetAppInsightsInstrumentationKey()
	if err != nil {
		panic(err)
	}
	// ========= SetupAdapters =========
	acore, err := adapters.SetupAdapters(iKey, "my-service")
	if err != nil {
		panic(err)
	}
	// ========= SetupInfra =========
	logger, trx, err := infra.SetupInfra(acore)
	if err != nil {
		panic(err)
	}
	// ========= Setup Repositories =========
	// ========= Setup Services =========
	// ========= Setup Usecases =========
	// ========= Start the app ========
	port, err := config.GetServerPort()
	if err != nil {
		panic(err)
	}
	go rest.SetupRoutes(port, logger, trx)
	// *** Sigterm handler ***
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down")
	acore.Close()
	logger.Sync()
}
