package cmd

import (
	"github.com/khpalwatan/portswarden/internal/ports"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var excludedCmd = &cobra.Command{
	Use:   "excluded",
	Short: "Show Windows-reserved (excluded) TCP port ranges",
	RunE: func(_ *cobra.Command, _ []string) error {
		ranges, err := ports.ExcludedRanges()
		if err != nil {
			return err
		}
		if len(ranges) == 0 {
			color.Green("no excluded ranges")
			return nil
		}
		color.Yellow("Windows has reserved these TCP port ranges:")
		for _, r := range ranges {
			color.Yellow("  %s", r)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(excludedCmd)
}