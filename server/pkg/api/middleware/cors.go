package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (c *middleware) Cors() gin.HandlerFunc {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:5173", "http://127.0.0.1:5173",
		"http://localhost:3000", "http://127.0.0.1:3000",
		"http://localhost:8080", "http://localhost:8080",
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	return cors.New(config)
}
