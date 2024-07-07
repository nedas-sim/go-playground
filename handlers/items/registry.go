package items

import (
	"crud-app/database"
	"crud-app/entities"

	"github.com/gorilla/mux"
)

type ItemDependencyAggregate struct {
	ItemRepository database.Repository[entities.Item]
}

func RegisterEndpoints(r *mux.Router) {
	da := ItemDependencyAggregate{
		ItemRepository: &database.GenericRepository[entities.Item]{DB: database.DB},
	}

	r.HandleFunc("/api/items", da.Get).Methods("GET")
	r.HandleFunc("/api/items", da.Create).Methods("POST")
}