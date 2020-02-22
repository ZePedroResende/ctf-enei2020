package main

import (
	"log"
	"net/http"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "template/index.html")
	})
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
