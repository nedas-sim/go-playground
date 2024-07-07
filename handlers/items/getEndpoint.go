package items

import (
	"crud-app/entities"
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
)

type ItemDto struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Price    decimal.Decimal `json:"price"`
	Quantity int             `json:"quantity"`
}

type GetItemsResponse struct {
	Items []ItemDto `json:"items"`
}

func mapToDto(item entities.Item) ItemDto {
	return ItemDto{
		ID:       item.ID,
		Name:     item.Name,
		Price:    item.Price,
		Quantity: item.Quantity,
	}
}

func mapResponse(items []entities.Item) GetItemsResponse {
	itemsDto := make([]ItemDto, len(items))
	for i, item := range items {
		itemsDto[i] = mapToDto(item)
	}

	return GetItemsResponse{Items: itemsDto}
}

func (da *ItemDependencyAggregate) Get(w http.ResponseWriter, r *http.Request) {
	items, err := da.ItemRepository.GetAll()
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	response := mapResponse(items)
	json.NewEncoder(w).Encode(response)
}