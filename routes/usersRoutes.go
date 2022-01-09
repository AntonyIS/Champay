package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AntonyIS/Champay/models"
	"github.com/gin-gonic/gin"
)

// ################ users routes #################

// Returns all users to the caller : READ
func GetUsers(ctx *gin.Context) {
	// Return all users in the database table
	var userModel models.User
	users := userModel.GetUsers()
	// Pull users from the users tables
	models.DB.Find(&users)
	// returns users to the caller
	ctx.IndentedJSON(http.StatusOK, users)
}

// Get a single user with a given user id : READ
func GetUserByID(ctx *gin.Context) {
	// return user with the given id ,initialize user

	// Get user id from the url
	userID := ctx.Param("id")
	var userModel models.User
	user := userModel.GetUser(userID)
	models.DB.Find(&user, userID)
	ctx.IndentedJSON(http.StatusOK, user)
}

// Add a user into the database table : POST
func CreateUser(ctx *gin.Context) {
	// Receives data from client and stores the data into the database
	// Returns new user to client
	// initialize user
	var newUser models.User

	// Bind request data ti user variable
	if err := ctx.BindJSON(&newUser); err != nil {
		fmt.Printf("ERROR ::: %v\n", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invlaid user data",
		})
		return
	}

	// Add user into the database
	newUser.CreateUser()
	// Return user data to the client: All users can be returned as well
	ctx.IndentedJSON(http.StatusCreated, newUser)
}

// Updte user data : UPDATE
func UpdateUser(ctx *gin.Context) {
	// Receives payload from client, validate and update the data
	// returns updated users back to client

	// Initilize user
	var user models.User
	// Get user ID from the request
	userID := ctx.Param("id")
	// Get user from the database and bind with initialized user
	models.DB.Find(&user, userID)

	// Initilize user to be updated
	var updateUser models.User
	// Bind incoming request data with updateUser
	if err := ctx.BindJSON(&updateUser); err != nil {
		log.Fatalf("Error ::: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Request body is not in the right format",
		})
		return
	}

	user, err := user.UpdateUser(updateUser)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	ctx.IndentedJSON(http.StatusBadRequest, user)

}

// Delete user from the database : DELETE
func DeleteUser(ctx *gin.Context) {
	// Deletes user with id from the database
	// Returns successful message

	// Initialize user
	var deleteUser models.User
	// Get id of user to be deleted from the request
	res, err := deleteUser.DeleteUser()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": res})
}
