package todo

import (
	"os"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	var lista List
	lista.Add("comprar leche")
	lista.Add("hacer ejercicio")

	if len(lista) != 2 {
		t.Errorf("esperaba 2 tareas pero hay %d", len(lista))
	}
	if lista[0].Task != "comprar leche" {
		t.Errorf("la primera tarea debería ser 'comprar leche' pero es '%s'", lista[0].Task)
	}
	if lista[1].Task != "hacer ejercicio" {
		t.Errorf("la segunda tarea debería ser 'hacer ejercicio' pero es '%s'", lista[1].Task)
	}
}

func TestComplete(t *testing.T) {
	var lista List
	lista.Add("comprar leche")
	lista.Add("hacer ejercicio")

	err := lista.Complete(1)
	if err != nil {
		t.Errorf("no esperaba un error aquí: %s", err)
	}
	if lista[0].Task != "comprar leche" {
		t.Errorf("tarea incorrecta: %s", lista[0].Task)
	}
	if !lista[0].Done {
		t.Errorf("'%s' debería estar marcada como completada", lista[0].Task)
	}

	err = lista.Complete(99)
	if err == nil {
		t.Error("con un índice inválido debería retornar error")
	}
}

func TestDelete(t *testing.T) {
	var lista List
	lista.Add("comprar leche")
	lista.Add("hacer ejercicio")
	lista.Add("llamar al médico")

	err := lista.Delete(2)
	if err != nil {
		t.Errorf("no esperaba un error aquí: %s", err)
	}
	if len(lista) != 2 {
		t.Errorf("después de borrar debería haber 2 tareas pero hay %d", len(lista))
	}
	if lista[0].Task != "comprar leche" {
		t.Errorf("la primera tarea debería seguir siendo 'comprar leche' pero es '%s'", lista[0].Task)
	}
	if lista[1].Task != "llamar al médico" {
		t.Errorf("la segunda tarea debería ser 'llamar al médico' pero es '%s'", lista[1].Task)
	}

	err = lista.Delete(99)
	if err == nil {
		t.Error("con un índice inválido debería retornar error")
	}
}

func TestSaveAndGet(t *testing.T) {
	temporal, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatalf("no se pudo crear el archivo temporal: %s", err)
	}
	defer os.Remove(temporal.Name())

	var lista1 List
	lista1.Add("comprar leche")
	lista1.Add("hacer ejercicio")
	lista1.Complete(1)

	err = lista1.Save(temporal.Name())
	if err != nil {
		t.Fatalf("no se pudo guardar la lista: %s", err)
	}

	var lista2 List
	err = lista2.Get(temporal.Name())
	if err != nil {
		t.Fatalf("no se pudo leer la lista: %s", err)
	}

	if len(lista2) != len(lista1) {
		t.Errorf("esperaba %d tareas pero hay %d", len(lista1), len(lista2))
	}
	if lista2[0].Task != lista1[0].Task {
		t.Errorf("la primera tarea debería ser '%s' pero es '%s'", lista1[0].Task, lista2[0].Task)
	}
	if lista2[0].Task == lista1[0].Task && lista2[0].Done != lista1[0].Done {
		t.Errorf("'%s' debería estar completada", lista2[0].Task)
	}
	if lista2[1].Task != lista1[1].Task {
		t.Errorf("la segunda tarea debería ser '%s' pero es '%s'", lista1[1].Task, lista2[1].Task)
	}
}

func TestString(t *testing.T) {
	var lista List
	lista.Add("comprar leche")
	lista.Add("hacer ejercicio")
	lista.Complete(2)

	salida := lista.String()

	esperadaIncompleta := "- [ ] 1: comprar leche\n"
	if !strings.Contains(salida, esperadaIncompleta) {
		t.Errorf("esperaba tarea incompleta %q en la salida, pero se obtuvo:\n%s", esperadaIncompleta, salida)
	}

	esperadaCompleta := "- [X] 2: hacer ejercicio\n"
	if !strings.Contains(salida, esperadaCompleta) {
		t.Errorf("esperaba tarea completa %q en la salida, pero se obtuvo:\n%s", esperadaCompleta, salida)
	}
}
