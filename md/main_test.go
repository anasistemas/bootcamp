package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	outName := "test_output"
	expectedFile := outName + ".html"

	defer os.Remove(expectedFile)

	err := run("testdata/test.md", outName)
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

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile("testdata/test.md")
	if err != nil {
		t.Fatalf("No se pudo leer testdata/test.md: %v", err)
	}

	result := parseContent(input)

	expected, err := os.ReadFile("testdata/test.html")
	if err != nil {
		t.Fatalf("No se pudo leer testdata/test.html: %v", err)
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("El resultado no coincide con el golden file.\nObtenido:\n%s\nEsperado:\n%s", result, expected)
	}
}
