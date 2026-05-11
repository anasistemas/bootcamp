package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", `
		echo "paso 1"
		sleep 0.3
		echo "paso 2"
		sleep 0.3
		echo "paso 3"
		sleep 0.3
		echo "listo!"
	`)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	fmt.Println("Corriendo el comando...")
	err := cmd.Run()
	fmt.Println("Comando terminado!")

	if err != nil {
		fmt.Println("Hubo un error:", err)
	}

	fmt.Println("\nLo que guardamos en el buffer:")
	fmt.Println(stdout.String())
}
