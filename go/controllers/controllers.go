package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jadson-medeiros/dc-a01/database"
	"github.com/jadson-medeiros/dc-a01/domain"
	"github.com/jadson-medeiros/dc-a01/helpers"
)

func CreateNewItem(w http.ResponseWriter, r *http.Request) {
	var newItem domain.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)

	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var response = helpers.JsonResponse{}
	db := database.ConnectionDB()
	defer db.Close()

	sql := `INSERT INTO tb01 (col_texto, col_dt) VALUES($1,$2) RETURNING id;`
	var lastId int64
	err = db.QueryRow(sql, newItem.Col_Text, time.Now().String()[0:19]).Scan(&lastId)

	if err != nil {
		log.Printf("Error when tried to insert the data: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newItem.Id = lastId

	response = helpers.JsonResponse{
		Type:    "success",
		Message: "item successfully inserted",
		Data:    newItem,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
