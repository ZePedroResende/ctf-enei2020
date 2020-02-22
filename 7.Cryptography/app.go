package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=nessie")
		http.ServeFile(w, r, "./nessie/target/release/nessie")
	})

	s := "It ferries people across the sea on its back. It may sing an enchanting cry if it is in a good mood."
	h := sha512.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))
	fmt.Println(hash)
	url := fmt.Sprintf("/%s", hash)

	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=rei_espadas.png")
		http.ServeFile(w, r, "./rei_espadas.png")
	})
	err := http.ListenAndServe(":8003", nil)

	if err != nil {
		fmt.Println(err)
	}
}
