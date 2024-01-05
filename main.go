package main

import (
	"os"

	"github.com/code-gorilla-au/goety/internal/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
