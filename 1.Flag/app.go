package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.Handle("/Zmlub3M=", http.HandlerFunc(downloadHandler))

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./flag/flag.png")

}
