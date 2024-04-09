package main

import (
	"log"
	"os"
	"os/signal"
	"{{.moduleName}}/apps/weather/internal/config"
	"{{.moduleName}}/apps/weather/internal/wires"
	"{{.moduleName}}/transports/rest"
	"syscall"

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

	// ========= SetupAdapters =========
	_, err := wires.SetupAdapters(wires.AdapterOptions{})
	assert(err)

	// ========= SetupInfra =========
	infras, err := wires.SetupInfra(wires.InfraOptions{})
	assert(err)

	// ========= Setup Repositories =========
	// ========= Setup Services =========
	err = wires.SetupServices(wires.ServiceOptions{})
	assert(err)
	// ========= Setup Usecase =========
	// ========= Setup Handlers ========
	// ========= Start the app ========
	port, err := config.GetServerPort()
	assert(err)

	go rest.SetupRoutes(&rest.Options{
		Port:        port,
		Logger:      zap.L(),
		Trx:         infras.Trx,
		SwaggerPath: "./swagger",
	})

	// *** Sigterm handler ***
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutting down")
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err) // just in case, i have trust issues
	}
}
