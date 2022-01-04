package routes

import (
	"net/http"

	"github.com/AntonyIS/Champay/models"
	"github.com/gin-gonic/gin"
)

// Returns all users to the caller
func GetGroups(context *gin.Context) {
	// initialize groups
	var groups []models.Group
	// Pull groups from the database
	models.DB.Find(&groups)
	// Returns the results back to the caller
	context.IndentedJSON(http.StatusOK, groups)
}
