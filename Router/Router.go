package Router

import (
	"gymfinity-backend-api/Controllers/ClassController"
	"gymfinity-backend-api/Controllers/UserController"

	"github.com/gin-gonic/gin"
)

var trustedProxies = []string{
	"127.0.0.1",
}

func SetupRoutes() {
	router := gin.Default();
	
	// Middlewares
	router.SetTrustedProxies(trustedProxies)

	// API Routes Group
	api := router.Group("/api")
	{
		// Users Routes
		api.GET("/users", UserController.Index)
		api.GET("/users/:id", UserController.Show)
		api.POST("/users", UserController.Create)
		api.PUT("/users/:id", UserController.Update)
		api.DELETE("/users/:id", UserController.Delete)

		// Classes Routes
		api.GET("/classes", ClassController.Index)
		api.GET("/classes/:id", ClassController.Show)
		api.POST("/classes", ClassController.Create)
		api.PUT("/classes/:id", ClassController.Update)
		api.DELETE("/classes/:id", ClassController.Delete)
	}

	router.Run("localhost:8080")
}