package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunUsingOutFlag(t *testing.T) {
	outName := "test_output"
	expectedFile := outName + ".html"
	defer os.Remove(expectedFile)

	var buf bytes.Buffer
	if err := run("testdata/test.md", outName, &buf); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatalf("el archivo %s no fue creado", expectedFile)
	}

	if got := strings.TrimSpace(buf.String()); got != expectedFile {
		t.Errorf("se esperaba %q pero se imprimió %q", expectedFile, got)
	}
}

func TestRunWithoutOutFlag(t *testing.T) {
	var buf bytes.Buffer
	if err := run("testdata/test.md", "", &buf); err != nil {
		t.Fatal(err)
	}

	generatedFile := filepath.Base(strings.TrimSpace(buf.String()))
	defer os.Remove(generatedFile)

	if !strings.HasPrefix(generatedFile, "md") || !strings.HasSuffix(generatedFile, ".html") {
		t.Errorf("nombre de archivo inesperado: %s", generatedFile)
	}

	if _, err := os.Stat(generatedFile); os.IsNotExist(err) {
		t.Fatalf("el archivo temporal %s no existe", generatedFile)
	}
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile("testdata/test.md")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile("testdata/test.html")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(parseContent(input), expected) {
		t.Error("el resultado no coincide con el golden file")
	}
}
