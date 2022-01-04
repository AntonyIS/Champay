package models

import (
	"fmt"
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
