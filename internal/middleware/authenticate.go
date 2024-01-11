package middleware

import (
	"server-go/internal/configs"
	"server-go/internal/services"

	"github.com/gin-gonic/gin"
)

func Authentication(cfg *configs.Config) gin.HandlerFunc {
	// create a new token service
	var tokenServices = services.NewTokenService(cfg)

	// return a middleware function that will be executed for every incoming request
	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{} // map to store the claims
		auth := c.GetHeader("Authorization") // get the token from the header with key Authorization

		if len(auth) < 7 {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		} else {
			claimMap, err = tokenServices.GetClaims(auth[7:]) // get the claims from the token and store it in claimMap
			if err != nil {
				c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
				return
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("UserID", claimMap["user_id"]) // set the claims in the context
		c.Set("Username", claimMap["username"])
		c.Set("Phone", claimMap["phone"])
		c.Set("Email", claimMap["email"])
		c.Set("Role", claimMap["role"])
		c.Next()
	}
}

func Authorization(r []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
	}
}
