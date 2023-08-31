package rest

import (
	"{{.moduleName}}/cmd/rest/middlewares"
	"{{.moduleName}}/pkg/infra/tracing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Options struct {
	Port   string
	Logger *zap.Logger
	Trx    tracing.Tracer
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
	// ================ Middlewares
	router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/", ""}}),
		middlewares.CorsMiddleware(),
		middlewares.TraceRequest(opts.Trx),
		middlewares.ErrorHandlerMiddleware(),
		middlewares.RecoveryMiddleware(),
		middlewares.RootRecoveryMiddleware(opts.Logger),
	)
	// ================ Routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "route not found"})
	})
	api := router.Group("api")
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
