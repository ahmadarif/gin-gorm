package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func FakeAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Login please."})
		} else if token != "mytoken" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid API token."})
		}
	}
}
