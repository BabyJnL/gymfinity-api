package Router

import "github.com/gin-gonic/gin"

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
		
	}

	router.Run("localhost:8080")
}