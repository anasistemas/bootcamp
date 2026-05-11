package fl

import "os"

type Flag struct {
	value       bool
	description string
}

var flags = map[string]*Flag{}

func Bool(cmd string, value bool, description string) *bool {
	f := &Flag{value: value, description: description}
	flags[cmd] = f
	return &f.value
}

func Parse() {
	args := os.Args[1:]
	for _, arg := range args {
		if flag, ok := flags[arg]; ok {
			flag.value = true
		}
	}
}
