package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Enemy struct {
	Name    string `json:"name"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

type BattleResult struct {
	EnemyName string `json:"enemy"`
	PlayerHP  int    `json:"player_hp"`
	Victory   bool   `json:"victory"`
	ExpGained int    `json:"exp"`
	GoldGained int   `json:"gold"`
}

func HandleBattle(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	enemy := Enemy{
		Name:    "Slime",
		HP:      rand.Intn(30) + 20,
		Attack:  rand.Intn(5) + 2,
		Defense: rand.Intn(3),
	}

	playerDamage := max(1, player.Attack - enemy.Defense)
	enemyDamage := max(1, enemy.Attack - player.Defense)

	rounds := 0
	for enemy.HP > 0 && player.HP > 0 && rounds < 10 {
		enemy.HP -= playerDamage
		if enemy.HP <= 0 {
			break
		}
		player.HP -= enemyDamage
		rounds++
	}

	result := BattleResult{
		EnemyName: enemy.Name,
		PlayerHP:  player.HP,
		Victory:   enemy.HP <= 0,
	}

	if result.Victory {
		result.ExpGained = 10
		result.GoldGained = 5
		player.Exp += result.ExpGained
		player.Gold += result.GoldGained
	}

	if player.Exp >= player.NextLevelExp {
		player.Level++
		player.Exp = 0
		player.NextLevelExp += 50
		player.MaxHP += 10
		player.HP = player.MaxHP
		player.Attack += 1
		player.Defense += 1
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
