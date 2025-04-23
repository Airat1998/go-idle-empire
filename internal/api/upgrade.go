package api

import (
	"encoding/json"
	"net/http"
)

type UpgradeRequest struct {
	Type string `json:"type"`
}

func HandleUpgrade(w http.ResponseWriter, r *http.Request) {
	var req UpgradeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid upgrade request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	const upgradeCost = 10

	if player.Gold < upgradeCost {
		http.Error(w, "Not enough gold", http.StatusPaymentRequired)
		return
	}

	switch req.Type {
	case "attack":
		player.Attack++
	case "defense":
		player.Defense++
	case "income":
		player.AutoIncome++
	case "hp":
		player.MaxHP += 10
		player.HP = player.MaxHP
	default:
		http.Error(w, "Invalid upgrade type", http.StatusBadRequest)
		return
	}

	player.Gold -= upgradeCost
	w.WriteHeader(http.StatusOK)
}
