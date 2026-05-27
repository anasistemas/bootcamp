package main

import (
	"log"
	"net/http"

	"mux/handlers"
)

func main() {
	mux := http.NewServeMux()

	home := handlers.HomeHandler{}
	mux.Handle("/", home)

	mux.Handle("/about", http.HandlerFunc(handlers.AboutHandler))

	mux.HandleFunc("/help", handlers.HelpHandler)

	log.Println("Server running on :9090")

	log.Fatal(http.ListenAndServe(":9090", mux))
}
