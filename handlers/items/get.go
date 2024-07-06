package items

import (
	"crud-app/database"
	"crud-app/entities"
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var items []entities.Item
	database.DB.Find(&items)
	json.NewEncoder(w).Encode(items)
}