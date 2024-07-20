package main

import (
	"log"

	"{{.moduleName}}/apps/{{.serviceName}}/internal/config"
	"{{.moduleName}}/apps/{{.serviceName}}/internal/wires"
	"{{.moduleName}}/transports/rest"

	"github.com/joho/godotenv"
	"github.com/karim-w/glose"
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
	adpt, err := wires.SetupAdapters(wires.AdapterOptions{})
	assert(err)

	// ========= SetupInfra =========
	infras, err := wires.SetupInfra(wires.InfraOptions{})
	assert(err)

	// ========= Setup Repositories =========
	// ========= Setup Services =========
	svcs, err := wires.SetupServices(wires.ServiceOptions{})
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

	// ========= Graceful Shutdown =========
	glose.Watch(
		&svcs,
		&infras,
		&adpt,
	)
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err) // just in case, i have trust issues
	}
}
