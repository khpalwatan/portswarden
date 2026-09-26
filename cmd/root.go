package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "portswarden",
	Short:   "Find and free Windows ports — no more EADDRINUSE",
	Version: "1.0.0",
}

func Execute() error {
	return rootCmd.Execute()
}