package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jadson-medeiros/dc-a01/domain"
	"github.com/jadson-medeiros/dc-a01/helpers"
	"github.com/vingarcia/ksql"
)

type Controller struct {
	db ksql.Provider
}

func NewController(db ksql.Provider) Controller {
	return Controller{db: db}
}

func (c Controller) CreateNewItem(w http.ResponseWriter, r *http.Request) {
	var newItem domain.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)

	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	newItem.Col_dt = time.Now()
	var response = helpers.JsonResponse{}
	err = c.db.Insert(r.Context(), ksql.NewTable("tb01"), &newItem)

	if err != nil {
		log.Printf("Error when tried to insert the data: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response = helpers.JsonResponse{
		Type:    "success",
		Message: "item successfully inserted",
		Data:    newItem,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
