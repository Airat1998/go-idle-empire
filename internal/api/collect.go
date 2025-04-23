package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type CollectResult struct {
	GoldCollected   int `json:"gold_collected"`
	PointsCollected int `json:"points_collected"`
}

func HandleCollect(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	elapsed := int(now.Sub(player.LastCollectTime).Seconds())
	player.LastCollectTime = now

	gold := elapsed * player.AutoIncome
	points := elapsed * player.AutoIncome

	player.Gold += gold
	player.Points += points

	json.NewEncoder(w).Encode(CollectResult{
		GoldCollected:   gold,
		PointsCollected: points,
	})
}
