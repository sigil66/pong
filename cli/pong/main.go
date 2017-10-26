package main

import (
	"os"

	"github.com/solvent-io/pong/cli/pong/commands"
)

func main() {
	command := commands.NewPongRootCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
