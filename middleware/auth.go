package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
