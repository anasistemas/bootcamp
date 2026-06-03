package main

import (
	"net/http"

	"github.com/anasistemas/todo"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorReply(w, r, http.StatusNotFound, "page not found")
		return
	}
	textReply(w, r, http.StatusOK, "Hello World")
}

func getAllHandler(datafile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list todo.List
		if err := list.Get(datafile); err != nil {
			errorReply(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		resp := &todoResponse{Results: list}
		jsonReply(w, r, http.StatusOK, resp)
	}
}
