package routes

import (
	"net/http"

	"github.com/AntonyIS/Champay/models"
	"github.com/gin-gonic/gin"
)

// Initializes route handler, and the server on port 8000
func RouteHandler() {
	router := gin.Default()
	router.GET("/api/v1/users", GetUsers)    // Get all users
	router.GET("/api/v1//groups", GetGroups) // Get all groups

	// Run the server on port 8000
	router.Run(":8000")

}

func GetUsers(context *gin.Context) {
	var users = models.Users
	context.IndentedJSON(http.StatusOK, users)
}

func GetGroups(context *gin.Context) {
	var groups = models.Groups
	context.IndentedJSON(http.StatusOK, groups)
}
