package main

import (
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to StreamAll!"))
}

func main() {
	addr := ":80"
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	log.Printf("server  is listening at https://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
