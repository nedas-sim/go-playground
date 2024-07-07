package items

import (
	"crud-app/entities"
	"encoding/json"
	"net/http"
)

type CreateItemResponse struct {
	ID uint `json:"id"`
}

func (da *ItemDependencyAggregate) Create(w http.ResponseWriter, r *http.Request) {
	var item entities.Item
	json.NewDecoder(r.Body).Decode(&item)
	
	da.ItemRepository.Create(&item)

	response := CreateItemResponse{
		ID: item.ID,
	}
	json.NewEncoder(w).Encode(response)
}