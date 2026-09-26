package cmd

import (
	"strconv"
	"time"

	"github.com/khpalwatan/portswarden/internal/ports"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch <port>",
	Short: "Auto-kill anything that grabs a port",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		port, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		color.Cyan("watching port %d — Ctrl+C to stop", port)
		for {
			list, _ := ports.List()
			for _, p := range list {
				if p.Port == uint32(port) {
					color.Yellow("killing %s (pid %d) on :%d", p.Process, p.PID, port)
					_ = ports.KillByPort(uint32(port), true)
				}
			}
			time.Sleep(time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}