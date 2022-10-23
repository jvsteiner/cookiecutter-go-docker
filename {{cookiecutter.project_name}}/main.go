package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	log.Fatal(server.ListenAndServe())
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{Greeting: "Hello " + request.Name}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)

	return &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
}
