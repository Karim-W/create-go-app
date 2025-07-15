package middlewares

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			if !ctx.Writer.Written() {
				slog.Error("No response was written")
				ctx.AbortWithStatusJSON(500, gin.H{
					"error": "No response was written, this is unexpected",
				})
			}

			// No errors, just return
			return
		}

		var err error

		for _, er := range ctx.Errors {
			err = errors.Join(err, er.Err)
		}

		slog.Error(
			"Error occurred during request handling",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
			slog.String("query", ctx.Request.URL.RawQuery),
			slog.String("traceparent", ctx.Request.Header.Get("traceparent")),
			slog.Any("error", err),
			slog.String("stack", string(debug.Stack())),
		)

		ctx.AbortWithStatusJSON(
			500, gin.H{
				"error":   "An unexpected error occurred",
				"details": err.Error(),
			})
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}

			// Check for a broken connection, as it is not really a
			// condition that warrants a panic stack trace.
			var brokenPipe bool

			if ne, ok := err.(*net.OpError); ok {
				var se *os.SyscallError
				if errors.As(ne, &se) {
					if strings.Contains(
						strings.ToLower(se.Error()), "broken pipe") ||
						strings.Contains(strings.ToLower(se.Error()),
							"connection reset by peer",
						) {
						brokenPipe = true
					}
				}
			}

			httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
			headers := strings.Split(string(httpRequest), "\r\n")
			for idx, header := range headers {
				current := strings.Split(header, ":")
				if current[0] == "Authorization" {
					headers[idx] = current[0] + ": *"
				}
			}

			headersToStr := strings.Join(headers, "\r\n")

			slog.Error(
				"Panic recovered",
				slog.String("headers", headersToStr),
				slog.Any("error", err),
				slog.String("stack", string(debug.Stack())),
				slog.String("method", ctx.Request.Method),
				slog.String("path", ctx.Request.URL.Path),
				slog.String("query", ctx.Request.URL.RawQuery),
				slog.String("traceparent", ctx.Request.Header.Get("traceparent")),
				slog.Bool("is_broken_pipe", brokenPipe),
			)

			response_err := gin.H{
				"error":   "An unexpected error occurred",
				"details": fmt.Sprintf("%v", err),
			}

			if brokenPipe {
				response_err["error"] = "Broken pipe error"
				response_err["details"] = "The connection was closed before the response could be sent."
			}

			ctx.AbortWithStatusJSON(500, response_err)
		}()

		ctx.Next()
	}
}

// Root level panic handler with minimal dependencies
func RootRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error(
					"panic reached root handling (this is bad)",
					slog.Any("error", err),
					slog.String("stack", string(debug.Stack())),
				)
				c.JSON(500, err)
			}
		}()
		c.Next()
	}
}
