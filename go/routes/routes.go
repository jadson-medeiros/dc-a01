package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jadson-medeiros/dc-a01/controllers"
	"github.com/jadson-medeiros/dc-a01/middleware"
)

func HandleRequest(controller controllers.Controller) {
	router := mux.NewRouter()
	router.Use(middleware.ContetTypeMiddleware)

	router.HandleFunc("/tb01", controller.CreateNewItem).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
