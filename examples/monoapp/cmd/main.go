package main

import (
	"{{.moduleName}}/apps/{{.serviceName}}/internal/config"
	"{{.moduleName}}/apps/{{.serviceName}}/internal/constants"
	"{{.moduleName}}/channels/rest"
	"{{.moduleName}}/pkg/utils/assert"
	"{{.moduleName}}/pkg/wires"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/karim-w/glose"
)

func main() {
	godotenv.Load()

	// ========= SetupConfig =========
	config.InitOrDie()

	// ========= Setup Application Wires =========
	err := wires.Run(config.Config, wires.Config{
		ServiceName: constants.SERVICE_NAME,
		Adapters:    wires.AdapterFlags{},
		Infra: wires.InfraOptions{
			PSQL:  false,
			Redis: false,
		},
	})

	assert.NoErr(err)

	// ========= Setup Repositories =========

	// ========= Setup Usecase =========

	// ========= Setup Handlers ========
	handlers := []rest.Controller[gin.IRouter]{
		// Add Handlers here
	}

	// ========= Start the app ========
	port, err := config.GetServerPort()
	assert.NoErr(err)

	go rest.SetupRoutes(&rest.Options{
		Port: port,
	}, handlers...)

	// ========= Graceful Shutdown =========
	glose.Watch()
}
