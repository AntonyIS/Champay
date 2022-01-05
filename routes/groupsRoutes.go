package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AntonyIS/Champay/models"
	"github.com/gin-gonic/gin"
)

// Returns all groups to the caller : READ
func GetGroups(ctx *gin.Context) {
	// initialize groups
	var groups []models.Group
	// Pull groups from the database
	models.DB.Find(&groups)
	// Returns the results back to the caller
	ctx.IndentedJSON(http.StatusOK, groups)
}

// Return a single group using group id : READ
func GetGroupByID(ctx *gin.Context) {
	// Initialize group
	var group models.Group
	// Get group id
	groupID := ctx.Param("id")
	// Get group from the database using group id
	models.DB.Find(&group, groupID)
	// Return group to the caller
	ctx.IndentedJSON(http.StatusOK, group)

}

// Create new group , returns new group: CREATE/POST
func CreateGroup(ctx *gin.Context) {
	// Initialize new group
	var newGroup models.Group

	// get group data from request payload
	if err := ctx.BindJSON(&newGroup); err != nil {
		log.Fatalf("Error ::: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Request payload has invalid data",
		})
		return
	}
	// Add group into the database
	models.DB.Save(&newGroup)

	ctx.IndentedJSON(http.StatusOK, newGroup)

}

// Update group data , returns updated group: UPDATE/PUT
func UpdateGroup(ctx *gin.Context) {
	// Initialize groups
	var groupFromRequest models.Group
	var groupFromDB models.Group
	groupID := ctx.Param("id")

	// Bind data from request with grogroupFromRequest
	if err := ctx.BindJSON(&groupFromRequest); err != nil {
		fmt.Println("Error bindig group")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Request body is not in the right format",
		})
		return
	}
	// Bind group from the database
	models.DB.Find(&groupFromDB, groupID)

	// Update group with request data
	groupFromDB.GroupName = groupFromRequest.GroupName
	groupFromDB.Contribution = groupFromRequest.Contribution
	groupFromDB.Amount = groupFromRequest.Amount
	groupFromDB.Users = groupFromRequest.Users

	// Save group into the database
	models.DB.Save(groupFromDB)

	// Return updated group
	ctx.IndentedJSON(http.StatusOK, groupFromDB)

}

// Delete group, returns successful delete message
func DeleteGroup(ctx *gin.Context) {
	// Initialize group
	var group models.Group
	// Get group id
	groupID := ctx.Param("id")

	// Delete group
	models.DB.Delete(&group, groupID)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Group delete",
	})
}
