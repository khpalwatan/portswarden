package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/khpalwatan/portswarden/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		color.Red("%v", err)
		os.Exit(1)
	}
}