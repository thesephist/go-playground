package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Route: " + r.RequestURI))
	})
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Could not start server", err)
	}
	log.Println("Started server at :8080")
}
