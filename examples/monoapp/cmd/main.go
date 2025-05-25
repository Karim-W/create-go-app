package main

import (
	"{{.moduleName}}/apps/{{.serviceName}}/internal/config"
	"{{.moduleName}}/apps/{{.serviceName}}/internal/constants"
	"{{.moduleName}}/pkg/wires"
	"{{.moduleName}}/transports/rest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/karim-w/glose"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load()

	// ========= SetupConfig =========
	config.InitOrDie()

	// ========= Setup Application Wires =========

	modules, err := wires.Run(config.Config, wires.Config{
		ServiceName: constants.SERVICE_NAME,
		Adapters:    wires.AdapterFlags{},
		Infra:       wires.InfraOptions{},
	})
	assert(err)

	// ========= Setup Repositories =========

	// ========= Setup Usecase =========

	// ========= Setup Handlers ========
	handlers := []rest.Controller[gin.IRouter]{
		// Add Handlers here
	}

	// ========= Start the app ========
	port, err := config.GetServerPort()
	assert(err)

	go rest.SetupRoutes(&rest.Options{
		Port:   port,
		Logger: zap.L(),
		Trx:    modules.Infra.Trx,
	}, handlers...)

	// ========= Graceful Shutdown =========
	glose.Watch()
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
