package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// You more than likely want your "Bot User OAuth Access Token" which starts with "xoxb-"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	endpoint := create()
	http.HandleFunc("/events-endpoint", endpoint.eventsEndpoint)
	http.HandleFunc("/ctf", endpoint.ctfEndpoint)

	http.HandleFunc("/674958312", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./flag.png")
	})
	http.HandleFunc("/audio", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./audio.wav")
	})
	http.HandleFunc("/submit", endpoint.submitEndpoint)
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":"+port, nil)
}
