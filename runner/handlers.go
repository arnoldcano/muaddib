package runner

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p, err := filepath.Abs("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	f, err := ioutil.ReadFile(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(f)
}

func RunHandler(w http.ResponseWriter, r *http.Request) {
	var c http.Client

	log.Printf("Received request from %s", r.UserAgent())
	r2, err := http.NewRequest("POST", "http://usul:8080/run", r.Body)
	r2.Header.Set("Content-Type", "application/json")
	log.Printf("Sent request to %s", r2.Host)
	w2, err := c.Do(r2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer w2.Body.Close()
	log.Printf("Received response from %s", r2.Host)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err = io.Copy(w, w2.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Sent response to %s", r.UserAgent())
}
