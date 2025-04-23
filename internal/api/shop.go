package api

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	Name     string `json:"name"`
	Type     string `json:"type"` // "attack", "defense"
	Bonus    int    `json:"bonus"`
	Cost     int    `json:"cost"`
	Equipped bool   `json:"equipped"`
}

var shopItems = []Item{
	{Name: "Bronze Sword", Type: "attack", Bonus: 2, Cost: 20},
	{Name: "Iron Shield", Type: "defense", Bonus: 3, Cost: 25},
}

var inventory []Item

func HandleShop(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shopItems)
}

func HandleInventory(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(inventory)
}

func HandleBuy(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid item", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, shopItem := range shopItems {
		if shopItem.Name == item.Name && player.Gold >= shopItem.Cost {
			player.Gold -= shopItem.Cost
			inventory = append(inventory, shopItem)
			json.NewEncoder(w).Encode(shopItem)
			return
		}
	}

	http.Error(w, "Item not found or not enough gold", http.StatusBadRequest)
}
