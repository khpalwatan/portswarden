package cmd

import (
	"strconv"

	"github.com/khpalwatan/portswarden/internal/ports"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var killForce bool

var killCmd = &cobra.Command{
	Use:   "kill <port>",
	Short: "Kill the process listening on a port",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		port, err := strconv.Atoi(args[0])
		if err != nil {
			return err
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
	rootCmd.AddCommand(killCmd)
}