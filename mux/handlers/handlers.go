package handlers

import (
	"net/http"
)

type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About"))
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Help"))
}
