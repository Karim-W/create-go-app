package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"{{.moduleName}}/cmd/apps"
	"{{.moduleName}}/cmd/rest"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	flag.Var(&appFlag, "app", "app(s) to run")
	flag.Parse()

	// ========= Setup Env =========
	godotenv.Load()
	if *envFlag != "" {
		godotenv.Load(*envFlag)
	}

	if len(appFlag) == 0 {
		log.Println("No app to run")
		return
	}

	// ========= Setup Apps =========
	apps := make(map[string]apps.Application)

	for _, app := range appFlag {
		if _, ok := apps[app]; ok {
			continue
		}

		// create app(s)
		switch app {
		case "rest":
			apps["rest"] = rest.RestApp()
		default:
			log.Fatalf("Unknown app: %s", app)
		}
	}

	for _, app := range apps {
		go assert(app.Setup())
	}

	// *** Sigterm handler ***
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// ========= Close Apps =========
	for _, app := range apps {
		app.Close()
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err) // just in case, i have trust issues
	}
}
