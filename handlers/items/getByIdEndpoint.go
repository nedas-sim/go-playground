package items

import (
	"crud-app/entities"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (da *ItemDependencyAggregate) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	itemId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var item entities.Item
	item, err = da.ItemRepository.GetById(uint(itemId))
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}