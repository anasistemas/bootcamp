package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "echo 'hola desde stdout'; echo 'hola desde stderr' >&2")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("El comando falló:", err)
	}

	fmt.Println("STDOUT:", stdout.String())
	fmt.Println("STDERR:", stderr.String())
}
