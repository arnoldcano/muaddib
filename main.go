package main

import (
	"log"
	"net/http"

	"github.com/arnoldcano/muaddib/runner"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/run", runner.RunHandler)

	log.Println("Listening on port 8081")
	http.ListenAndServe(":8081", nil)
}
