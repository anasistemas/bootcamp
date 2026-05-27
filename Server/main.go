package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("p", 8080, "puerto del servidor")
	host := flag.String("h", "localhost", "host del servidor")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	server := &http.Server{
		Addr:    addr,
		Handler: newMux(),
	}

	log.Printf("Servidor corriendo en %s", addr)
	log.Fatal(server.ListenAndServe())
}
