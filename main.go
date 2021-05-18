package main

import (
	"os"

	"tcp_echo_server/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
