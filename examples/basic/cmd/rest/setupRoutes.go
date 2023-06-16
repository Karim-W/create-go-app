package rest

import (
	"{{.moduleName}}/cmd/rest/middlewares"
	"{{.moduleName}}/pkg/infra/tracing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SetupRoutes initializes the rest package.
// It is called by the main package.
func SetupRoutes(
	port string,
	logger *zap.Logger,
	trx tracing.Tracer,
	handlers ...RestHandler[gin.RouterGroup],
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
		middlewares.TraceRequest(trx),
		middlewares.ErrorHandlerMiddleware(),
		middlewares.RecoveryMiddleware(),
		middlewares.RootRecoveryMiddleware(logger),
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
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
	return router
}
