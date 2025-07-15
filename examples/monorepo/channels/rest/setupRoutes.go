package rest

import (
	"{{.moduleName}}/channels/rest/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/karim-w/glose"
)

type Options struct {
	Port        string
	SwaggerPath string
	Middlewares []gin.HandlerFunc
}

// SetupRoutes initializes the rest package.
// It is called by the main package.
func SetupRoutes(
	opts *Options,
	handlers ...Controller[gin.IRouter],
) *gin.Engine {
	router := gin.New()

	// ================ Health Check
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "alive",
		})
	})

	// ================ Swagger
	if opts.SwaggerPath != "" {
		router.Static("/swagger", opts.SwaggerPath)
	}

	// ================ Middlewares
	router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/", ""}}),
		middlewares.CorsMiddleware(),
		middlewares.TraceRequest(),
		middlewares.ErrorHandlerMiddleware(),
		middlewares.RecoveryMiddleware(),
		middlewares.RootRecoveryMiddleware(),
	)

	// ================ add custom middlewares
	for _, middleware := range opts.Middlewares {
		router.Use(middleware)
	}

	// ================ Routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "route not found"})
	})

	api := router.Group("")

	for _, handler := range handlers {
		handler.SetupRoutes(api)
	}

	// ================ Run Server
	err := router.Run(opts.Port)
	if err != nil {
		glose.Panik(err)
	}
	return router
}
