package api

import (
	"encoding/json"
	"net/http"
)

type EquipRequest struct {
	Name string `json:"name"`
}

func HandleEquip(w http.ResponseWriter, r *http.Request) {
	var req EquipRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid item", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, item := range inventory {
		if item.Name == req.Name && !item.Equipped {
			switch item.Type {
			case "attack":
				player.Attack += item.Bonus
			case "defense":
				player.Defense += item.Bonus
			}
			inventory[i].Equipped = true
			json.NewEncoder(w).Encode(inventory[i])
			return
		}
	}

	http.Error(w, "Item not found or already equipped", http.StatusNotFound)
}
