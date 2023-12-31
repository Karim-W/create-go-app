package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Cors is a disgusting thing that exists. I always allow all origins, methods, and headers.
// Feel free to change this to your liking.
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().
			Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, traceparent, Cache-Control, X-Requested-With")
		c.Writer.Header().
			Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT , PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
