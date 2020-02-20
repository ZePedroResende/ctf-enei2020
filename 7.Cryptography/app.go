package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=nessie")
		http.ServeFile(w, r, "./nessie/target/release/nessie")
	})
	err := http.ListenAndServe(":8003", nil)

	if err != nil {
		fmt.Println(err)
	}
}
