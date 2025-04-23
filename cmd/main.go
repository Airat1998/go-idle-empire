package main

import (
	"log"
	"net/http"

	"github.com/Airat1998/go-idle-empire/internal/api"
)

func main() {
	http.HandleFunc("/status", api.HandleStatus)
	http.HandleFunc("/click", api.HandleClick)
	http.HandleFunc("/upgrade", api.HandleUpgrade)
	http.HandleFunc("/battle", api.HandleBattle)
	http.HandleFunc("/heal", api.HandleHeal)
	http.HandleFunc("/collect", api.HandleCollect)
	http.HandleFunc("/shop", api.HandleShop)
	http.HandleFunc("/inventory", api.HandleInventory)
	http.HandleFunc("/buy", api.HandleBuy)
	http.HandleFunc("/equip", api.HandleEquip)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
