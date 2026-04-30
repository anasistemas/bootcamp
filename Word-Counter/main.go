package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	contarLineas := false

	for _, arg := range os.Args[1:] {
		if arg == "-l" {
			contarLineas = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	palabras := 0
	lineas := 0

	for scanner.Scan() {
		texto := scanner.Text()

		if texto == "salida" {
			break
		}

		lineas++
		palabras += len(strings.Fields(texto))
	}

	if contarLineas {
		fmt.Println(lineas)
	} else {
		fmt.Println(palabras)
	}
}

func procesarEntrada(entrada string, contarLineas bool) int {
	lineas := 0
	palabras := 0

	scanner := bufio.NewScanner(strings.NewReader(entrada))

	for scanner.Scan() {
		texto := scanner.Text()

		if strings.ToLower(texto) == "exit" {
			break
		}

		lineas++
		palabras += len(strings.Fields(texto))
	}

	if contarLineas {
		return lineas
	}
	return palabras
}
