package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/anasistemas/todo"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorReply(w, r, http.StatusNotFound, "page not found")
		return
	}
	textReply(w, r, http.StatusOK, "Hello World")
}

func router(datafile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list todo.List
		if err := list.Get(datafile); err != nil {
			errorReply(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		parts := strings.Split(r.URL.Path, "/")

		if len(parts) == 3 && parts[2] != "" {
			id, err := validateID(parts[2], &list)
			if err != nil {
				errorReply(w, r, http.StatusBadRequest, err.Error())
				return
			}

			switch r.Method {
			case http.MethodGet:
				getOneHandler(w, r, &list, id)
			default:
				errorReply(w, r, http.StatusMethodNotAllowed, "method not allowed")
			}
			return
		}

		switch r.Method {
		case http.MethodGet:
			getAllHandler(w, r, &list)
		case http.MethodPost:
			addHandler(w, r, &list, datafile)
		default:
			errorReply(w, r, http.StatusMethodNotAllowed, "method not allowed")
		}
	}
}

func validateID(idStr string, list *todo.List) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("id inválido: %q no es un número", idStr)
	}
	if id < 1 {
		return 0, errors.New("id debe ser mayor o igual a 1")
	}
	if id > len(*list) {
		return 0, fmt.Errorf("id %d no existe en la lista", id)
	}
	return id, nil
}

func getAllHandler(w http.ResponseWriter, r *http.Request, list *todo.List) {
	resp := &todoResponse{Results: *list}
	jsonReply(w, r, http.StatusOK, resp)
}

func getOneHandler(w http.ResponseWriter, r *http.Request, list *todo.List, id int) {
	result := todo.List{(*list)[id-1]}
	resp := &todoResponse{Results: result}
	jsonReply(w, r, http.StatusOK, resp)
}

func addHandler(w http.ResponseWriter, r *http.Request, list *todo.List, datafile string) {
	type NewTask struct {
		Task string `json:"task"`
	}

	var item NewTask

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		errorReply(w, r, http.StatusBadRequest, err.Error())
		return
	}

	list.Add(item.Task)

	if err := list.Save(datafile); err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	textReply(w, r, http.StatusCreated, "todo created successfully")
}
