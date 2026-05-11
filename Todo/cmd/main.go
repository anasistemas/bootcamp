package main

import (
	"fmt"
	"os"
	"strings"

	"todo"
)

const archivoTareas = ".todo.json"

func main() {
	lista := &todo.List{}

	lista.Get(archivoTareas)

	if len(os.Args) == 1 {
		for _, t := range *lista {
			fmt.Println(t.Task)
		}
		return
	}

	tarea := strings.Join(os.Args[1:], " ")
	lista.Add(tarea)
	lista.Save(archivoTareas)
}
