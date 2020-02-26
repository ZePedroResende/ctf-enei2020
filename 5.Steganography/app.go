package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=cat.gif")
		http.ServeFile(w, r, "./assets/cat.gif")
	})
	err := http.ListenAndServe(":1235", nil)

	if err != nil {
		fmt.Println(err)
	}
}
