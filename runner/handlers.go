package runner

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RunHandler(w http.ResponseWriter, r *http.Request) {
	var c http.Client

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	log.Printf("Received request from %s", r.RemoteAddr)
	r2, err := http.NewRequest("POST", "http://usul:8080/run", r.Body)
	r2.Header.Set("Content-Type", "application/json; charset=UTF-8")
	log.Printf("Sent request to %s", r2.Host)
	w2, err := c.Do(r2)
	if err != nil {
		writeError(w, err)
		return
	}
	defer w2.Body.Close()
	log.Printf("Received response from %s", r2.Host)

	_, err = io.Copy(w, w2.Body)
	if err != nil {
		writeError(w, err)
		return
	}
	log.Printf("Sent response to %s", r.RemoteAddr)
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	var c http.Client

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	log.Printf("Received request from %s", r.RemoteAddr)
	r2, err := http.NewRequest("POST", "http://usul:8080/analyze", r.Body)
	r2.Header.Set("Content-Type", "application/json; charset=UTF-8")
	log.Printf("Sent request to %s", r2.Host)
	w2, err := c.Do(r2)
	if err != nil {
		writeError(w, err)
		return
	}
	defer w2.Body.Close()
	log.Printf("Received response from %s", r2.Host)

	_, err = io.Copy(w, w2.Body)
	if err != nil {
		writeError(w, err)
		return
	}
	log.Printf("Sent response to %s", r.RemoteAddr)
}

func writeError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}
