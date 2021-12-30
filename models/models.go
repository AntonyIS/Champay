package models

import (
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
	Phone      uint    `json:"phone"`
	Amount     float64 `json:"amount"`
	NationalId uint    `json:"nationalId"`
	GroupID    uint    `json:"groupId"`
}

type Group struct {
	gorm.Model
	GroupName    string  `json:"groupName"`
	Contribution float64 `json:"contribution"`
	Amount       float64 `json:"amount"`
	Users        *[]User `json:"users"`
}

// Initilize the database and prep database tables
func InitDB() {
	godotenv.Load(".env")
	var (
		dialect  = os.Getenv("DIALECT")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		dbname   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	// Connection link to the database
	db_connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	// Execute database connection
	DB, err := gorm.Open(dialect, db_connection)

	if err != nil {
		log.Fatalf("Error connecting to the database %s", error.Error(err))
	}

	// Migrations
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Group{})

}

// var Groups = []Group{
// 	{Id: 1, GroupName: "TwawezaGroup", Contribution: 0, Amount: 0},
// 	{Id: 2, GroupName: "Chama kuu", Contribution: 0, Amount: 0},
// 	{Id: 3, GroupName: "Tusonge", Contribution: 0, Amount: 0},
// }

// var Users = []User{
// 	{Id: 1, FirstName: "Antony", LastName: "Injila", Email: "antonyshikubu@gmail.com", Phone: 723308900, Amount: 50000, NationalId: 89786745, GroupID: 1},
// 	{Id: 1, FirstName: "Lilian", LastName: "Ts", Email: "lilian@gmail.com", Phone: 723308900, Amount: 50000, NationalId: 89786745, GroupID: 1},
// 	{Id: 1, FirstName: "Brian", LastName: "Aliwa", Email: "al@gmail.com", Phone: 7236709900, Amount: 100000, NationalId: 69786745, GroupID: 2},
// 	{Id: 1, FirstName: "Elon", LastName: "Musk", Email: "elon@gmail.com", Phone: 723308900, Amount: 100000, NationalId: 89346745, GroupID: 2},
// 	{Id: 1, FirstName: "Edith", LastName: "Ay", Email: "ay@gmail.com", Phone: 723308900, Amount: 20000, NationalId: 89786745, GroupID: 3},
// 	{Id: 1, FirstName: "Lavin", LastName: "Wi", Email: "wi@gmail.com", Phone: 723308900, Amount: 20000, NationalId: 89700745, GroupID: 3},
// }
