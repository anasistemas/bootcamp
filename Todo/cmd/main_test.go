package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	if err := exec.Command("go", "build", "-o", binName).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	if err := os.WriteFile(fileName, []byte{}, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file %s", fileName)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up....")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "New Task"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	run := func(t *testing.T, args ...string) {
		t.Helper()
		if err := exec.Command(cmdPath, args...).Run(); err != nil {
			t.Fatal(err)
		}
	}

	output := func(t *testing.T, args ...string) string {
		t.Helper()
		out, err := exec.Command(cmdPath, args...).CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		return string(out)
	}

	t.Run("AddNewTask", func(t *testing.T) {
		run(t, "-task", task)
	})

	t.Run("ListTasks", func(t *testing.T) {
		out := output(t, "-list")
		expected := "- [ ] 1: " + task
		if !strings.Contains(out, expected) {
			t.Errorf("esperaba %q en la salida, pero se obtuvo: %s", expected, out)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		run(t, "-complete", "1")

		out := output(t, "-list")
		expected := "- [X] 1: " + task
		if !strings.Contains(out, expected) {
			t.Errorf("esperaba tarea completada %q en la salida, pero se obtuvo: %s", expected, out)
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		run(t, "-task", "Tarea para borrar")

		run(t, "-delete", "2")

		out := output(t, "-list")
		if strings.Contains(out, "Tarea para borrar") {
			t.Errorf("la tarea borrada no debería aparecer en la lista, pero se obtuvo: %s", out)
		}
	})
}
