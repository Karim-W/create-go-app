package rest

import (
	"{{.moduleName}}/pkg/infrastructure/tracing"
	"{{.moduleName}}/transports/rest/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Options struct {
	Port        string
	Logger      *zap.Logger
	Trx         tracing.Tracer
	SwaggerPath string
	Middlewares []gin.HandlerFunc
}

// SetupRoutes initializes the rest package.
// It is called by the main package.
func SetupRoutes(
	opts *Options,
	handlers ...RestHandler[gin.IRouter],
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
		middlewares.TraceRequest(opts.Trx),
		middlewares.ErrorHandlerMiddleware(),
		middlewares.RecoveryMiddleware(),
		middlewares.RootRecoveryMiddleware(opts.Logger),
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
		panic(err)
	}
	return router
}
