package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const header = `<!DOCTYPE html>
  <html>
    <head>
      <meta http-equiv="content-type" content="text/html; charset=utf-8" />
      <title>Markdown Preview Tool</title>
    </head>
    <body>
`

const footer = `
    </body>
  </html>
`

func main() {
	in := flag.String("in", "", "Archivo markdown que quieres convertir")
	out := flag.String("out", "", "Nombre del archivo HTML resultante (sin .html)")
	flag.Parse()

	if *in == "" {
		fmt.Fprintln(os.Stderr, "Falta el archivo de entrada. Usa: -in archivo.md")
		os.Exit(1)
	}

	if err := run(*in, *out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(inFile, outFile string) error {
	input, err := os.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo %s: %w", inFile, err)
	}

	body := parseContent(input)

	if outFile == "" {
		outFile = filepath.Base(inFile)
	}

	return saveHTML(outFile+".html", body)
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var result []byte
	result = append(result, []byte(header)...)
	result = append(result, body...)
	result = append(result, []byte(footer)...)

	return result
}

func saveHTML(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}
