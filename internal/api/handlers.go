package api

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Player struct {
	Name            string    `json:"name"`
	Level           int       `json:"level"`
	Exp             int       `json:"exp"`
	NextLevelExp    int       `json:"next_level_exp"`
	Points          int       `json:"points"`
	Attack          int       `json:"attack"`
	Defense         int       `json:"defense"`
	HP              int       `json:"hp"`
	MaxHP           int       `json:"max_hp"`
	Gold            int       `json:"gold"`
	AutoIncome      int       `json:"auto_income"`
	LastCollectTime time.Time `json:"-"`
}

var (
	mu     sync.Mutex
	player = Player{
		Name:            "Hero",
		Level:           1,
		Exp:             0,
		NextLevelExp:    100,
		Points:          0,
		Attack:          5,
		Defense:         3,
		HP:              100,
		MaxHP:           100,
		Gold:            0,
		AutoIncome:      1,
		LastCollectTime: time.Now(),
	}
)

// /click — получить очки вручную
func HandleClick(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	player.Points += player.Attack
	player.Gold += 1
	player.Exp += 1

	// Уровап
	if player.Exp >= player.NextLevelExp {
		player.Level++
		player.Exp = 0
		player.NextLevelExp += 50
		player.MaxHP += 10
		player.HP = player.MaxHP
		player.Attack += 1
		player.Defense += 1
	}

	w.WriteHeader(http.StatusOK)
}

// /status — показать статус игрока с учётом авто-дохода
func HandleStatus(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	elapsed := int(now.Sub(player.LastCollectTime).Seconds())

	player.Points += elapsed * player.AutoIncome
	player.Gold += elapsed * player.AutoIncome
	player.LastCollectTime = now

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}
