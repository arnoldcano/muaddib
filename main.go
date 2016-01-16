package main

import (
	"io/ioutil"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(page)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
