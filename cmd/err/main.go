package main

import (
	"SimpleExchangeRates/storage"
	"net/http"
)

func main() {
	http.HandleFunc("/", storage.FuncProvider)
	if err := http.ListenAndServe(":8087", nil); err != nil {
		panic(err)
	}
}
