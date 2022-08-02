package main

import (
	"fmt"

	"github.com/jadson-medeiros/dc-a01/controllers"
	"github.com/jadson-medeiros/dc-a01/database"
	"github.com/jadson-medeiros/dc-a01/routes"
)

func main() {
	fmt.Println("Starting the server...")
	db := database.ConnectionDB()

	defer db.Close()

	controller := controllers.NewController(db)
	routes.HandleRequest(controller)
}
