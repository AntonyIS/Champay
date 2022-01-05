package routes

import (
	"github.com/gin-gonic/gin"
)

// Initializes route handler, and the server on port 8000
func RouteHandler() {
	// Initialize router
	router := gin.Default()

	// ##############  Users routes
	router.POST("/api/v1/users/add", CreateUser)          // Add user : POST
	router.GET("/api/v1/users", GetUsers)                 // Get all users
	router.GET("/api/v1/users/:id", GetUserByID)          // Get user with user id
	router.PUT("/api/v1/users/put/:id", UpdateUser)       // Update user : UPDATE
	router.DELETE("/api/v1/users/delete/:id", DeleteUser) // Delete user : DELETE

	// ##############  Groups routes
	router.POST("/api/v1/groups/add", CreateGroup)          // Add grpup : POST
	router.GET("/api/v1/groups", GetGroups)                 // Get all groups
	router.GET("/api/v1/groups/:id", GetGroupByID)          // Get group with group id
	router.PUT("/api/v1/groups/put/:id", UpdateGroup)       // Update group : UPDATE
	router.DELETE("/api/v1/groups/delete/:id", DeleteGroup) // Delete group : DELETE

	// Run the server on port 8000
	router.Run(":8000")

}
