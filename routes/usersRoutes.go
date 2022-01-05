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
	// Initialize users: gets assigned all users fron the database
	var users []models.User
	// Pull users from the users tables
	models.DB.Find(&users)
	// returns users to the caller
	ctx.IndentedJSON(http.StatusOK, users)
}

// Get a single user with a given user id : READ
func GetUserByID(ctx *gin.Context) {
	// return user with the given id
	// initialize user
	var user models.User
	// Get user id from the url
	userID := ctx.Param("id")
	// Pull user with userID from the database
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
		return
	}

	// Add user into the database
	models.DB.Create(&newUser)
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

	// Check if firstname, lastname, email and phone fields are not empty fields
	if updateUser.FirstName == "" && updateUser.LastName == "" && updateUser.Email == "" && updateUser.Phone == "" {
		log.Fatal("Error ::: Invalid user data")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "firstname, lastname, email and phone fields are empty",
		})
		return
	}

	// Update user
	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName
	user.Email = updateUser.Email
	user.Phone = updateUser.Phone
	user.Amount = updateUser.Amount
	user.NationalId = updateUser.NationalId
	user.GroupID = updateUser.GroupID

	// Save user into the database
	models.DB.Save(&user)

	// Return the updated user
	ctx.IndentedJSON(http.StatusOK, user)

}

// Delete user from the database : DELETE
func DeleteUser(ctx *gin.Context) {
	// Deletes user with id from the database
	// Returns successful message

	// Initialize user
	var deleteUser models.User
	// Get id of user to be deleted from the request
	userID := ctx.Param("id")

	// Delete user from the database
	models.DB.Delete(&deleteUser, userID)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "User delete",
	})
}
