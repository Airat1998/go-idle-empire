package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Airat1998/go-idle-empire/internal/api"
)

func main() {
	http.HandleFunc("/click", api.HandleClick)
	http.HandleFunc("/status", api.HandleStatus)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}