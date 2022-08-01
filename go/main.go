package main

import (
	"fmt"

	"github.com/jadson-medeiros/dc-a01/routes"
)

func main() {
	fmt.Println("Starting the server...")
	routes.HandleRequest()
}
