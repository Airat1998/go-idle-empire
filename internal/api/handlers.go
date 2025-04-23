package api

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	mu     sync.Mutex
	points int
)

func HandleClick(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	points++
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
