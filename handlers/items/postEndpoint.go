package items

import (
	"crud-app/entities"
	"encoding/json"
	"net/http"
)

func (da *ItemDependencyAggregate) Create(w http.ResponseWriter, r *http.Request) {
	var item entities.Item
	json.NewDecoder(r.Body).Decode(&item)
	da.ItemRepository.Create(&item)
	json.NewEncoder(w).Encode(item)
}