package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	l := flag.Bool("l", false, "contar lineas")
	b := flag.Bool("b", false, "contar bytes")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	words := 0
	lines := 0
	bytes := 0

	for scanner.Scan() {
		text := scanner.Text()

		if strings.ToLower(text) == "exit" {
			break
		}

		lines++
		words += len(strings.Fields(text))
		bytes += len(text)
	}

	if *l {
		fmt.Println(lines)
	} else if *b {
		fmt.Println(bytes)
	} else {
		fmt.Println(words)
	}
}

func procesarEntrada(entrada string, l bool, b bool) int {
	words := 0
	lines := 0
	bytes := 0

	scanner := bufio.NewScanner(strings.NewReader(entrada))

	for scanner.Scan() {
		text := scanner.Text()

		if strings.ToLower(text) == "exit" {
			break
		}

		lines++
		words += len(strings.Fields(text))
		bytes += len(text)
	}

	if l {
		return lines
	}
	if b {
		return bytes
	}
	return words
}
