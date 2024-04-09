package routes

import (
	"playdates/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Public routes
	public := r.Group("/public")
	{
		public.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Public route"})
		})
	}

	// Apply authentication middleware to all subsequent routes.
	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the social media platform!"})
		})
	}

	// Handle undefined routes - ideally, place this after all other route definitions.
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
}
