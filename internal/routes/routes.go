package routes

import (
	"playdates/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Apply authentication middleware to all routes by default
	r.Use(middleware.AuthMiddleware())

	// Define protected routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the social media platform!"})
	})

	// Public routes
	public := r.Group("/public")
	{
		public.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Public route"})
		})
	}
}
