package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func minmax(min, max float64, values ...float64) []float64 {
	var result []float64
	for _, v := range values {
		if v >= min && v <= max {
			result = append(result, v)
		}
	}
	return result
}

func parseFloats(input string) ([]float64, error) {
	parts := strings.Fields(input)
	var nums []float64
	for _, p := range parts {
		n, err := strconv.ParseFloat(p, 64)
		if err != nil {
			return nil, fmt.Errorf("valor inválido: %s", p)
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func main() {
	fmt.Print("Ingresa el valor mínimo: ")
	minInput := getInput()
	minVal, err := strconv.ParseFloat(strings.TrimSpace(minInput), 64)
	if err != nil {
		fmt.Println("Error: valor mínimo inválido.")
		return
	}

	fmt.Print("Ingresa el valor máximo: ")
	maxInput := getInput()
	maxVal, err := strconv.ParseFloat(strings.TrimSpace(maxInput), 64)
	if err != nil {
		fmt.Println("Error: valor máximo inválido.")
		return
	}

	fmt.Print("Ingresa la lista de valores (separados por espacios): ")
	valuesInput := getInput()
	values, err := parseFloats(valuesInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := minmax(minVal, maxVal, values...)
	fmt.Println("Valores en rango:", result)
}