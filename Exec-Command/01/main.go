package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	err := cmd.Run()
	if err != nil {
		fmt.Println("El comando falló:", err)
		return
	}

	fmt.Println("El comando funcionó")
}
