package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	cmd.Stderr = &buffer

	err := cmd.Run()
	if err != nil {
		fmt.Println("El comando falló:", err)
		return
	}

	fmt.Println("Salida capturada manualmente:")
	fmt.Println(buffer.String())
}
