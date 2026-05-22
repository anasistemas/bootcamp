package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World"))
}

func main() {
	port := flag.Int("p", 8080, "puerto del servidor")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Servidor corriendo en puerto %d", *port)
	log.Fatal(http.ListenAndServe(addr, mux))
}
