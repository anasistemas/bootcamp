package main

import (
	"flag"
	"fmt"
	"os"
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
	out := flag.String("out", "", "Nombre del archivo HTML de salida (sin extensión .html)")
	flag.Parse()

	if *out == "" {
		fmt.Fprintln(os.Stderr, "Error: debes especificar el nombre del archivo con el flag -out")
		os.Exit(1)
	}

	if err := run(*out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(outFile string) error {
	content := []byte(header + footer)

	fileName := outFile + ".html"

	return saveHTML(fileName, content)
}

func saveHTML(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}
