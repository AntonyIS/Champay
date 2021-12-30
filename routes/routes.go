package routes

import (
	"net/http"

	"github.com/AntonyIS/Champay/models"
	"github.com/gin-gonic/gin"
)

// Initializes route handler, and the server on port 8000
func RouteHandler() {
	router := gin.Default()
	router.GET("/api/v1/users", GetUsers)        // Get all users
	router.GET("/api/v1/users/:id", GetUserByID) // Get user with user id
	router.GET("/api/v1/groups", GetGroups)      // Get all groups

	// Run the server on port 8000
	router.Run(":8000")

}

// ################ users routes #################
// Returns all users to the caller
func GetUsers(context *gin.Context) {
	// Initialize users
	var users []models.User
	// Pull users from the users tables
	models.DB.Find(&users)
	// returns users to the caller
	context.IndentedJSON(http.StatusOK, users)
}

// Get a single user with a given user id
func GetUserByID(context *gin.Context) {
	// initialize user
	var user models.User
	// Get user id from the url
	userID := context.Param("id")
	// Pull user with userID from the database
	models.DB.Find(&user, userID)
	context.IndentedJSON(http.StatusOK, user)
}

// Returns all users to the caller
func GetGroups(context *gin.Context) {
	// initialize groups
	var groups []models.Group
	// Pull groups from the database
	models.DB.Find(&groups)
	// Returns the results back to the caller
	context.IndentedJSON(http.StatusOK, groups)
}
