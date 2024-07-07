package items

import (
	"encoding/json"
	"net/http"
)

func (da *ItemDependencyAggregate) Get(w http.ResponseWriter, r *http.Request) {
	items, _ := da.ItemRepository.GetAll()
	json.NewEncoder(w).Encode(items)
}