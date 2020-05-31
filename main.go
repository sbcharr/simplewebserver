package main

import (
	"log"
	"net/http"
)

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	http.HandleFunc("/", ServeHttp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
