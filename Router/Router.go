package Router

import (
	"gymfinity-backend-api/Controllers/ClassController"
	"gymfinity-backend-api/Controllers/ClassScheduleController"
	"gymfinity-backend-api/Controllers/FacilityController"
	"gymfinity-backend-api/Controllers/FacilityStatusController"
	"gymfinity-backend-api/Controllers/PaymentController"
	"gymfinity-backend-api/Controllers/ReservationController"
	"gymfinity-backend-api/Controllers/UserController"
	"gymfinity-backend-api/Controllers/UserRoleController"
	"gymfinity-backend-api/Controllers/UserStatusController"

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
		// User Routes
		api.GET("/users", UserController.Index)
		api.GET("/users/:id", UserController.Show)
		api.POST("/users", UserController.Create)
		api.PUT("/users/:id", UserController.Update)
		api.DELETE("/users/:id", UserController.Delete)

		// User Status Routes
		api.GET("/user-statuses", UserStatusController.Index)

		// User Role Routes
		api.GET("/user-roles", UserRoleController.Index)

		// Class Routes
		api.GET("/classes", ClassController.Index)
		api.GET("/classes/:id", ClassController.Show)
		api.POST("/classes", ClassController.Create)
		api.PUT("/classes/:id", ClassController.Update)
		api.DELETE("/classes/:id", ClassController.Delete)

		// Class Schedule Routes
		api.GET("/class-schedules", ClassScheduleController.Index)
		api.GET("/class-schedules/:id", ClassScheduleController.Show)
		api.POST("/class-schedules", ClassScheduleController.Create)
		api.PUT("/class-schedules/:id", ClassScheduleController.Update)
		api.DELETE("/class-schedules/:id", ClassScheduleController.Delete)

		// Facility Routes
		api.GET("/facilities", FacilityController.Index)
		api.GET("/facilities/:id", FacilityController.Show)
		api.POST("/facilities", FacilityController.Create)
		api.PUT("/facilities/:id", FacilityController.Update)
		api.DELETE("/facilities/:id", FacilityController.Delete)

		// Facility Status Routes
		api.GET("/facility-statuses", FacilityStatusController.Index)
		
		// Reservation Routes
		api.GET("/reservations", ReservationController.Index)
		api.GET("/reservations/:id", ReservationController.Show)
		api.POST("/reservations", ReservationController.Create)
		api.PUT("/reservations/:id", ReservationController.Update)

		// Payment Routes
		api.GET("/payments", PaymentController.Index)
		api.GET("/payments/:id", PaymentController.Show)
		api.POST("/payments", PaymentController.Create)
	}

	router.Run("localhost:8080")
}