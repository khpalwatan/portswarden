package cmd

import (
	"fmt"
	"os"

	"github.com/khpalwatan/portswarden/internal/ports"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listPortFilter string

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List listening ports",
	RunE: func(c *cobra.Command, _ []string) error {
		data, err := ports.List()
		if err != nil {
			return err
		}

		table := tablewriter.NewTable(os.Stdout)
		table.Header([]string{"PORT", "PROTO", "STATE", "PID", "PROCESS", "PATH"})

		for _, p := range data {
			if listPortFilter != "" && fmt.Sprint(p.Port) != listPortFilter {
				continue
			}
			_ = table.Append([]any{
				fmt.Sprint(p.Port),
				p.Proto,
				p.State,
				fmt.Sprint(p.PID),
				p.Process,
				p.Path,
			})
		}
		_ = table.Render()
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&listPortFilter, "port", "p", "", "filter by port")
	rootCmd.AddCommand(listCmd)
}