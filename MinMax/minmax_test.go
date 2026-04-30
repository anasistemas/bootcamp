package main

import (
	"reflect"
	"testing"
)

func checkResult(t *testing.T, operacion string, got, expected []float64) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("operación: %s\n  resultado: %v\n  esperado:  %v", operacion, got, expected)
	}
}

func TestMinmax(t *testing.T) {
	t.Run("varios valores mixtos", func(t *testing.T) {
		got := minmax(2, 8, 1, 3, 5, 8, 10)
		checkResult(t, "minmax(2, 8, 1, 3, 5, 8, 10)", got, []float64{3, 5, 8})
	})

	t.Run("un solo valor en rango", func(t *testing.T) {
		got := minmax(1, 10, 5)
		checkResult(t, "minmax(1, 10, 5)", got, []float64{5})
	})

	t.Run("un solo valor fuera del rango", func(t *testing.T) {
		got := minmax(1, 10, 15)
		checkResult(t, "minmax(1, 10, 15)", got, nil)
	})

	t.Run("min mayor que max", func(t *testing.T) {
		got := minmax(10, 2, 5, 7, 9)
		checkResult(t, "minmax(10, 2, 5, 7, 9)", got, nil)
	})

	t.Run("ningún valor en rango", func(t *testing.T) {
		got := minmax(5, 10, 1, 2, 3, 4, 11)
		checkResult(t, "minmax(5, 10, 1, 2, 3, 4, 11)", got, nil)
	})

	t.Run("min y max negativos", func(t *testing.T) {
		got := minmax(-10, -2, -15, -8, -5, -1, 0)
		checkResult(t, "minmax(-10, -2, -15, -8, -5, -1, 0)", got, []float64{-8, -5})
	})

	t.Run("min igual a max con valor exacto", func(t *testing.T) {
		got := minmax(5, 5, 3, 5, 7)
		checkResult(t, "minmax(5, 5, 3, 5, 7)", got, []float64{5})
	})

	t.Run("min igual a max sin coincidencia", func(t *testing.T) {
		got := minmax(5, 5, 3, 4, 6)
		checkResult(t, "minmax(5, 5, 3, 4, 6)", got, nil)
	})

	t.Run("valores en los límites", func(t *testing.T) {
		got := minmax(2, 8, 2, 8)
		checkResult(t, "minmax(2, 8, 2, 8)", got, []float64{2, 8})
	})
}
