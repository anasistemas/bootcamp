package main

import "testing"

func check(t *testing.T, operacion string, got, esperado int) {
	t.Helper()
	if got != esperado {
		t.Errorf("%s -> resultado: %d, esperado: %d", operacion, got, esperado)
	}
}

func TestPalabras(t *testing.T) {
	t.Run("una frase", func(t *testing.T) {
		got := procesarEntrada("hola mundo desde go", false, false)
		check(t, "palabras una frase", got, 4)
	})
	t.Run("varias frases", func(t *testing.T) {
		got := procesarEntrada("hola mundo\nesto es una prueba", false, false)
		check(t, "palabras varias frases", got, 6)
	})
	t.Run("una palabra", func(t *testing.T) {
		got := procesarEntrada("hola", false, false)
		check(t, "palabras una palabra", got, 1)
	})
	t.Run("palabra compuesta", func(t *testing.T) {
		got := procesarEntrada("solo lectura", false, false)
		check(t, "palabras compuesta", got, 2)
	})
	t.Run("lineas vacias", func(t *testing.T) {
		got := procesarEntrada("\n\n", false, false)
		check(t, "palabras lineas vacias", got, 0)
	})
	t.Run("exit corta ejecucion", func(t *testing.T) {
		got := procesarEntrada("hola mundo\nEXIT\notra linea", false, false)
		check(t, "palabras con exit", got, 2)
	})
}

func TestLineas(t *testing.T) {
	t.Run("una linea", func(t *testing.T) {
		got := procesarEntrada("hola mundo", true, false)
		check(t, "lineas una linea", got, 1)
	})
	t.Run("sin saltos", func(t *testing.T) {
		got := procesarEntrada("hola mundo esto es una sola linea", true, false)
		check(t, "lineas sin saltos", got, 1)
	})
	t.Run("con saltos", func(t *testing.T) {
		got := procesarEntrada("linea1\n\nlinea2", true, false)
		check(t, "lineas con saltos", got, 3)
	})
	t.Run("exit en medio", func(t *testing.T) {
		got := procesarEntrada("linea1\nexit\nlinea3", true, false)
		check(t, "lineas exit medio", got, 1)
	})
	t.Run("exit al inicio", func(t *testing.T) {
		got := procesarEntrada("exit\nlinea2", true, false)
		check(t, "lineas exit inicio", got, 0)
	})
	t.Run("exit al final", func(t *testing.T) {
		got := procesarEntrada("linea1\nlinea2\nEXIT", true, false)
		check(t, "lineas exit final", got, 2)
	})
}

func TestBytes(t *testing.T) {
	t.Run("una frase", func(t *testing.T) {
		got := procesarEntrada("hola", false, true)
		check(t, "bytes una frase", got, 4)
	})
	t.Run("varias frases", func(t *testing.T) {
		got := procesarEntrada("hola\nmundo", false, true)
		check(t, "bytes varias frases", got, 9)
	})
	t.Run("exit corta ejecucion", func(t *testing.T) {
		got := procesarEntrada("hola\nEXIT\notra", false, true)
		check(t, "bytes con exit", got, 4)
	})
}
