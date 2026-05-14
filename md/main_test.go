package main

import (
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	outName := "test_output"
	expectedFile := outName + ".html"

	defer os.Remove(expectedFile)

	err := run(outName)
	if err != nil {
		t.Fatalf("run() devolvió un error inesperado: %v", err)
	}

	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatalf("Se esperaba que existiera el archivo %s, pero no fue creado", expectedFile)
	}

	content, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("No se pudo leer el archivo generado: %v", err)
	}

	if !strings.Contains(string(content), strings.TrimSpace(header)) {
		t.Error("El archivo generado no contiene el header esperado")
	}

	if !strings.Contains(string(content), strings.TrimSpace(footer)) {
		t.Error("El archivo generado no contiene el footer esperado")
	}
}
