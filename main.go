package main

import (
	"github.com/AntonyIS/Champay/models"
	"github.com/AntonyIS/Champay/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Initilize database features
	models.InitDB()

	// Close DB after a connection
	defer models.DB.Close()
	// Call the route handler to handler request
	routes.RouteHandler()
}
