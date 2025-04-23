package api

import (
	"encoding/json"
	"net/http"
)

type HealResult struct {
	HPBefore int `json:"hp_before"`
	HPAfter  int `json:"hp_after"`
	GoldLeft int `json:"gold_left"`
}

func HandleHeal(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	const cost = 5
	const healAmount = 10

	if player.Gold < cost {
		http.Error(w, "Not enough gold", http.StatusPaymentRequired)
		return
	}

	if player.HP == player.MaxHP {
		http.Error(w, "HP is already full", http.StatusConflict)
		return
	}

	player.Gold -= cost
	hpBefore := player.HP
	player.HP += healAmount
	if player.HP > player.MaxHP {
		player.HP = player.MaxHP
	}

	result := HealResult{
		HPBefore: hpBefore,
		HPAfter:  player.HP,
		GoldLeft: player.Gold,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
