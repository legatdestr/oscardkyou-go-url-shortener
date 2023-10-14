package main

import (
	"net/http"

	"GoUrlShotener/internal/shotener"
)

func main() {
	http.HandleFunc("/url", shotener.SaveURLHandler)
	http.HandleFunc("/", shotener.RedirectHandler)
	http.ListenAndServe(":8080", nil)
}
