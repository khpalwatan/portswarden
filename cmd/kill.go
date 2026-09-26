package cmd

import (
	"fmt"
	"strconv"

	"github.com/khpalwatan/portswarden/internal/ports"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	killForce bool
	killYes   bool
)

var killCmd = &cobra.Command{
	Use:   "kill <port>",
	Short: "Kill the process listening on a port",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		port, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		list, _ := ports.List()
		var target *ports.PortInfo
		for i := range list {
			if list[i].Port == uint32(port) {
				target = &list[i]
				break
			}
		}
		if target == nil {
			return fmt.Errorf("no listening process found on port %d", port)
		}

		if !killYes {
			color.Yellow("About to kill:")
			fmt.Printf("  Port    : %d\n", target.Port)
			fmt.Printf("  Process : %s\n", target.Process)
			fmt.Printf("  PID     : %d\n", target.PID)
			fmt.Print("\nContinue? [y/N] ")

			var resp string
			_, _ = fmt.Scanln(&resp)
			if resp != "y" && resp != "Y" {
				color.Cyan("aborted")
				return nil
			}
		}

		if err := ports.KillByPort(uint32(port), killForce); err != nil {
			return err
		}
		color.Green("✔ freed port %d", port)
		return nil
	},
}

func init() {
	killCmd.Flags().BoolVarP(&killForce, "force", "f", true, "force kill")
	killCmd.Flags().BoolVarP(&killYes, "yes", "y", false, "skip confirmation")
	rootCmd.AddCommand(killCmd)
}