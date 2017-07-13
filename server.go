package main

import (
	"net/http"

	"github.com/l0n3r4n83r/goshort/handler"
)

func main() {
	http.HandleFunc("/encode", handler.URLHandler)
	http.ListenAndServe(":8080", nil)
}
