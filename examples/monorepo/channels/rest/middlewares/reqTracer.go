package middlewares

import (
	"context"

	"{{.moduleName}}/pkg/utils/w3c"

	"github.com/gin-gonic/gin"
	"github.com/karim-w/coreruntime"
	"github.com/karim-w/gopts"
)

func TraceRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceparent := c.Request.Header.Get("traceparent")

		var (
			trace_id_opt gopts.Option[string]
			span_id_opt  gopts.Option[string]
		)

		is_valid, trace_id, parent_id, _ := w3c.ParseTraceHeaders(traceparent)
		if is_valid {
			trace_id_opt = gopts.Some(trace_id)
			span_id_opt = gopts.Some(parent_id)
		}

		engine := coreruntime.NewWithOptions(coreruntime.EngineOptions{
			Context:  gopts.Some[context.Context](c),
			TraceID:  trace_id_opt,
			ParentID: span_id_opt,
		})

		defer engine.Cleanup()

		// TODO:
		// method := c.Request.Method
		// path := c.Request.URL.Path
		// query := c.Request.URL.RawQuery
		// agent := c.Request.UserAgent()
		// ip := c.ClientIP()
		// start := time.Now()

		// TODO:
		c.Set("engine", engine)

		c.Next()

		// TODO:
		// end := time.Now()
		// status := c.Writer.Status()
		// bytes := c.Writer.Size()
	}
}
