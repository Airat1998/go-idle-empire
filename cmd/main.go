package main

import (
	"log"
	"net/http"

	"github.com/Airat1998/go-idle-empire/internal/api"
)

func main() {
	http.HandleFunc("/api/new", api.HandleNewGame)
	http.HandleFunc("/api/move", api.HandleMakeMove)
	http.HandleFunc("/api/state", api.HandleGetState)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
