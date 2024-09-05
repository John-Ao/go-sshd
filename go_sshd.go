package main

import (
	"os"

	"github.com/John-Ao/go-sshd/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		os.Exit(-1)
	}
}
