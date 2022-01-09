package models

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

// Define structure for users of th application
type User struct {
	gorm.Model
	FirstName  string  `json:"firstname"`
	LastName   string  `json:"lastname"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Amount     float64 `json:"amount"`
	NationalId string  `json:"nationalId"`
	GroupID    string  `json:"groupId"`
}

type Group struct {
	gorm.Model
	GroupName    string  `json:"groupName"`
	Contribution float64 `json:"contribution"`
	Amount       float64 `json:"amount"`
	Users        *[]User `json:"users"`
}

func (*User) GetUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}

func (*User) GetUser(userID string) User {
	var user User
	DB.Find(&user, userID)
	return user
}

func (usr *User) UpdateUser(userData User) (User, error) {
	// Receives payload from client, validate and update the data

	// Initilize user
	var user User
	// Get user ID from the request
	userID := usr.ID
	// Get user from the database and bind with initialized user
	DB.Find(&user, userID)

	// Check if firstname, lastname, email and phone fields are not empty fields
	if userData.FirstName == "" && userData.LastName == "" && userData.Email == "" && userData.Phone == "" {
		log.Fatal("Error ::: Invalid user data")
		return User{}, errors.New("Invalid data")
	}

	// Update user
	user.FirstName = userData.FirstName
	user.LastName = userData.LastName
	user.Email = userData.Email
	user.Phone = userData.Phone
	user.Amount = userData.Amount
	user.NationalId = userData.NationalId
	user.GroupID = userData.GroupID

	// Save user into the database
	DB.Save(&user)

	// Return the updated user
	return user, nil
}

func (usr *User) CreateUser() (User, error) {
	// Receives data from client and stores the data into the database
	// Returns new user to client
	// Add user into the database
	DB.Create(&usr)
	// Return user data to the client: All users can be returned as well
	return *usr, nil
}

func (usr *User) DeleteUser() (string, error) {
	// Deletes user with id from the database
	// Returns successful message

	// Initialize user
	var deleteUser User
	// Get id of user to be deleted from the request
	// Delete user from the database
	DB.Delete(&deleteUser, usr.ID)

	return "User deleted successfuly", nil
}

// Initilize the database and prep database tables
func Setup() {
	godotenv.Load(".env")
	var (
		dialect  = os.Getenv("DIALECT")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		dbname   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)

	// Connection string to the data
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	// // Execute database connection
	DB, err = gorm.Open(dialect, conn)

	// // Check if connection to database has error or issues
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database")
	}

	// Create a Items table in the database if it does exists
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Group{})
	// // Setup database connection
}
