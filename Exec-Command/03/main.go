package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	salida, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("El comando falló:", err)
		return
	}

	fmt.Println("Esto se capturó:")
	fmt.Println(string(salida))
}
