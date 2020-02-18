package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./bin/flag")
	})
	err := http.ListenAndServe(":8001", nil)

	if err != nil {
		fmt.Println(err)
	}
}
