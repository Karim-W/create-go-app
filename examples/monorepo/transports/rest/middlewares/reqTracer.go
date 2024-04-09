package middlewares

import (
	"{{.moduleName}}/pkg/infrastructure/tracing"
	"{{.moduleName}}/services/factory"
	"time"

	"github.com/gin-gonic/gin"
)

func TraceRequest(
	trx tracing.Tracer,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		agent := c.Request.UserAgent()
		ip := c.ClientIP()
		traceparent := c.Request.Header.Get("traceparent")
		ftx, _ := factory.NewFactoryFromTraceParent(traceparent)
		c.Set("ftx", ftx)
		c.Request.Header.Set("traceparent", ftx.TraceParent())
		c.Header("traceparent", ftx.TraceParent())
		start := time.Now()
		c.Next()
		end := time.Now()
		status := c.Writer.Status()
		bytes := c.Writer.Size()
		trx.TraceRequest(
			ftx.Context(),
			method,
			path,
			query,
			status,
			bytes,
			ip,
			agent,
			start,
			end,
			map[string]string{},
		)
	}
}
