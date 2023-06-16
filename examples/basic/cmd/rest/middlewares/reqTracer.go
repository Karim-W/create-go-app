package middlewares

import (
	"context"
	"{{.moduleName}}/pkg/infra/tracing"
	"{{.moduleName}}/pkg/services/factory"
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
		var ftx factory.Service
		var err error
		if traceparent == "" {
			ftx = factory.NewFactory(context.Background())
			c.Request.Header.Set("traceparent", ftx.GetTraceParent())
		} else {
			ftx, err = factory.NewFactoryFromTraceParent(traceparent)
			if err != nil {
				ftx = factory.NewFactory(context.Background())
			}
		}
		start := time.Now()
		c.Next()
		end := time.Now()
		status := c.Writer.Status()
		bytes := c.Writer.Size()
		trx.TraceRequest(
			ftx.GetContext(),
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
