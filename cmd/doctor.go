package cmd

import (
	"fmt"
	"os/exec"

	"github.com/khpalwatan/portswarden/internal/ports"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check environment and print diagnostics",
	RunE: func(_ *cobra.Command, _ []string) error {
		color.Cyan("portswarden doctor")
		fmt.Println()

		// netsh
		check("netsh reachable", func() bool {
			_, err := exec.LookPath("netsh")
			return err == nil
		})

		// taskkill
		check("taskkill reachable", func() bool {
			_, err := exec.LookPath("taskkill")
			return err == nil
		})

		// port list
		list, err := ports.List()
		check("port list readable", func() bool { return err == nil })
		if err == nil {
			fmt.Printf("    %d listening ports found\n", len(list))
		}
		fmt.Println()

		// reserved ranges
		ranges, err := ports.ExcludedRanges()
		check("excluded ranges readable", func() bool { return err == nil })
		if err == nil {
			if len(ranges) > 0 {
				color.Yellow("    Windows has reserved %d TCP port ranges:", len(ranges))
				for _, r := range ranges {
					fmt.Printf("      %s\n", r)
				}
			} else {
				fmt.Println("    no reserved ranges detected")
			}
		}

		return nil
	},
}

func check(name string, fn func() bool) {
	if fn() {
		color.Green("  OK  %s", name)
	} else {
		color.Red(" FAIL %s", name)
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}