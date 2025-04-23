package api

import (
	"encoding/json"
	"net/http"
)

func HandleNewGame(w http.ResponseWriter, r *http.Request) {
	game := NewGame()
	storeGame(game)
	json.NewEncoder(w).Encode(game)
}

func HandleMakeMove(w http.ResponseWriter, r *http.Request) {
	var move MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&move); err != nil {
		http.Error(w, "Invalid move", http.StatusBadRequest)
		return
	}

	game := getGame()
	err := game.MakeMove(move.Row, move.Col, move.Direction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(game)
}

func HandleGetState(w http.ResponseWriter, r *http.Request) {
	game := getGame()
	json.NewEncoder(w).Encode(game)
}
