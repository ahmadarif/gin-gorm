package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// on request handler
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
