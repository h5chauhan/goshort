package handler

import "net/http"

// URLHandler handles requets to shorten urls
func URLHandler(w http.ResponseWriter, req *http.Request) {
	url := req.FormValue("url")
}

func encode(in string) string {

}
