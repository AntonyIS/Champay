// Author : Antony Injila
// Contributors : Nil :)
// Project name : Champay
// Github : https://github.com/AntonyIS/Champay
// Description : A simple Golang microservice application show casing how different technologies can be used to make a system(CRUD)
// Technologies : Golang,Go Gin (Go REST API framework) Docker, PostgreSQL
package main

import (
	"github.com/AntonyIS/Champay/models"
	"github.com/AntonyIS/Champay/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Initilize database features
	models.Setup()

	// Call the route handler to handler request
	routes.RouteHandler()
}
