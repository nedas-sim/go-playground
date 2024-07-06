package items

import (
	"crud-app/database"
	"crud-app/entities"
	"encoding/json"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var item entities.Item
	json.NewDecoder(r.Body).Decode(&item)
	database.DB.Create(&item)
	json.NewEncoder(w).Encode(item)
}