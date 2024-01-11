package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("X-API-KEY") // TODO: Change this to the header you need
		if apiKey != "123" {                        // TODO: Change this to the API key you need
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"}) // TODO: Change this to the error you need
			c.Abort()
			return
		}
		c.Next()
	}
}
