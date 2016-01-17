package main

import (
	"log"
	"net/http"

	"github.com/arnoldcano/muaddib/runner"
)

func main() {
	http.HandleFunc("/", runner.IndexHandler)
	http.HandleFunc("/run", runner.RunHandler)

	log.Println("Listening on 8081")
	http.ListenAndServe(":8081", nil)
}
