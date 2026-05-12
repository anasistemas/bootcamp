package main

import (
	"flag"
	"fmt"
	"os"

	"todo"
)

const archivoTareas = ".todo.json"

func salir(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	list := flag.Bool("list", false, "Ver tareas incompletas")
	task := flag.String("task", "", "Agregar una nueva tarea")
	complete := flag.Int("complete", 0, "Completar una tarea por número")
	delete := flag.Int("delete", 0, "Eliminar una tarea por número")

	flag.Parse()

	if !*list && *task == "" && *complete == 0 && *delete == 0 {
		fmt.Fprintln(os.Stderr, "Indica un comando: -list, -task, -complete o -delete")
		os.Exit(1)
	}

	l := &todo.List{}
	if err := l.Get(archivoTareas); err != nil {
		salir(err)
	}

	switch {
	case *list:
		for _, t := range *l {
			if !t.Done {
				fmt.Printf("Title: %s, Done: %t, CreatedAt: %s, CompletedAt: %s\n",
					t.Task, t.Done, t.CreatedAt, t.CompletedAt)
			}
		}

	case *complete != 0:
		if err := l.Complete(*complete); err != nil {
			salir(err)
		}
		if err := l.Save(archivoTareas); err != nil {
			salir(err)
		}

	case *delete != 0:
		if err := l.Delete(*delete); err != nil {
			salir(err)
		}
		if err := l.Save(archivoTareas); err != nil {
			salir(err)
		}

	case *task != "":
		l.Add(*task)
		if err := l.Save(archivoTareas); err != nil {
			salir(err)
		}
	}
}
