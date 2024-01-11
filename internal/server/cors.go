package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                         // TODO: Change this to the domain of your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},                     // TODO: Change this to the methods you need
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // TODO: Change this to the headers you need
		AllowCredentials: true,                                                                  // TODO: Change this if you need cookies
		ExposeHeaders:    []string{"Content-Length"},                                            // TODO: Change this if you need cookies
		MaxAge:           12 * 60 * 60,                                                          // TODO: Change this to the timeout you need
	})
}
