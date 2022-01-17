package main

import (
	"os"

	"github.com/ionut-maxim/notarry/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
