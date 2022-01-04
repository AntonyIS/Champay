package routes

import (
	"github.com/gin-gonic/gin"
)

// Initializes route handler, and the server on port 8000
func RouteHandler() {
	router := gin.Default()
	router.POST("/api/v1/users/add", CreateUser)          // Add user : POST
	router.GET("/api/v1/users", GetUsers)                 // Get all users
	router.GET("/api/v1/users/:id", GetUserByID)          // Get user with user id
	router.PUT("/api/v1/users/put/:id", UpdateUser)       // Add user : UPDATE
	router.DELETE("/api/v1/users/delete/:id", DeleteUser) // Add user : DELETE

	// Run the server on port 8000
	router.Run(":8000")

}
